package avalanchevm

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/version"
	"github.com/ia-mm-ai/THIRD-1/kernel"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var _ smblock.ChainVM = (*VM)(nil)

var (
	genesisBytesKey = []byte("continuity/meta/genesis-bytes")
	genesisIDKey    = []byte("continuity/meta/genesis-id")
	lastAcceptedKey = []byte("continuity/meta/last-accepted")
	currentStateKey = []byte("continuity/state/current")
	pendingPrefix   = []byte("continuity/pending/")
	retiredPrefix   = []byte("continuity/retired/")
)

const (
	pendingOperationSchema = "CONTINUITY_PENDING_OPERATION_001"
	retiredOperationSchema = "CONTINUITY_RETIRED_OPERATION_001"
	pendingRetryInitialDelay = 100 * time.Millisecond
	pendingRetryMaximumDelay = 5 * time.Second
)

type durablePendingOperation struct {
	Schema            string `json:"schema"`
	Version           uint16 `json:"version"`
	Sequence          uint64 `json:"sequence"`
	OperationBytesHex string `json:"operation_bytes_hex"`
}

type durableRetiredOperation struct {
	Schema            string      `json:"schema"`
	Version           uint16      `json:"version"`
	Sequence          uint64      `json:"sequence"`
	OperationBytesHex string      `json:"operation_bytes_hex"`
	Code              kernel.Code `json:"code"`
}

type VM struct {
	lock sync.RWMutex

	db             database.Database
	genesis        *Genesis
	genesisBytes   []byte
	genesisBlock   *vmBlock
	blocks         map[ids.ID]*vmBlock
	lastAcceptedID ids.ID
	preferredID    ids.ID
	chainState     snow.State
	initialized    bool

	pending             map[string]*Operation
	pendingSequence     map[string]uint64
	pendingOrder        []string
	durablePendingCount uint32
	durablePendingBytes uint64
	nextPendingSequence uint64
	notify              chan struct{}
	shutdown            chan struct{}
	pendingRetryDelay   time.Duration
	pendingRetryAfter   time.Time
}

