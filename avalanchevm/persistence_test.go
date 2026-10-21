package avalanchevm

import (
	"context"
	"testing"

	"github.com/ava-labs/avalanchego/database/memdb"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

func TestPendingReconciliationRetainsDependentSuffix(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	first, err := NewOperation(ActionOpenRelation, signedOpen(t, "same-relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	second, err := NewOperation(ActionOpenRelation, signedOpen(t, "same-relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	for sequence, operation := range []*Operation{first, second} {
		entry, err := encodeDurablePendingOperation(uint64(sequence+1), operation.Bytes())
		if err != nil {
			t.Fatal(err)
		}
		if err := vm.db.Put(pendingKey(operation.ID()), entry); err != nil {
			t.Fatal(err)
		}
	}
	vm.lock.Lock()
	resolutions, snapshot, err := vm.pendingResolutionsAfterAcceptLocked(state, map[string]struct{}{})
	vm.lock.Unlock()
	if err != nil {
		t.Fatal(err)
	}
	if len(resolutions) != 0 {
		t.Fatal("a speculative pending predecessor unlawfully retired its dependent suffix")
	}
	if len(snapshot.pendingOrder) != 2 {
		t.Fatal("dependent pending suffix was not retained in the post-accept queue snapshot")
	}
}

func TestRejectIgnoresLawfullyRetiredLocalOperation(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "retired", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	retired, err := encodeDurableRetiredOperation(1, operation.Bytes(), kernel.CodeRelationAlreadyExists)
	if err != nil {
		t.Fatal(err)
	}
	if err := vm.db.Put(retiredKey(operation.ID()), retired); err != nil {
		t.Fatal(err)
	}
	vm.pending[operation.ID()] = operation
	vm.pendingSequence[operation.ID()] = 1
	vm.pendingOrder = []string{operation.ID()}
	block := &vmBlock{wire: &WireBlock{Operations: []*Operation{operation}}}
	vm.lock.Lock()
	err = vm.requeueBlockOperationsLocked(block)
	vm.lock.Unlock()
	if err != nil {
		t.Fatal(err)
	}
	if len(vm.pendingOrder) != 0 {
		t.Fatal("retired local operation was resurrected on reject")
	}
}

func TestAcceptedStateBoundsAreExecutable(t *testing.T) {
	state := testChainState(t)
	limits := fixedLimits()
	limits.MaxAcceptedStateBytes = 1
	if _, err := encodeBoundedState(state, limits); err == nil {
		t.Fatal("accepted state byte bound was not enforced")
	}
}
