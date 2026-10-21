package avalanchevm

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

type pendingQueueItem struct {
	id        string
	sequence  uint64
	operation *Operation
}

type pendingQueueSnapshot struct {
	pending             map[string]*Operation
	pendingSequence     map[string]uint64
	pendingOrder        []string
	durablePendingCount uint32
	durablePendingBytes uint64
	nextPendingSequence uint64
}

func buildPendingQueueSnapshot(items []pendingQueueItem, limits Limits) (*pendingQueueSnapshot, error) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].sequence != items[j].sequence {
			return items[i].sequence < items[j].sequence
		}
		return items[i].id < items[j].id
	})
	snapshot := &pendingQueueSnapshot{
		pending:             make(map[string]*Operation, len(items)),
		pendingSequence:     make(map[string]uint64, len(items)),
		pendingOrder:        make([]string, 0, len(items)),
		nextPendingSequence: 1,
	}
	seenSequences := make(map[uint64]string, len(items))
	maximumSequence := uint64(0)
	for _, item := range items {
		if item.id == "" || item.operation == nil || item.sequence == 0 || item.id != item.operation.ID() {
			return nil, errors.New("invalid pending queue item")
		}
		if _, duplicate := snapshot.pending[item.id]; duplicate {
			return nil, fmt.Errorf("duplicate pending operation %s", item.id)
		}
		if prior, duplicate := seenSequences[item.sequence]; duplicate {
			return nil, fmt.Errorf("pending operations %s and %s share sequence %d", prior, item.id, item.sequence)
		}
		if snapshot.durablePendingCount >= limits.MaxPendingOperations {
			return nil, errors.New("durable pending queue exceeds genesis operation bound")
		}
		operationBytes := uint64(len(item.operation.Bytes()))
		if snapshot.durablePendingBytes > limits.MaxPendingBytes ||
			operationBytes > limits.MaxPendingBytes-snapshot.durablePendingBytes {
			return nil, errors.New("durable pending queue exceeds genesis byte bound")
		}
		seenSequences[item.sequence] = item.id
		snapshot.pending[item.id] = item.operation
		snapshot.pendingSequence[item.id] = item.sequence
		snapshot.pendingOrder = append(snapshot.pendingOrder, item.id)
		snapshot.durablePendingCount++
		snapshot.durablePendingBytes += operationBytes
		if item.sequence > maximumSequence {
			maximumSequence = item.sequence
		}
	}
	if maximumSequence == ^uint64(0) {
		return nil, errors.New("durable pending sequence exhausted")
	}
	if maximumSequence > 0 {
		snapshot.nextPendingSequence = maximumSequence + 1
	}
	return snapshot, nil
}

func (vm *VM) installPendingQueueSnapshotLocked(snapshot *pendingQueueSnapshot) {
	vm.pending = snapshot.pending
	vm.pendingSequence = snapshot.pendingSequence
	vm.pendingOrder = snapshot.pendingOrder
	vm.durablePendingCount = snapshot.durablePendingCount
	vm.durablePendingBytes = snapshot.durablePendingBytes
	vm.nextPendingSequence = snapshot.nextPendingSequence
}

func encodeDurablePendingOperation(sequence uint64, operationBytes []byte) ([]byte, error) {
	if sequence == 0 {
		return nil, errors.New("durable pending sequence must be positive")
	}
	return canonicalJSON(durablePendingOperation{
		Schema: pendingOperationSchema, Version: 1, Sequence: sequence,
		OperationBytesHex: hex.EncodeToString(operationBytes),
	})
}

func decodeDurablePendingOperation(data []byte, maximumOperationBytes uint32) (*Operation, string, uint64, error) {
	entry := new(durablePendingOperation)
	if err := decodeCanonical(data, entry); err != nil {
		return nil, "", 0, err
	}
	if entry.Schema != pendingOperationSchema || entry.Version != 1 || entry.Sequence == 0 {
		return nil, "", 0, errors.New("invalid durable pending operation envelope")
	}
	operationBytes, err := decodeOperationHex(entry.OperationBytesHex, maximumOperationBytes)
	if err != nil {
		return nil, "", 0, err
	}
	operation, operationID, err := parseOperation(operationBytes)
	if err != nil {
		return nil, "", 0, err
	}
	return operation, operationID, entry.Sequence, nil
}

func encodeDurableRetiredOperation(sequence uint64, operationBytes []byte, code kernel.Code) ([]byte, error) {
	if sequence == 0 || code == "" {
		return nil, errors.New("retired operation requires sequence and refusal code")
	}
	return canonicalJSON(durableRetiredOperation{
		Schema: retiredOperationSchema, Version: 1, Sequence: sequence,
		OperationBytesHex: hex.EncodeToString(operationBytes), Code: code,
	})
}