func (vm *VM) Initialize(
	ctx context.Context,
	chainContext *snow.Context,
	db database.Database,
	genesisBytes []byte,
	upgradeBytes []byte,
	configBytes []byte,
	_ []*common.Fx,
	_ common.AppSender,
) error {
	if uint(version.RPCChainVMProtocol) != RequiredRPCProtocol {
		return fmt.Errorf("RPCChainVM protocol %d does not equal required %d", version.RPCChainVMProtocol, RequiredRPCProtocol)
	}
	genesis, err := ParseGenesis(genesisBytes)
	if err != nil {
		return fmt.Errorf("parse Continuity genesis: %w", err)
	}
	chainConfig, err := genesis.chainConfigForContext(chainContext)
	if err != nil {
		return fmt.Errorf("validate Continuity genesis security domain: %w", err)
	}
	initialState, err := kernel.NewChainState(chainConfig)
	if err != nil {
		return fmt.Errorf("initialize kernel state: %w", err)
	}
	if _, err := encodeBoundedState(initialState, genesis.Limits); err != nil {
		return err
	}
	if len(upgradeBytes) != 0 {
		return errors.New("Continuity candidate does not implement VM upgrade bytes")
	}
	if len(configBytes) != 0 {
		return errors.New("Continuity candidate does not admit operator config bytes")
	}
	genesisBlock := newGenesisBlock(vm, genesisBytes, initialState)
	storedGenesis, err := db.Get(genesisBytesKey)
	switch {
	case errors.Is(err, database.ErrNotFound):
		stateBytes, err := encodeBoundedState(initialState, genesis.Limits)
		if err != nil {
			return err
		}
		vdb := versiondb.New(db)
		for _, write := range []struct{ key, value []byte }{
			{genesisBytesKey, genesisBytes},
			{genesisIDKey, genesisBlock.id[:]},
			{lastAcceptedKey, genesisBlock.id[:]},
			{currentStateKey, stateBytes},
			{blockKey(genesisBlock.id), genesisBytes},
			{stateByBlockKey(genesisBlock.id), stateBytes},
			{heightKey(0), genesisBlock.id[:]},
		} {
			if err := vdb.Put(write.key, write.value); err != nil {
				return err
			}
		}
		if err := vdb.Commit(); err != nil {
			return fmt.Errorf("commit genesis state: %w", err)
		}
	case err != nil:
		return fmt.Errorf("read stored genesis: %w", err)
	case !bytes.Equal(storedGenesis, genesisBytes):
		return errors.New("stored immutable genesis differs from supplied genesis")
	default:
		storedID, err := db.Get(genesisIDKey)
		if err != nil || !bytes.Equal(storedID, genesisBlock.id[:]) {
			return errors.New("stored genesis ID mismatch")
		}
	}
	lastAcceptedBytes, err := db.Get(lastAcceptedKey)
	if err != nil {
		return err
	}
	lastAcceptedID, err := ids.ToID(lastAcceptedBytes)
	if err != nil {
		return err
	}

	vm.lock.Lock()
	defer vm.lock.Unlock()
	vm.db = db
	vm.genesis = genesis
	vm.genesisBytes = bytes.Clone(genesisBytes)
	vm.genesisBlock = genesisBlock
	vm.blocks = map[ids.ID]*vmBlock{genesisBlock.id: genesisBlock}
	vm.lastAcceptedID = lastAcceptedID
	vm.preferredID = lastAcceptedID
	vm.chainState = snow.Initializing
	vm.pending = make(map[string]*Operation)
	vm.pendingSequence = make(map[string]uint64)
	vm.nextPendingSequence = 1
	vm.notify = make(chan struct{}, 1)
	vm.shutdown = make(chan struct{})
	vm.resetPendingRetryLocked()
	accepted, err := vm.replayAcceptedHistoryLocked(ctx, lastAcceptedID, initialState)
	if err != nil {
		return fmt.Errorf("replay accepted Continuity history: %w", err)
	}
	vm.blocks[lastAcceptedID] = accepted
	if err := vm.loadPendingOperationsLocked(); err != nil {
		return fmt.Errorf("load durable pending operations: %w", err)
	}
	vm.initialized = true
	vm.signalPendingLocked()
	return nil
}

func (vm *VM) SetState(_ context.Context, state snow.State) error {
	if state != snow.Initializing && state != snow.StateSyncing && state != snow.Bootstrapping && state != snow.NormalOp {
		return snow.ErrUnknownState
	}
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return errNotInitialized
	}
	enteringNormalOp := vm.chainState != snow.NormalOp && state == snow.NormalOp
	vm.chainState = state
	if enteringNormalOp {
		vm.resetPendingRetryLocked()
		vm.signalPendingLocked()
	}
	return nil
}

func (vm *VM) Shutdown(context.Context) error {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return nil
	}
	vm.initialized = false
	close(vm.shutdown)
	vm.resetPendingRetryLocked()
	return vm.db.Close()
}

func (*VM) Version(context.Context) (string, error) { return VMVersion, nil }

func (vm *VM) CreateHandlers(context.Context) (map[string]http.Handler, error) {
	return map[string]http.Handler{
		"/operations":             exactMethod(http.MethodPost, vm.handleIssueOperation),
		"/receipts/{operationID}": exactMethod(http.MethodGet, vm.handleGetReceipt),
		"/state":                  exactMethod(http.MethodGet, vm.handleState),
		"/genesis":                exactMethod(http.MethodGet, vm.handleGenesis),
		"/status":                 exactMethod(http.MethodGet, vm.handleStatus),
	}, nil
}

func (vm *VM) NewHTTPHandler(context.Context) (http.Handler, error) { return vm.apiHandler(), nil }

