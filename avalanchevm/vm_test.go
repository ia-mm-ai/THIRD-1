package avalanchevm

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/database/memdb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

func testChainContext() *snow.Context {
	return &snow.Context{NetworkID: testNetworkID, SubnetID: testSubnetID, ChainID: testChainID}
}

func testGenesisBytesForSubnet(t *testing.T, subnetID ids.ID) []byte {
	t.Helper()
	semanticProfile := testHash("profile")
	data, err := canonicalJSON(Genesis{
		Schema: GenesisSchema, Version: 1,
		SecurityDomain: AvalancheSecurityDomain{
			NetworkID: testNetworkID, SubnetID: subnetID.String(),
		},
		Config: GenesisKernelConfig{
			ProtocolVersion: kernel.ProtocolVersion,
			KernelVersion: kernel.Version,
			SemanticProfileCommitment: semanticProfile,
		},
		Limits: fixedLimits(),
	})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func testGenesisBytes(t *testing.T) []byte {
	return testGenesisBytesForSubnet(t, testSubnetID)
}

func TestAcceptedHistoryReplaysExactly(t *testing.T) {
	ctx := context.Background()
	database := memdb.New()
	genesisBytes := testGenesisBytes(t)
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), database, genesisBytes, nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := vm.SetState(ctx, snow.NormalOp); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := vm.IssueOperation(operation.Bytes()); err != nil {
		t.Fatal(err)
	}
	block, err := vm.BuildBlock(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err := block.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	if err := block.Accept(ctx); err != nil {
		t.Fatal(err)
	}
	vm.lock.Lock()
	accepted, err := vm.acceptedStateLocked()
	vm.lock.Unlock()
	if err != nil {
		t.Fatal(err)
	}
	expectedHash, err := kernel.ChainStateHash(accepted)
	if err != nil {
		t.Fatal(err)
	}

	restarted := new(VM)
	if err := restarted.Initialize(ctx, testChainContext(), database, genesisBytes, nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	restarted.lock.Lock()
	replayed, err := restarted.acceptedStateLocked()
	restarted.lock.Unlock()
	if err != nil {
		t.Fatal(err)
	}
	replayedHash, err := kernel.ChainStateHash(replayed)
	if err != nil {
		t.Fatal(err)
	}
	if replayedHash != expectedHash {
		t.Fatalf("accepted-history replay changed state hash: %s != %s", replayedHash, expectedHash)
	}
}

func TestOperationAPIEnforcesByteCapBeforeDecode(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/operations", strings.NewReader(strings.Repeat("x", int(maxOperationBytes)+1)))
	recorder := httptest.NewRecorder()
	vm.handleIssueOperation(recorder, request)
	if recorder.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("oversized request returned %d", recorder.Code)
	}
}

func TestInitializeRejectsOperatorConfigBytes(t *testing.T) {
	vm := new(VM)
	err := vm.Initialize(context.Background(), testChainContext(), memdb.New(), testGenesisBytes(t), nil, []byte(`{}`), nil, nil)
	if err == nil || !strings.Contains(err.Error(), "operator config") {
		t.Fatalf("non-empty operator config was not rejected: %v", err)
	}
}

func TestGenesisSecurityDomainPreventsCrossL1Replay(t *testing.T) {
	otherSubnetID := testSubnetID
	otherSubnetID[0] ^= 0xff
	otherChainID := testChainID
	otherChainID[0] ^= 0xff
	semanticProfile := testHash("profile")
	profileA, err := deriveChainProfileCommitment(semanticProfile, testNetworkID, testSubnetID, testChainID, kernel.Version)
	if err != nil {
		t.Fatal(err)
	}
	profileOtherSubnet, err := deriveChainProfileCommitment(semanticProfile, testNetworkID, otherSubnetID, testChainID, kernel.Version)
	if err != nil {
		t.Fatal(err)
	}
	if profileA == profileOtherSubnet {
		t.Fatal("distinct Avalanche subnet/L1 IDs produced the same effective profile commitment")
	}
	profileB, err := deriveChainProfileCommitment(semanticProfile, testNetworkID, testSubnetID, otherChainID, kernel.Version)
	if err != nil {
		t.Fatal(err)
	}
	if profileA == profileB {
		t.Fatal("distinct Avalanche chain IDs in one subnet/L1 produced the same effective profile commitment")
	}

	wrongContext := &snow.Context{NetworkID: testNetworkID, SubnetID: otherSubnetID, ChainID: testChainID}
	vm := new(VM)
	err = vm.Initialize(context.Background(), wrongContext, memdb.New(), testGenesisBytes(t), nil, nil, nil, nil)
	if err == nil || !strings.Contains(err.Error(), "subnet/L1 ID") {
		t.Fatalf("genesis was admitted under a different Avalanche subnet/L1: %v", err)
	}
	chainBContext := &snow.Context{NetworkID: testNetworkID, SubnetID: testSubnetID, ChainID: otherChainID}
	chainBVM := new(VM)
	if err := chainBVM.Initialize(context.Background(), chainBContext, memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	chainBVM.lock.Lock()
	chainB, err := chainBVM.acceptedStateLocked()
	chainBVM.lock.Unlock()
	if err != nil {
		t.Fatal(err)
	}
	if chainB.ProfileCommitment != profileB {
		t.Fatal("Initialize did not derive the expected chain-instance profile commitment")
	}

	replayed, err := NewOperation(ActionOpenRelation, signedOpen(t, "cross-l1-replay", profileA))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := applyOperation(chainB, replayed); err == nil {
		t.Fatal("operation signed for one L1 was accepted by a distinct L1")
	} else if code, ok := refusalCode(err); !ok || code != kernel.CodeProfileMismatch {
		t.Fatalf("cross-L1 replay refusal = %q, %v; want PROFILE_MISMATCH", code, err)
	}
}

func TestParseWireBlockRejectsOversizeBeforeDecode(t *testing.T) {
	data := make([]byte, int(maxBlockBytes)+1)
	if _, _, _, err := parseWireBlock(data, maxBlockBytes, maxOperationsPerBlock, maxOperationBytes); err == nil {
		t.Fatal("oversized block reached decoder")
	}
}

func TestUnverifiedBlocksAreNotCachedAndRejectedBlocksAreEvicted(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	vm.lock.RLock()
	genesisID := vm.genesisBlock.id
	vm.lock.RUnlock()
	parse := func(label string) snowman.Block {
		t.Helper()
		operation, err := NewOperation(ActionOpenRelation, signedOpen(t, label, state.ProfileCommitment))
		if err != nil {
			t.Fatal(err)
		}
		data, err := canonicalJSON(WireBlock{
			Schema: BlockSchema, ParentID: genesisID.String(), Timestamp: time.Now().UTC().Unix(),
			Height: 1, Operations: []*Operation{operation},
		})
		if err != nil {
			t.Fatal(err)
		}
		block, err := vm.ParseBlock(ctx, data)
		if err != nil {
			t.Fatal(err)
		}
		return block
	}
	first := parse("processing-first")
	second := parse("processing-second")
	vm.lock.RLock()
	_, firstCachedBeforeVerify := vm.blocks[first.ID()]
	_, secondCachedBeforeVerify := vm.blocks[second.ID()]
	vm.lock.RUnlock()
	if firstCachedBeforeVerify || secondCachedBeforeVerify {
		t.Fatal("ParseBlock retained unverified adversarial blocks")
	}
	if err := first.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	if err := first.Reject(ctx); err != nil {
		t.Fatal(err)
	}
	vm.lock.RLock()
	_, firstCachedAfterReject := vm.blocks[first.ID()]
	vm.lock.RUnlock()
	if firstCachedAfterReject {
		t.Fatal("rejected block remained in the processing cache")
	}
	if err := second.Verify(ctx); err != nil {
		t.Fatalf("independent block failed after rejected-cache eviction: %v", err)
	}
}

func TestPendingRetryBackoffAndCausalReset(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := vm.SetState(ctx, snow.NormalOp); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "retry", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := vm.IssueOperation(operation.Bytes()); err != nil {
		t.Fatal(err)
	}
	vm.lock.Lock()
	vm.armPendingRetryLocked()
	deadline := vm.pendingRetryAfter
	vm.lock.Unlock()
	if deadline.IsZero() {
		t.Fatal("pending retry deadline was not armed")
	}
	waitCtx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	if _, err := vm.WaitForEvent(waitCtx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("backoff did not suppress immediate PendingTxs: %v", err)
	}
	vm.lock.Lock()
	vm.resetPendingRetryLocked()
	vm.signalPendingLocked()
	vm.lock.Unlock()
	message, err := vm.WaitForEvent(ctx)
	if err != nil || message != common.PendingTxs {
		t.Fatalf("causal reset did not restore PendingTxs: %v, %v", message, err)
	}
}

func TestCompetingAcceptReloadsDurableLocalOperation(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := vm.SetState(ctx, snow.NormalOp); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	localOperation, err := NewOperation(ActionOpenRelation, signedOpen(t, "local", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := vm.IssueOperation(localOperation.Bytes()); err != nil {
		t.Fatal(err)
	}
	localBlock, err := vm.BuildBlock(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(vm.pendingOrder) != 0 {
		t.Fatal("locally selected operation remained in the live queue")
	}
	externalOperation, err := NewOperation(ActionOpenRelation, signedOpen(t, "external", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	vm.lock.RLock()
	genesisID := vm.genesisBlock.id
	vm.lock.RUnlock()
	wireBytes, err := canonicalJSON(WireBlock{
		Schema: BlockSchema, ParentID: genesisID.String(), Timestamp: time.Now().UTC().Unix(),
		Height: 1, Operations: []*Operation{externalOperation},
	})
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := vm.ParseBlock(ctx, wireBytes)
	if err != nil {
		t.Fatal(err)
	}
	if err := parsed.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	if err := parsed.Accept(ctx); err != nil {
		t.Fatal(err)
	}
	vm.lock.RLock()
	_, reloaded := vm.pending[localOperation.ID()]
	orderLength := len(vm.pendingOrder)
	vm.lock.RUnlock()
	if !reloaded || orderLength != 1 {
		t.Fatal("competing accept did not rehydrate the still-valid durable operation")
	}
	if err := localBlock.Reject(ctx); err != nil {
		t.Fatal(err)
	}
	vm.lock.RLock()
	defer vm.lock.RUnlock()
	if len(vm.pendingOrder) != 1 || vm.pendingOrder[0] != localOperation.ID() {
		t.Fatal("reject duplicated or lost the rehydrated operation")
	}
}

func TestPendingQueueCapsAndRestartAccounting(t *testing.T) {
	ctx := context.Background()
	database := memdb.New()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), database, testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := vm.SetState(ctx, snow.NormalOp); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "bounded", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := vm.IssueOperation(operation.Bytes()); err != nil {
		t.Fatal(err)
	}
	if vm.durablePendingCount != 1 || vm.durablePendingBytes != uint64(len(operation.Bytes())) {
		t.Fatal("pending queue accounting did not track durable admission")
	}
	restarted := new(VM)
	if err := restarted.Initialize(ctx, testChainContext(), database, testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if restarted.durablePendingCount != 1 || restarted.durablePendingBytes != uint64(len(operation.Bytes())) {
		t.Fatal("restart did not reconstruct pending queue accounting")
	}
	restarted.durablePendingCount = restarted.genesis.Limits.MaxPendingOperations
	second, err := NewOperation(ActionOpenRelation, signedOpen(t, "bounded-second", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := restarted.IssueOperation(second.Bytes()); err == nil || !strings.Contains(err.Error(), "operation bound") {
		t.Fatalf("pending operation cap was not enforced: %v", err)
	}
}

func TestIssueRefusesStateThatWouldCrossAdapterBound(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	vm.genesis.Limits.MaxRelations = 0
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "over-bound", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := vm.IssueOperation(operation.Bytes()); err == nil {
		t.Fatal("operation crossing the adapter state bound was durably admitted")
	} else if code, ok := refusalCode(err); !ok || code != codeStateBoundExceeded {
		t.Fatalf("state-bound refusal code = %q, %v", code, err)
	}
	if vm.durablePendingCount != 0 || len(vm.pendingOrder) != 0 {
		t.Fatal("state-bound refusal mutated the pending queue")
	}
}

func TestAcceptedAncestorsDoNotRetainFullStateSnapshots(t *testing.T) {
	ctx := context.Background()
	vm := new(VM)
	if err := vm.Initialize(ctx, testChainContext(), memdb.New(), testGenesisBytes(t), nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := vm.SetState(ctx, snow.NormalOp); err != nil {
		t.Fatal(err)
	}
	state := testChainState(t)
	var firstID string
	for index, label := range []string{"cache-first", "cache-second"} {
		operation, err := NewOperation(ActionOpenRelation, signedOpen(t, label, state.ProfileCommitment))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := vm.IssueOperation(operation.Bytes()); err != nil {
			t.Fatal(err)
		}
		block, err := vm.BuildBlock(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if err := block.Verify(ctx); err != nil {
			t.Fatal(err)
		}
		if err := block.Accept(ctx); err != nil {
			t.Fatal(err)
		}
		if index == 0 {
			firstID = block.ID().String()
		}
	}
	vm.lock.RLock()
	for id, block := range vm.blocks {
		if id.String() == firstID && block.postState != nil {
			vm.lock.RUnlock()
			t.Fatal("accepted ancestor retained a full state snapshot in the live cache")
		}
	}
	vm.lock.RUnlock()
}