func decodeDurableRetiredOperation(data []byte, maximumOperationBytes uint32) (*Operation, string, uint64, kernel.Code, error) {
	entry := new(durableRetiredOperation)
	if err := decodeCanonical(data, entry); err != nil {
		return nil, "", 0, "", err
	}
	if entry.Schema != retiredOperationSchema || entry.Version != 1 || entry.Sequence == 0 || entry.Code == "" {
		return nil, "", 0, "", errors.New("invalid durable retired operation envelope")
	}
	operationBytes, err := decodeOperationHex(entry.OperationBytesHex, maximumOperationBytes)
	if err != nil {
		return nil, "", 0, "", err
	}
	operation, operationID, err := parseOperation(operationBytes)
	if err != nil {
		return nil, "", 0, "", err
	}
	return operation, operationID, entry.Sequence, entry.Code, nil
}

func decodeOperationHex(value string, maximum uint32) ([]byte, error) {
	if len(value)%2 != 0 || uint64(len(value)/2) > uint64(maximum) {
		return nil, errors.New("operation_bytes_hex is malformed or exceeds compiled bounds")
	}
	decoded, err := hex.DecodeString(value)
	if err != nil || hex.EncodeToString(decoded) != value {
		return nil, errors.New("operation_bytes_hex is not canonical lowercase hexadecimal")
	}
	return decoded, nil
}

func (vm *VM) replayAcceptedHistoryLocked(_ context.Context, lastAcceptedID ids.ID, initial kernel.ChainState) (*vmBlock, error) {
	state := initial
	genesisStateBytes, err := encodeBoundedState(state, vm.genesis.Limits)
	if err != nil {
		return nil, err
	}
	storedGenesisBlock, err := vm.db.Get(blockKey(vm.genesisBlock.id))
	if err != nil || !bytes.Equal(storedGenesisBlock, vm.genesisBytes) {
		return nil, errors.New("stored genesis block bytes mismatch")
	}
	storedGenesisState, err := vm.db.Get(stateByBlockKey(vm.genesisBlock.id))
	if err != nil || !bytes.Equal(storedGenesisState, genesisStateBytes) {
		return nil, errors.New("stored genesis state bytes mismatch")
	}
	heightZero, err := vm.db.Get(heightKey(0))
	if err != nil || !bytes.Equal(heightZero, vm.genesisBlock.id[:]) {
		return nil, errors.New("height-zero index does not name genesis")
	}
	currentStateBytes, err := vm.db.Get(currentStateKey)
	if err != nil {
		return nil, err
	}
	if uint64(len(currentStateBytes)) > uint64(vm.genesis.Limits.MaxAcceptedStateBytes) {
		return nil, errors.New("stored current state exceeds genesis byte bound")
	}
	if lastAcceptedID == vm.genesisBlock.id {
		if !bytes.Equal(currentStateBytes, genesisStateBytes) {
			return nil, errors.New("current state differs from replayed genesis state")
		}
		return newGenesisBlock(vm, vm.genesisBytes, state), nil
	}
	lastBytes, err := vm.db.Get(blockKey(lastAcceptedID))
	if err != nil {
		return nil, err
	}
	lastHeader, err := blockFromBytes(vm, lastBytes, choices.Accepted)
	if err != nil || lastHeader.id != lastAcceptedID {
		return nil, errors.New("last-accepted pointer does not match stored block")
	}
	parentID := vm.genesisBlock.id
	parentTimestamp := vm.genesisBlock.timestamp
	var replayed *vmBlock
	for height := uint64(1); height <= lastHeader.height; height++ {
		indexedBytes, err := vm.db.Get(heightKey(height))
		if err != nil {
			return nil, fmt.Errorf("read height %d index: %w", height, err)
		}
		indexedID, err := ids.ToID(indexedBytes)
		if err != nil {
			return nil, err
		}
		blockBytes, err := vm.db.Get(blockKey(indexedID))
		if err != nil {
			return nil, err
		}
		block, err := blockFromBytes(vm, blockBytes, choices.Accepted)
		if err != nil {
			return nil, err
		}
		if block.id != indexedID || block.parentID != parentID || block.height != height || block.timestamp.Before(parentTimestamp) {
			return nil, fmt.Errorf("height %d breaks accepted-chain continuity", height)
		}
		receipts := make([]*Receipt, 0, len(block.wire.Operations))
		for index, operation := range block.wire.Operations {
			operationID := operation.ID()
			storedOperation, err := vm.db.Get(operationKey(operationID))
			if err != nil || !bytes.Equal(storedOperation, operation.Bytes()) {
				return nil, fmt.Errorf("height %d operation %d persistence mismatch", height, index)
			}
			result, err := vm.applyBoundedOperation(state, operation)
			if err != nil {
				return nil, fmt.Errorf("height %d operation %d replay: %w", height, index, err)
			}
			state = result.State
			receipt, err := receiptFor(operation, result, block.id.String(), height, uint16(index))
			if err != nil {
				return nil, err
			}
			receiptBytes, err := canonicalJSON(receipt)
			if err != nil {
				return nil, err
			}
			storedReceipt, err := vm.db.Get(receiptKey(operationID))
			if err != nil || !bytes.Equal(storedReceipt, receiptBytes) {
				return nil, fmt.Errorf("height %d operation %d receipt mismatch", height, index)
			}
			receipts = append(receipts, receipt)
		}
		if err := kernel.ValidateChainState(state); err != nil {
			return nil, fmt.Errorf("height %d replayed state invalid: %w", height, err)
		}
		stateBytes, err := encodeBoundedState(state, vm.genesis.Limits)
		if err != nil {
			return nil, err
		}
		persistedState, err := vm.db.Get(stateByBlockKey(block.id))
		if err == nil && uint64(len(persistedState)) > uint64(vm.genesis.Limits.MaxAcceptedStateBytes) {
			return nil, fmt.Errorf("height %d stored state exceeds genesis byte bound", height)
		}
		if err != nil || !bytes.Equal(persistedState, stateBytes) {
			return nil, fmt.Errorf("height %d state differs from deterministic replay", height)
		}
		postState := state
		block.postState, block.receipts = &postState, receipts
		replayed = block
		parentID, parentTimestamp = block.id, block.timestamp
	}
	if replayed == nil || replayed.id != lastAcceptedID {
		return nil, errors.New("last-accepted pointer is not terminal height index")
	}
	finalStateBytes, err := encodeBoundedState(*replayed.postState, vm.genesis.Limits)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(currentStateBytes, finalStateBytes) {
		return nil, errors.New("current state differs from deterministic accepted-history replay")
	}
	return replayed, nil
}