func (vm *VM) HealthCheck(context.Context) (interface{}, error) {
	vm.lock.Lock()
	if !vm.initialized {
		vm.lock.Unlock()
		return nil, errNotInitialized
	}
	block, err := vm.getBlockLocked(vm.lastAcceptedID)
	if err != nil || block.postState == nil {
		vm.lock.Unlock()
		return nil, err
	}
	state := *block.postState
	lastAcceptedID := vm.lastAcceptedID.String()
	height := block.height
	pendingOperations := vm.durablePendingCount
	pendingBytes := vm.durablePendingBytes
	vm.lock.Unlock()
	stateHash, err := kernel.ChainStateHash(state)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"healthy": true, "vmVersion": VMVersion,
		"rpcChainVMProtocol": version.RPCChainVMProtocol,
		"lastAcceptedBlockID": lastAcceptedID, "height": height,
		"stateHash": stateHash, "profileCommitment": state.ProfileCommitment,
		"relations": len(state.Relations),
		"closures": len(state.Closures), "pendingOperations": pendingOperations,
		"pendingBytes": pendingBytes,
	}, nil
}

func (vm *VM) WaitForEvent(ctx context.Context) (common.Message, error) {
	for {
		vm.lock.RLock()
		if !vm.initialized {
			vm.lock.RUnlock()
			return 0, errNotInitialized
		}
		hasPending := len(vm.pendingOrder) > 0
		retryAfter := vm.pendingRetryAfter
		notify := vm.notify
		shutdown := vm.shutdown
		vm.lock.RUnlock()
		if hasPending && (retryAfter.IsZero() || !time.Now().Before(retryAfter)) {
			return common.PendingTxs, nil
		}
		var retryTimer *time.Timer
		var retry <-chan time.Time
		if hasPending {
			retryTimer = time.NewTimer(time.Until(retryAfter))
			retry = retryTimer.C
		}
		select {
		case <-ctx.Done():
			stopAndDrainTimer(retryTimer)
			return 0, ctx.Err()
		case <-shutdown:
			stopAndDrainTimer(retryTimer)
			return 0, errNotInitialized
		case <-notify:
			stopAndDrainTimer(retryTimer)
		case <-retry:
		}
	}
}

func (vm *VM) BuildBlock(context.Context) (built snowman.Block, err error) {
	vm.lock.Lock()
	defer func() {
		if err != nil {
			vm.armPendingRetryLocked()
		} else {
			vm.resetPendingRetryLocked()
		}
		vm.lock.Unlock()
	}()
	if !vm.initialized {
		return nil, errNotInitialized
	}
	if vm.chainState != snow.NormalOp {
		return nil, fmt.Errorf("cannot build while chain state is %s", vm.chainState)
	}
	parent, err := vm.getBlockLocked(vm.preferredID)
	if err != nil || parent.postState == nil || parent.status == choices.Rejected {
		return nil, errInvalidPreference
	}
	workingState := *parent.postState
	timestamp := time.Now().UTC().Truncate(time.Second)
	if timestamp.Before(parent.timestamp) {
		timestamp = parent.timestamp
	}
	limit := int(vm.genesis.Limits.MaxOperationsPerBlock)
	selected := make([]*Operation, 0, limit)
	selectedIDs := make([]string, 0, limit)
	remaining := make([]string, 0, len(vm.pendingOrder))
	var firstRefusal error
	for _, operationID := range vm.pendingOrder {
		operation, exists := vm.pending[operationID]
		if !exists {
			continue
		}
		if len(selected) >= limit {
			remaining = append(remaining, operationID)
			continue
		}
		result, err := vm.applyBoundedOperation(workingState, operation)
		if err != nil {
			remaining = append(remaining, operationID)
			if firstRefusal == nil {
				firstRefusal = fmt.Errorf("pending operation %s: %w", operationID, err)
			}
			continue
		}
		candidate := append(append([]*Operation(nil), selected...), operation)
		candidateBytes, encodeErr := canonicalJSON(&WireBlock{
			Schema: BlockSchema, ParentID: parent.id.String(), Timestamp: timestamp.Unix(),
			Height: parent.height + 1, Operations: candidate,
		})
		if encodeErr != nil {
			return nil, encodeErr
		}
		if uint64(len(candidateBytes)) > uint64(vm.genesis.Limits.MaxBlockBytes) {
			remaining = append(remaining, operationID)
			continue
		}
		workingState = result.State
		selected = candidate
		selectedIDs = append(selectedIDs, operationID)
	}
	if len(selected) == 0 {
		vm.pendingOrder = remaining
		if firstRefusal != nil {
			return nil, fmt.Errorf("%w: %w", errNoPendingOperations, firstRefusal)
		}
		return nil, errNoPendingOperations
	}
	var block *vmBlock
	for {
		block, err = newWireBlock(vm, &WireBlock{
			Schema: BlockSchema, ParentID: parent.id.String(), Timestamp: timestamp.Unix(),
			Height: parent.height + 1, Operations: selected,
		}, true)
		if err != nil {
			return nil, err
		}
		if _, exists := vm.blocks[block.id]; !exists {
			break
		}
		timestamp = timestamp.Add(time.Second)
		if time.Until(timestamp) > maxClockSkew {
			return nil, errors.New("cannot construct a fresh block ID inside maximum clock skew")
		}
	}
	for _, operationID := range selectedIDs {
		delete(vm.pending, operationID)
	}
	vm.pendingOrder = remaining
	vm.signalPendingLocked()
	return block, nil
}

