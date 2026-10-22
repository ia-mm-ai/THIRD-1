package avalanchevm

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

func exactMethod(method string, handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != method {
			writer.Header().Set("Allow", method)
			writeAPIError(writer, http.StatusMethodNotAllowed, fmt.Errorf("method %s is not allowed", request.Method))
			return
		}
		handler(writer, request)
	})
}

func routeValue(request *http.Request, name string) string {
	if value := request.PathValue(name); value != "" {
		return value
	}
	return path.Base(strings.TrimRight(request.URL.Path, "/"))
}

func (vm *VM) apiHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /operations", vm.handleIssueOperation)
	mux.HandleFunc("GET /receipts/{operationID}", vm.handleGetReceipt)
	mux.HandleFunc("GET /state", vm.handleState)
	mux.HandleFunc("GET /genesis", vm.handleGenesis)
	mux.HandleFunc("GET /status", vm.handleStatus)
	return mux
}

func (vm *VM) handleIssueOperation(writer http.ResponseWriter, request *http.Request) {
	if !vm.isInitialized() {
		writeAPIError(writer, http.StatusServiceUnavailable, errNotInitialized)
		return
	}
	body, err := io.ReadAll(http.MaxBytesReader(writer, request.Body, int64(vm.maxOperationBytes())))
	if err != nil {
		writeAPIError(writer, http.StatusRequestEntityTooLarge, err)
		return
	}
	operationID, err := vm.IssueOperation(body)
	if err != nil {
		if code, ok := refusalCode(err); ok {
			writeJSON(writer, http.StatusConflict, map[string]any{"accepted": false, "code": code})
			return
		}
		writeAPIError(writer, http.StatusBadRequest, err)
		return
	}
	writeJSON(writer, http.StatusAccepted, map[string]any{
		"operation_id": operationID, "status": "PENDING_CONSENSUS",
	})
}

func (vm *VM) handleGetReceipt(writer http.ResponseWriter, request *http.Request) {
	id := routeValue(request, "operationID")
	if err := requireHash("operation_id", id); err != nil {
		writeAPIError(writer, http.StatusBadRequest, err)
		return
	}
	vm.lock.RLock()
	if !vm.initialized {
		vm.lock.RUnlock()
		writeAPIError(writer, http.StatusServiceUnavailable, errNotInitialized)
		return
	}
	data, err := vm.db.Get(receiptKey(id))
	data = bytes.Clone(data)
	vm.lock.RUnlock()
	if errors.Is(err, database.ErrNotFound) {
		writeAPIError(writer, http.StatusNotFound, err)
		return
	}
	if err != nil {
		writeAPIError(writer, http.StatusInternalServerError, err)
		return
	}
	receipt := new(Receipt)
	if err := decodeCanonical(data, receipt); err != nil {
		writeAPIError(writer, http.StatusInternalServerError, err)
		return
	}
	writeJSON(writer, http.StatusOK, receipt)
}

func (vm *VM) handleState(writer http.ResponseWriter, _ *http.Request) {
	vm.lock.Lock()
	if !vm.initialized {
		vm.lock.Unlock()
		writeAPIError(writer, http.StatusServiceUnavailable, errNotInitialized)
		return
	}
	state, err := vm.acceptedStateLocked()
	vm.lock.Unlock()
	if err != nil {
		writeAPIError(writer, http.StatusInternalServerError, err)
		return
	}
	stateHash, err := kernel.ChainStateHash(state)
	if err != nil {
		writeAPIError(writer, http.StatusInternalServerError, err)
		return
	}
	writeJSON(writer, http.StatusOK, map[string]any{"state_hash": stateHash, "state": state})
}

func (vm *VM) handleGenesis(writer http.ResponseWriter, _ *http.Request) {
	vm.lock.RLock()
	if !vm.initialized {
		vm.lock.RUnlock()
		writeAPIError(writer, http.StatusServiceUnavailable, errNotInitialized)
		return
	}
	genesis := *vm.genesis
	genesisSHA256 := hex.EncodeToString(vm.genesisBlock.id[:])
	genesisBlockID := vm.genesisBlock.id.String()
	vm.lock.RUnlock()
	writeJSON(writer, http.StatusOK, map[string]any{
		"genesis": genesis, "genesis_sha256": genesisSHA256,
		"genesis_block_id": genesisBlockID,
	})
}

func (vm *VM) handleStatus(writer http.ResponseWriter, _ *http.Request) {
	vm.lock.Lock()
	if !vm.initialized {
		vm.lock.Unlock()
		writeAPIError(writer, http.StatusServiceUnavailable, errNotInitialized)
		return
	}
	block, err := vm.getBlockLocked(vm.lastAcceptedID)
	if err != nil || block.postState == nil {
		vm.lock.Unlock()
		writeAPIError(writer, http.StatusInternalServerError, err)
		return
	}
	state := *block.postState
	lastAcceptedID := vm.lastAcceptedID.String()
	height := block.height
	pendingOperations := vm.durablePendingCount
	pendingBytes := vm.durablePendingBytes
	vm.lock.Unlock()
	stateHash, err := kernel.ChainStateHash(state)
	if err != nil {
		writeAPIError(writer, http.StatusInternalServerError, err)
		return
	}
	writeJSON(writer, http.StatusOK, map[string]any{
		"healthy": true, "implementation": ImplementationID, "vm_version": VMVersion,
		"rpcchainvm_protocol": version.RPCChainVMProtocol, "avalanchego_target": AvalancheGoTarget,
		"last_accepted_block": lastAcceptedID, "height": height,
		"state_hash": stateHash, "profile_commitment": state.ProfileCommitment,
		"relations": len(state.Relations),
		"closures": len(state.Closures), "successors": len(state.Successors),
		"pending_operations": pendingOperations,
		"pending_bytes": pendingBytes,
	})
}

func (vm *VM) isInitialized() bool {
	vm.lock.RLock()
	defer vm.lock.RUnlock()
	return vm.initialized
}

func (vm *VM) maxOperationBytes() uint32 {
	vm.lock.RLock()
	defer vm.lock.RUnlock()
	if vm.genesis == nil {
		return 1024
	}
	return vm.genesis.Limits.MaxOperationBytes
}

func writeJSON(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(value)
}

func writeAPIError(writer http.ResponseWriter, status int, err error) {
	if err == nil {
		err = errors.New("unknown error")
	}
	writeJSON(writer, status, map[string]string{"error": err.Error()})
}
