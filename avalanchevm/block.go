package avalanchevm

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

const maxClockSkew = 10 * time.Second

type WireBlock struct {
	Schema     string       `json:"schema"`
	ParentID   string       `json:"parent_id"`
	Timestamp  int64        `json:"timestamp"`
	Height     uint64       `json:"height"`
	Operations []*Operation `json:"operations"`
}

type vmBlock struct {
	vm           *VM
	wire         *WireBlock
	id           ids.ID
	bytes        []byte
	parentID     ids.ID
	height       uint64
	timestamp    time.Time
	status       choices.Status
	postState    *kernel.ChainState
	receipts     []*Receipt
	genesis      bool
	builtLocally bool
}

func newGenesisBlock(vm *VM, genesisBytes []byte, state kernel.ChainState) *vmBlock {
	digest := sha256.Sum256(genesisBytes)
	cloned := state
	return &vmBlock{
		vm: vm, id: ids.ID(digest), bytes: bytes.Clone(genesisBytes), parentID: ids.Empty,
		height: 0, timestamp: time.Unix(0, 0).UTC(), status: choices.Accepted,
		postState: &cloned, genesis: true,
	}
}

func parseWireBlock(data []byte, maximumBlockBytes uint32, maximumOperations uint16, maximumOperationBytes uint32) (*WireBlock, ids.ID, ids.ID, error) {
	if len(data) == 0 || uint64(len(data)) > uint64(maximumBlockBytes) {
		return nil, ids.Empty, ids.Empty, errors.New("block byte length is outside compiled bounds")
	}
	wire := new(WireBlock)
	if err := decodeCanonical(data, wire); err != nil {
		return nil, ids.Empty, ids.Empty, err
	}
	if wire.Schema != BlockSchema {
		return nil, ids.Empty, ids.Empty, errors.New("unsupported block schema")
	}
	parentID, err := ids.FromString(wire.ParentID)
	if err != nil {
		return nil, ids.Empty, ids.Empty, fmt.Errorf("parent_id: %w", err)
	}
	if wire.Timestamp < 0 || wire.Height == 0 {
		return nil, ids.Empty, ids.Empty, errors.New("non-genesis block requires a non-negative timestamp and positive height")
	}
	if len(wire.Operations) == 0 || len(wire.Operations) > int(maximumOperations) {
		return nil, ids.Empty, ids.Empty, fmt.Errorf("block must contain 1..%d operations", maximumOperations)
	}
	seen := make(map[string]struct{}, len(wire.Operations))
	for index, operation := range wire.Operations {
		if operation == nil {
			return nil, ids.Empty, ids.Empty, fmt.Errorf("operation %d is null", index)
		}
		operationBytes := operation.Bytes()
		if uint64(len(operationBytes)) > uint64(maximumOperationBytes) {
			return nil, ids.Empty, ids.Empty, fmt.Errorf("operation %d exceeds genesis byte limit", index)
		}
		parsed, operationID, err := parseOperation(operationBytes)
		if err != nil {
			return nil, ids.Empty, ids.Empty, fmt.Errorf("operation %d: %w", index, err)
		}
		wire.Operations[index] = parsed
		if _, duplicate := seen[operationID]; duplicate {
			return nil, ids.Empty, ids.Empty, fmt.Errorf("duplicate operation %s", operationID)
		}
		seen[operationID] = struct{}{}
	}
	digest := sha256.Sum256(data)
	return wire, ids.ID(digest), parentID, nil
}

func newWireBlock(vm *VM, wire *WireBlock, builtLocally bool) (*vmBlock, error) {
	data, err := canonicalJSON(wire)
	if err != nil {
		return nil, err
	}
	parsed, blockID, parentID, err := parseWireBlock(data, vm.genesis.Limits.MaxBlockBytes, vm.genesis.Limits.MaxOperationsPerBlock, vm.genesis.Limits.MaxOperationBytes)
	if err != nil {
		return nil, err
	}
	return &vmBlock{
		vm: vm, wire: parsed, id: blockID, bytes: data, parentID: parentID,
		height: parsed.Height, timestamp: time.Unix(parsed.Timestamp, 0).UTC(),
		status: choices.Processing, builtLocally: builtLocally,
	}, nil
}

func blockFromBytes(vm *VM, data []byte, status choices.Status) (*vmBlock, error) {
	wire, blockID, parentID, err := parseWireBlock(data, vm.genesis.Limits.MaxBlockBytes, vm.genesis.Limits.MaxOperationsPerBlock, vm.genesis.Limits.MaxOperationBytes)
	if err != nil {
		return nil, err
	}
	return &vmBlock{
		vm: vm, wire: wire, id: blockID, bytes: bytes.Clone(data), parentID: parentID,
		height: wire.Height, timestamp: time.Unix(wire.Timestamp, 0).UTC(), status: status,
	}, nil
}

func (block *vmBlock) ID() ids.ID           { return block.id }
func (block *vmBlock) Parent() ids.ID       { return block.parentID }
func (block *vmBlock) Height() uint64       { return block.height }
func (block *vmBlock) Timestamp() time.Time { return block.timestamp }
func (block *vmBlock) Bytes() []byte        { return bytes.Clone(block.bytes) }

func (block *vmBlock) Status() choices.Status {
	block.vm.lock.RLock()
	defer block.vm.lock.RUnlock()
	return block.status
}

func (block *vmBlock) Verify(ctx context.Context) error {
	if block.genesis {
		return nil
	}
	return block.vm.verifyBlock(ctx, block)
}

func (block *vmBlock) Accept(context.Context) error {
	if block.genesis {
		return nil
	}
	return block.vm.acceptBlock(block)
}

func (block *vmBlock) Reject(context.Context) error {
	if block.genesis {
		return errCannotRejectGenesis
	}
	return block.vm.rejectBlock(block)
}