func (vm *VM) ParseBlock(_ context.Context, data []byte) (snowman.Block, error) {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return nil, errNotInitialized
	}
	if bytes.Equal(data, vm.genesisBytes) {
		return vm.genesisBlock, nil
	}
	block, err := blockFromBytes(vm, data, choices.Processing)
	if err != nil {
		return nil, err
	}
	if existing, exists := vm.blocks[block.id]; exists {
		return existing, nil
	}
	if _, err := vm.db.Get(blockKey(block.id)); err == nil {
		accepted, err := vm.loadAcceptedBlockLocked(block.id)
		if err != nil {
			return nil, err
		}
		if block.id == vm.lastAcceptedID || block.id == vm.genesisBlock.id {
			vm.blocks[block.id] = accepted
		}
		return accepted, nil
	} else if !errors.Is(err, database.ErrNotFound) {
		return nil, err
	}
	return block, nil
}

func (vm *VM) GetBlock(_ context.Context, blockID ids.ID) (snowman.Block, error) {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return nil, errNotInitialized
	}
	return vm.getBlockLocked(blockID)
}

func (vm *VM) SetPreference(_ context.Context, blockID ids.ID) error {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return errNotInitialized
	}
	block, err := vm.getBlockLocked(blockID)
	if err != nil || block.status == choices.Rejected || block.postState == nil {
		return errInvalidPreference
	}
	if vm.preferredID != blockID {
		vm.preferredID = blockID
		vm.resetPendingRetryLocked()
		vm.signalPendingLocked()
	}
	return nil
}

func (vm *VM) LastAccepted(context.Context) (ids.ID, error) {
	vm.lock.RLock()
	defer vm.lock.RUnlock()
	if !vm.initialized {
		return ids.Empty, errNotInitialized
	}
	return vm.lastAcceptedID, nil
}

func (vm *VM) GetBlockIDAtHeight(_ context.Context, height uint64) (ids.ID, error) {
	vm.lock.RLock()
	defer vm.lock.RUnlock()
	if !vm.initialized {
		return ids.Empty, errNotInitialized
	}
	data, err := vm.db.Get(heightKey(height))
	if err != nil {
		return ids.Empty, err
	}
	return ids.ToID(data)
}