func (vm *VM) loadPendingOperationsLocked() error {
	loaded := make([]pendingQueueItem, 0)
	stale := make([]string, 0)
	iterator := vm.db.NewIteratorWithPrefix(pendingPrefix)
	defer iterator.Release()
	for iterator.Next() {
		operationID := string(iterator.Key()[len(pendingPrefix):])
		operation, derivedID, sequence, err := decodeDurablePendingOperation(bytes.Clone(iterator.Value()), vm.genesis.Limits.MaxOperationBytes)
		if err != nil {
			return fmt.Errorf("decode pending operation %s: %w", operationID, err)
		}
		if operationID != derivedID {
			return errors.New("pending key differs from canonical operation ID")
		}
		if _, err := vm.db.Get(operationKey(operationID)); err == nil {
			stale = append(stale, operationID)
			continue
		} else if !errors.Is(err, database.ErrNotFound) {
			return err
		}
		if _, err := vm.db.Get(retiredKey(operationID)); err == nil {
			stale = append(stale, operationID)
			continue
		} else if !errors.Is(err, database.ErrNotFound) {
			return err
		}
		loaded = append(loaded, pendingQueueItem{id: operationID, sequence: sequence, operation: operation})
	}
	if err := iterator.Error(); err != nil {
		return err
	}
	snapshot, err := buildPendingQueueSnapshot(loaded, vm.genesis.Limits)
	if err != nil {
		return err
	}
	if len(stale) > 0 {
		vdb := versiondb.New(vm.db)
		for _, operationID := range stale {
			if err := vdb.Delete(pendingKey(operationID)); err != nil {
				return err
			}
		}
		if err := vdb.Commit(); err != nil {
			return err
		}
	}
	vm.installPendingQueueSnapshotLocked(snapshot)
	return nil
}