func (vm *VM) IssueOperation(data []byte) (string, error) {
	vm.lock.RLock()
	if !vm.initialized {
		vm.lock.RUnlock()
		return "", errNotInitialized
	}
	maximum := vm.genesis.Limits.MaxOperationBytes
	vm.lock.RUnlock()
	if uint64(len(data)) > uint64(maximum) {
		return "", errors.New("operation exceeds genesis byte limit")
	}
	operation, operationID, err := parseOperation(data)
	if err != nil {
		return "", err
	}
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return "", errNotInitialized
	}
	if accepted, err := vm.db.Get(operationKey(operationID)); err == nil {
		if !bytes.Equal(accepted, data) {
			return "", errors.New("accepted operation bytes differ from canonical operation ID")
		}
		if err := vm.purgeAcceptedPendingCollisionLocked(operationID, data); err != nil {
			return "", err
		}
		vm.rebuildPendingOrderLocked()
		return operationID, nil
	} else if !errors.Is(err, database.ErrNotFound) {
		return "", err
	}
	if _, exists := vm.pending[operationID]; exists {
		return operationID, nil
	}
	if persisted, err := vm.db.Get(pendingKey(operationID)); err == nil {
		stored, persistedID, _, decodeErr := decodeDurablePendingOperation(persisted, maximum)
		if decodeErr != nil {
			return "", fmt.Errorf("decode existing durable pending operation: %w", decodeErr)
		}
		if persistedID != operationID || !bytes.Equal(stored.Bytes(), data) {
			return "", errors.New("durable pending operation differs from canonical operation ID")
		}
		// A locally built unresolved block may temporarily own the in-memory
		// entry. The durable queue position remains authoritative.
		return operationID, nil
	} else if !errors.Is(err, database.ErrNotFound) {
		return "", err
	}
	if retired, err := vm.db.Get(retiredKey(operationID)); err == nil {
		stored, retiredID, _, code, decodeErr := decodeDurableRetiredOperation(retired, maximum)
		if decodeErr != nil {
			return "", fmt.Errorf("decode retired operation: %w", decodeErr)
		}
		if retiredID != operationID || !bytes.Equal(stored.Bytes(), data) {
			return "", errors.New("retired operation differs from canonical operation ID")
		}
		return "", &refusalError{code: code}
	} else if !errors.Is(err, database.ErrNotFound) {
		return "", err
	}
	preferred, err := vm.getBlockLocked(vm.preferredID)
	if err != nil || preferred.postState == nil {
		return "", errInvalidPreference
	}
	preview := *preferred.postState
	for _, pendingID := range vm.pendingOrder {
		pending, exists := vm.pending[pendingID]
		if !exists {
			continue
		}
		result, err := vm.applyBoundedOperation(preview, pending)
		if err != nil {
			return "", fmt.Errorf("pending operation %s became invalid: %w", pendingID, err)
		}
		preview = result.State
	}
	if _, err := vm.applyBoundedOperation(preview, operation); err != nil {
		return "", err
	}
	if vm.durablePendingCount >= vm.genesis.Limits.MaxPendingOperations {
		return "", errors.New("durable pending queue reached genesis operation bound")
	}
	operationBytes := uint64(len(data))
	if vm.durablePendingBytes > vm.genesis.Limits.MaxPendingBytes ||
		operationBytes > vm.genesis.Limits.MaxPendingBytes-vm.durablePendingBytes {
		return "", errors.New("durable pending queue reached genesis byte bound")
	}
	if vm.nextPendingSequence == 0 || vm.nextPendingSequence == ^uint64(0) {
		return "", errors.New("durable pending sequence exhausted")
	}
	sequence := vm.nextPendingSequence
	persisted, err := encodeDurablePendingOperation(sequence, data)
	if err != nil {
		return "", err
	}
	if err := vm.db.Put(pendingKey(operationID), persisted); err != nil {
		return "", fmt.Errorf("persist pending operation: %w", err)
	}
	vm.pending[operationID] = operation
	vm.pendingSequence[operationID] = sequence
	vm.pendingOrder = append(vm.pendingOrder, operationID)
	vm.durablePendingCount++
	vm.durablePendingBytes += operationBytes
	vm.nextPendingSequence++
	vm.resetPendingRetryLocked()
	vm.signalPendingLocked()
	return operationID, nil
}

func (vm *VM) verifyBlock(_ context.Context, block *vmBlock) (err error) {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return errNotInitialized
	}
	if err := vm.requireCanonicalBlockReferenceLocked(block); err != nil {
		return err
	}
	if block.status == choices.Rejected {
		return errors.New("cannot verify a rejected block")
	}
	if block.builtLocally {
		defer func() {
			if err == nil {
				return
			}
			block.postState, block.receipts = nil, nil
			if cached, exists := vm.blocks[block.id]; exists && cached == block {
				delete(vm.blocks, block.id)
			}
			if requeueErr := vm.requeueBlockOperationsLocked(block); requeueErr != nil {
				err = errors.Join(err, fmt.Errorf("requeue locally built block after verify failure: %w", requeueErr))
			}
			vm.armPendingRetryLocked()
		}()
	}
	if block.postState != nil {
		return nil
	}
	if time.Until(block.timestamp) > maxClockSkew {
		return errors.New("block timestamp exceeds maximum clock skew")
	}
	parent, err := vm.getBlockLocked(block.parentID)
	if err != nil {
		return err
	}
	if parent.status == choices.Rejected || parent.postState == nil {
		return errors.New("parent rejected or unverified")
	}
	if block.height != parent.height+1 || block.timestamp.Before(parent.timestamp) {
		return errors.New("block does not follow parent")
	}
	state := *parent.postState
	receipts := make([]*Receipt, 0, len(block.wire.Operations))
	for index, operation := range block.wire.Operations {
		result, err := vm.applyBoundedOperation(state, operation)
		if err != nil {
			return fmt.Errorf("operation %d (%s): %w", index, operation.ID(), err)
		}
		state = result.State
		receipt, err := receiptFor(operation, result, block.id.String(), block.height, uint16(index))
		if err != nil {
			return err
		}
		receipts = append(receipts, receipt)
	}
	if _, err := encodeBoundedState(state, vm.genesis.Limits); err != nil {
		return err
	}
	block.postState, block.receipts = &state, receipts
	block.status = choices.Processing
	vm.blocks[block.id] = block
	return nil
}

func (vm *VM) acceptBlock(block *vmBlock) (err error) {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return errNotInitialized
	}
	if err := vm.requireCanonicalBlockReferenceLocked(block); err != nil {
		return err
	}
	if block.status == choices.Accepted {
		return nil
	}
	if block.builtLocally {
		defer func() {
			if err == nil {
				return
			}
			if requeueErr := vm.requeueBlockOperationsLocked(block); requeueErr != nil {
				err = errors.Join(err, fmt.Errorf("requeue locally built block after accept failure: %w", requeueErr))
			}
			vm.armPendingRetryLocked()
		}()
	}
	if block.status == choices.Rejected || block.postState == nil {
		return errors.New("cannot accept rejected or unverified block")
	}
	if block.parentID != vm.lastAcceptedID {
		return errors.New("accepted block does not extend last accepted")
	}
	stateBytes, err := encodeBoundedState(*block.postState, vm.genesis.Limits)
	if err != nil {
		return err
	}
	vdb := versiondb.New(vm.db)
	for _, write := range []struct{ key, value []byte }{
		{blockKey(block.id), block.bytes}, {stateByBlockKey(block.id), stateBytes},
		{heightKey(block.height), block.id[:]}, {lastAcceptedKey, block.id[:]}, {currentStateKey, stateBytes},
	} {
		if err := vdb.Put(write.key, write.value); err != nil {
			return err
		}
	}
	accepted := make(map[string]struct{}, len(block.wire.Operations))
	for index, operation := range block.wire.Operations {
		operationID := operation.ID()
		receiptBytes, err := canonicalJSON(block.receipts[index])
		if err != nil {
			return err
		}
		if err := vdb.Put(operationKey(operationID), operation.Bytes()); err != nil {
			return err
		}
		if err := vdb.Put(receiptKey(operationID), receiptBytes); err != nil {
			return err
		}
		if err := vdb.Delete(pendingKey(operationID)); err != nil {
			return err
		}
		accepted[operationID] = struct{}{}
	}
	resolutions, pendingSnapshot, err := vm.pendingResolutionsAfterAcceptLocked(*block.postState, accepted)
	if err != nil {
		return err
	}
	for operationID, retired := range resolutions {
		if len(retired) > 0 {
			if err := vdb.Put(retiredKey(operationID), retired); err != nil {
				return err
			}
		}
		if err := vdb.Delete(pendingKey(operationID)); err != nil {
			return err
		}
	}
	if err := vdb.Commit(); err != nil {
		return fmt.Errorf("commit accepted Continuity block: %w", err)
	}
	for operationID := range accepted {
		delete(vm.pending, operationID)
		delete(vm.pendingSequence, operationID)
	}
	for operationID := range resolutions {
		delete(vm.pending, operationID)
		delete(vm.pendingSequence, operationID)
	}
	block.status = choices.Accepted
	vm.lastAcceptedID, vm.preferredID = block.id, block.id
	vm.blocks[block.id] = block
	vm.pruneAcceptedCacheLocked(block.id)
	vm.installPendingQueueSnapshotLocked(pendingSnapshot)
	vm.resetPendingRetryLocked()
	if len(vm.pendingOrder) == 0 {
		vm.drainNotifyLocked()
	} else {
		vm.signalPendingLocked()
	}
	return nil
}