func (vm *VM) pendingResolutionsAfterAcceptLocked(postState kernel.ChainState, accepted map[string]struct{}) (map[string][]byte, *pendingQueueSnapshot, error) {
	items := make([]pendingQueueItem, 0)
	iterator := vm.db.NewIteratorWithPrefix(pendingPrefix)
	defer iterator.Release()
	for iterator.Next() {
		id := string(iterator.Key()[len(pendingPrefix):])
		op, derivedID, sequence, err := decodeDurablePendingOperation(bytes.Clone(iterator.Value()), vm.genesis.Limits.MaxOperationBytes)
		if err != nil || id != derivedID {
			return nil, nil, fmt.Errorf("invalid durable pending operation %s", id)
		}
		items = append(items, pendingQueueItem{id: id, sequence: sequence, operation: op})
	}
	if err := iterator.Error(); err != nil {
		return nil, nil, err
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].sequence != items[j].sequence {
			return items[i].sequence < items[j].sequence
		}
		return items[i].id < items[j].id
	})
	resolutions := make(map[string][]byte)
	retained := make([]pendingQueueItem, 0, len(items))
	retainDependentSuffix := false
	for _, item := range items {
		raw := item.operation.Bytes()
		if _, ok := accepted[item.id]; ok {
			resolutions[item.id] = nil
			continue
		}
		if acceptedBytes, err := vm.db.Get(operationKey(item.id)); err == nil {
			if !bytes.Equal(acceptedBytes, raw) {
				return nil, nil, fmt.Errorf("accepted operation %s differs from colliding pending bytes", item.id)
			}
			resolutions[item.id] = nil
			continue
		} else if !errors.Is(err, database.ErrNotFound) {
			return nil, nil, err
		}
		if retainDependentSuffix {
			retained = append(retained, item)
			continue
		}
		_, err := vm.applyBoundedOperation(postState, item.operation)
		if err == nil {
			retainDependentSuffix = true
			retained = append(retained, item)
			continue
		}
		code, ok := refusalCode(err)
		if !ok {
			return nil, nil, err
		}
		retired, err := encodeDurableRetiredOperation(item.sequence, raw, code)
		if err != nil {
			return nil, nil, err
		}
		resolutions[item.id] = retired
	}
	snapshot, err := buildPendingQueueSnapshot(retained, vm.genesis.Limits)
	if err != nil {
		return nil, nil, err
	}
	return resolutions, snapshot, nil
}

func (vm *VM) requeueBlockOperationsLocked(block *vmBlock) error {
	for _, operation := range block.wire.Operations {
		operationID := operation.ID()
		operationBytes := operation.Bytes()
		if acceptedBytes, err := vm.db.Get(operationKey(operationID)); err == nil {
			if !bytes.Equal(acceptedBytes, operationBytes) {
				return fmt.Errorf("accepted operation %s differs from rejected local block bytes", operationID)
			}
			if err := vm.purgeAcceptedPendingCollisionLocked(operationID, operationBytes); err != nil {
				return err
			}
			continue
		} else if !errors.Is(err, database.ErrNotFound) {
			return err
		}
		if retiredBytes, err := vm.db.Get(retiredKey(operationID)); err == nil {
			stored, retiredID, _, _, decodeErr := decodeDurableRetiredOperation(retiredBytes, vm.genesis.Limits.MaxOperationBytes)
			if decodeErr != nil {
				return decodeErr
			}
			if retiredID != operationID || !bytes.Equal(stored.Bytes(), operationBytes) {
				return errors.New("retired operation differs from rejected local block operation")
			}
			delete(vm.pending, operationID)
			delete(vm.pendingSequence, operationID)
			continue
		} else if !errors.Is(err, database.ErrNotFound) {
			return err
		}
		persisted, err := vm.db.Get(pendingKey(operationID))
		if err != nil {
			return fmt.Errorf("locally built operation %s lost durable queue entry: %w", operationID, err)
		}
		stored, derivedID, sequence, err := decodeDurablePendingOperation(persisted, vm.genesis.Limits.MaxOperationBytes)
		if err != nil || derivedID != operationID || !bytes.Equal(stored.Bytes(), operationBytes) {
			return errors.New("requeued operation differs from durable queue entry")
		}
		vm.pending[operationID] = operation
		vm.pendingSequence[operationID] = sequence
	}
	vm.rebuildPendingOrderLocked()
	return nil
}

func (vm *VM) purgeAcceptedPendingCollisionLocked(operationID string, operationBytes []byte) error {
	acceptedBytes, err := vm.db.Get(operationKey(operationID))
	if err != nil {
		return err
	}
	if !bytes.Equal(acceptedBytes, operationBytes) {
		return errors.New("accepted and pending operation bytes differ for one canonical ID")
	}
	vdb := versiondb.New(vm.db)
	if err := vdb.Delete(pendingKey(operationID)); err != nil {
		return err
	}
	if err := vdb.Commit(); err != nil {
		return fmt.Errorf("purge accepted pending collision: %w", err)
	}
	delete(vm.pending, operationID)
	delete(vm.pendingSequence, operationID)
	return vm.loadPendingOperationsLocked()
}

func (vm *VM) rebuildPendingOrderLocked() {
	vm.pendingOrder = vm.pendingOrder[:0]
	for operationID := range vm.pending {
		vm.pendingOrder = append(vm.pendingOrder, operationID)
	}
	sort.Slice(vm.pendingOrder, func(i, j int) bool {
		left, right := vm.pendingSequence[vm.pendingOrder[i]], vm.pendingSequence[vm.pendingOrder[j]]
		if left != right {
			return left < right
		}
		return vm.pendingOrder[i] < vm.pendingOrder[j]
	})
}