func (vm *VM) rejectBlock(block *vmBlock) error {
	vm.lock.Lock()
	defer vm.lock.Unlock()
	if !vm.initialized {
		return errNotInitialized
	}
	if err := vm.requireCanonicalBlockReferenceLocked(block); err != nil {
		return err
	}
	if block.status == choices.Rejected {
		return nil
	}
	if block.status == choices.Accepted {
		return errors.New("cannot reject an accepted block")
	}
	if block.builtLocally {
		if err := vm.requeueBlockOperationsLocked(block); err != nil {
			return err
		}
	}
	block.status = choices.Rejected
	block.postState, block.receipts = nil, nil
	if cached, exists := vm.blocks[block.id]; exists && cached == block {
		delete(vm.blocks, block.id)
	}
	vm.resetPendingRetryLocked()
	if len(vm.pendingOrder) == 0 {
		vm.drainNotifyLocked()
	} else {
		vm.signalPendingLocked()
	}
	return nil
}

func (vm *VM) requireCanonicalBlockReferenceLocked(block *vmBlock) error {
	if block == nil {
		return errors.New("block reference is nil")
	}
	if canonical, exists := vm.blocks[block.id]; exists && canonical != block {
		return fmt.Errorf("%w: %s", errNonCanonicalBlock, block.id)
	}
	return nil
}

func (vm *VM) getBlockLocked(id ids.ID) (*vmBlock, error) {
	if block, exists := vm.blocks[id]; exists {
		return block, nil
	}
	block, err := vm.loadAcceptedBlockLocked(id)
	if err != nil {
		return nil, err
	}
	if id == vm.lastAcceptedID || id == vm.genesisBlock.id {
		vm.blocks[id] = block
	}
	return block, nil
}

func (vm *VM) pruneAcceptedCacheLocked(keep ids.ID) {
	for id, block := range vm.blocks {
		if id == keep || id == vm.genesisBlock.id {
			continue
		}
		if block.status == choices.Accepted {
			delete(vm.blocks, id)
		}
	}
}

func (vm *VM) loadAcceptedBlockLocked(id ids.ID) (*vmBlock, error) {
	blockBytes, err := vm.db.Get(blockKey(id))
	if err != nil {
		return nil, err
	}
	stateBytes, err := vm.db.Get(stateByBlockKey(id))
	if err != nil {
		return nil, err
	}
	if uint64(len(stateBytes)) > uint64(vm.genesis.Limits.MaxAcceptedStateBytes) {
		return nil, errors.New("stored accepted state exceeds genesis byte bound")
	}
	state := new(kernel.ChainState)
	if err := decodeCanonical(stateBytes, state); err != nil {
		return nil, err
	}
	if err := kernel.ValidateChainState(*state); err != nil {
		return nil, err
	}
	if _, err := encodeBoundedState(*state, vm.genesis.Limits); err != nil {
		return nil, err
	}
	if id == vm.genesisBlock.id {
		return newGenesisBlock(vm, vm.genesisBytes, *state), nil
	}
	block, err := blockFromBytes(vm, blockBytes, choices.Accepted)
	if err != nil {
		return nil, err
	}
	if block.id != id {
		return nil, errors.New("stored block bytes do not match requested block ID")
	}
	block.postState = state
	return block, nil
}

func (vm *VM) acceptedStateLocked() (kernel.ChainState, error) {
	block, err := vm.getBlockLocked(vm.lastAcceptedID)
	if err != nil || block.postState == nil {
		return kernel.ChainState{}, err
	}
	return *block.postState, nil
}

func (vm *VM) signalPendingLocked() {
	if len(vm.pendingOrder) == 0 {
		return
	}
	select {
	case vm.notify <- struct{}{}:
	default:
	}
}

func (vm *VM) armPendingRetryLocked() {
	if !vm.initialized || len(vm.pendingOrder) == 0 {
		vm.resetPendingRetryLocked()
		return
	}
	if vm.pendingRetryDelay <= 0 {
		vm.pendingRetryDelay = pendingRetryInitialDelay
	} else if vm.pendingRetryDelay >= pendingRetryMaximumDelay/2 {
		vm.pendingRetryDelay = pendingRetryMaximumDelay
	} else {
		vm.pendingRetryDelay *= 2
	}
	vm.pendingRetryAfter = time.Now().Add(vm.pendingRetryDelay)
	vm.drainNotifyLocked()
}

func (vm *VM) resetPendingRetryLocked() {
	vm.pendingRetryDelay = 0
	vm.pendingRetryAfter = time.Time{}
}

func (vm *VM) drainNotifyLocked() {
	select {
	case <-vm.notify:
	default:
	}
}

func stopAndDrainTimer(timer *time.Timer) {
	if timer == nil || timer.Stop() {
		return
	}
	select {
	case <-timer.C:
	default:
	}
}

func blockKey(id ids.ID) []byte        { return append([]byte("continuity/block/"), id[:]...) }
func stateByBlockKey(id ids.ID) []byte { return append([]byte("continuity/state/block/"), id[:]...) }
func heightKey(height uint64) []byte {
	key := make([]byte, len("continuity/height/")+8)
	copy(key, "continuity/height/")
	binary.BigEndian.PutUint64(key[len("continuity/height/"):], height)
	return key
}
func operationKey(id string) []byte { return []byte("continuity/operation/" + id) }
func pendingKey(id string) []byte   { return []byte("continuity/pending/" + id) }
func retiredKey(id string) []byte   { return []byte("continuity/retired/" + id) }
func receiptKey(id string) []byte   { return []byte("continuity/receipt/" + id) }

func (*VM) Connected(context.Context, ids.NodeID, *version.Application) error            { return nil }
func (*VM) Disconnected(context.Context, ids.NodeID) error                               { return nil }
func (*VM) AppRequest(context.Context, ids.NodeID, uint32, time.Time, []byte) error      { return nil }
func (*VM) AppRequestFailed(context.Context, ids.NodeID, uint32, *common.AppError) error { return nil }
func (*VM) AppResponse(context.Context, ids.NodeID, uint32, []byte) error                { return nil }
func (*VM) AppGossip(context.Context, ids.NodeID, []byte) error                          { return nil }
