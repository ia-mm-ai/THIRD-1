package kernel

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"os"
	"reflect"
	"sort"
	"testing"
)

type vectorOperation struct {
	Kind  string          `json:"kind"`
	Value json.RawMessage `json:"value"`
}

type vectorExpected struct {
	Accepted          bool       `json:"accepted"`
	Code              Code       `json:"code"`
	OperationHash     *string    `json:"operation_hash"`
	RelationID        string     `json:"relation_id"`
	RelationStateHash *string    `json:"relation_state_hash"`
	ClosureHash       *string    `json:"closure_hash"`
	State             ChainState `json:"state"`
}

type vectorCase struct {
	ID                   string          `json:"id"`
	InitialState         ChainState      `json:"initial_state"`
	Operation            vectorOperation `json:"operation"`
	CanonicalUnsignedHex string          `json:"canonical_unsigned_hex"`
	Expected             vectorExpected  `json:"expected"`
}

type vectorStep struct {
	Operation            vectorOperation `json:"operation"`
	CanonicalUnsignedHex string          `json:"canonical_unsigned_hex"`
	Expected             vectorExpected  `json:"expected"`
}

type vectorSuite struct {
	SuiteHash  string       `json:"suite_hash"`
	Operations []vectorCase `json:"operations"`
	Lifecycle  struct {
		InitialState ChainState   `json:"initial_state"`
		Steps        []vectorStep `json:"steps"`
		FinalState   ChainState   `json:"final_state"`
	} `json:"lifecycle_sequence"`
}

func loadVectorSuite(t *testing.T) vectorSuite {
	t.Helper()
	encoded, err := os.ReadFile("../bridge/vectors/continuity_vectors_001.json")
	if err != nil {
		t.Fatalf("read bridge vectors: %v", err)
	}
	var suite vectorSuite
	if err := json.Unmarshal(encoded, &suite); err != nil {
		t.Fatalf("decode bridge vectors: %v", err)
	}
	const expectedSuiteHash = "ce74eea2aa34cc3ded66b9213de37f2ccdde3bc769a14c265e7c84fb53b7ad01"
	if suite.SuiteHash != expectedSuiteHash {
		t.Fatalf("bridge suite hash = %s, want %s", suite.SuiteHash, expectedSuiteHash)
	}
	const expectedFileHash = "8c0469d0e97f901c6ff3c88e09b98b9adf8439d093674d362a07f2c4695121b6"
	digest := sha256.Sum256(encoded)
	if got := hex.EncodeToString(digest[:]); got != expectedFileHash {
		t.Fatalf("bridge vector file hash = %s, want %s", got, expectedFileHash)
	}
	return suite
}

func decodeAndApply(t *testing.T, state ChainState, operation vectorOperation) (ApplyResult, []byte) {
	t.Helper()
	switch operation.Kind {
	case "OPEN_RELATION":
		var value OpenRelation
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			t.Fatalf("decode open: %v", err)
		}
		canonical, err := CanonicalOpenRelationBytes(value)
		if err != nil {
			t.Fatalf("canonical open: %v", err)
		}
		return ApplyOpen(state, value), canonical
	case "APPLY_CROSSING":
		var value CrossingProposal
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			t.Fatalf("decode crossing: %v", err)
		}
		canonical, err := CanonicalProposalBytes(value)
		if err != nil {
			t.Fatalf("canonical crossing: %v", err)
		}
		return ApplyCrossing(state, value), canonical
	case "ADMIT_SUCCESSOR":
		var value SuccessorAdmission
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			t.Fatalf("decode successor: %v", err)
		}
		canonical, err := CanonicalFreshAdmissionBytes(value.Admission)
		if err != nil {
			t.Fatalf("canonical successor: %v", err)
		}
		return ApplySuccessor(state, value), canonical
	case "WITHDRAW_CHAMBER":
		var value WithdrawChamber
		if err := json.Unmarshal(operation.Value, &value); err != nil {
			t.Fatalf("decode withdrawal: %v", err)
		}
		canonical, err := CanonicalWithdrawalBytes(value)
		if err != nil {
			t.Fatalf("canonical withdrawal: %v", err)
		}
		return ApplyWithdrawal(state, value), canonical
	default:
		t.Fatalf("unknown operation kind %q", operation.Kind)
		return ApplyResult{}, nil
	}
}

func expectedString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func assertVectorResult(t *testing.T, expected vectorExpected, actual ApplyResult) {
	t.Helper()
	if actual.Accepted != expected.Accepted || actual.Code != expected.Code {
		t.Fatalf("result = (%t, %s), want (%t, %s)", actual.Accepted, actual.Code, expected.Accepted, expected.Code)
	}
	if actual.RelationID != expected.RelationID {
		t.Fatalf("relation id = %s, want %s", actual.RelationID, expected.RelationID)
	}
	if actual.OperationHash != expectedString(expected.OperationHash) {
		t.Fatalf("operation hash = %s, want %s", actual.OperationHash, expectedString(expected.OperationHash))
	}
	if actual.RelationStateHash != expectedString(expected.RelationStateHash) {
		t.Fatalf("relation state hash = %s, want %s", actual.RelationStateHash, expectedString(expected.RelationStateHash))
	}
	if actual.ClosureHash != expectedString(expected.ClosureHash) {
		t.Fatalf("closure hash = %s, want %s", actual.ClosureHash, expectedString(expected.ClosureHash))
	}
	if !reflect.DeepEqual(actual.State, expected.State) {
		t.Fatalf("state mismatch\ngot:  %#v\nwant: %#v", actual.State, expected.State)
	}
}

func TestBridgeVectors(t *testing.T) {
	suite := loadVectorSuite(t)
	for _, vector := range suite.Operations {
		vector := vector
		t.Run(vector.ID, func(t *testing.T) {
			before := cloneChainState(vector.InitialState)
			result, canonical := decodeAndApply(t, vector.InitialState, vector.Operation)
			if got := hex.EncodeToString(canonical); got != vector.CanonicalUnsignedHex {
				t.Fatalf("canonical bytes = %s, want %s", got, vector.CanonicalUnsignedHex)
			}
			assertVectorResult(t, vector.Expected, result)
			if !reflect.DeepEqual(vector.InitialState, before) {
				t.Fatal("apply mutated its input state")
			}
			if !result.Accepted {
				second, _ := decodeAndApply(t, vector.InitialState, vector.Operation)
				if !reflect.DeepEqual(result, second) {
					t.Fatal("refusal is not deterministic")
				}
			}
		})
	}
}

func TestApplyResultWireUsesExplicitNulls(t *testing.T) {
	suite := loadVectorSuite(t)
	for _, id := range []string{"O001", "O003"} {
		var vector vectorCase
		for _, candidate := range suite.Operations {
			if candidate.ID == id {
				vector = candidate
				break
			}
		}
		result, _ := decodeAndApply(t, vector.InitialState, vector.Operation)
		encoded, err := json.Marshal(result)
		if err != nil {
			t.Fatalf("marshal %s result: %v", id, err)
		}
		var wire map[string]json.RawMessage
		if err := json.Unmarshal(encoded, &wire); err != nil {
			t.Fatalf("decode %s result wire: %v", id, err)
		}
		for _, field := range []string{"operation_hash", "relation_id", "relation_state_hash", "closure_hash"} {
			got, found := wire[field]
			if !found {
				t.Fatalf("%s result omitted %s", id, field)
			}
			var expected *string
			switch field {
			case "operation_hash":
				expected = vector.Expected.OperationHash
			case "relation_id":
				if vector.Expected.RelationID != "" {
					expected = &vector.Expected.RelationID
				}
			case "relation_state_hash":
				expected = vector.Expected.RelationStateHash
			case "closure_hash":
				expected = vector.Expected.ClosureHash
			}
			if expected == nil {
				if string(got) != "null" {
					t.Fatalf("%s %s = %s, want null", id, field, got)
				}
			} else {
				var actual string
				if err := json.Unmarshal(got, &actual); err != nil || actual != *expected {
					t.Fatalf("%s %s = %s, want %s", id, field, got, *expected)
				}
			}
		}
	}
}

func TestLifecycleSequence(t *testing.T) {
	suite := loadVectorSuite(t)
	state := suite.Lifecycle.InitialState
	for index, step := range suite.Lifecycle.Steps {
		result, canonical := decodeAndApply(t, state, step.Operation)
		if got := hex.EncodeToString(canonical); got != step.CanonicalUnsignedHex {
			t.Fatalf("step %d canonical bytes mismatch", index+1)
		}
		assertVectorResult(t, step.Expected, result)
		if !result.Accepted {
			t.Fatalf("step %d refused with %s", index+1, result.Code)
		}
		state = result.State
	}
	if !reflect.DeepEqual(state, suite.Lifecycle.FinalState) {
		t.Fatal("lifecycle final state mismatch")
	}
}

func TestDeclaredPrecedence(t *testing.T) {
	wantOpen := []Code{
		CodeMalformedChainState, CodeMalformedOpen, CodeUnsupportedVersion,
		CodeProfileMismatch, CodeRelationAlreadyExists, CodeVerifierScopesNotDistinct,
		CodeOpenSignatureAInvalid, CodeOpenSignatureBInvalid,
	}
	wantCrossing := []Code{
		CodeMalformedChainState, CodeMalformedProposal, CodeUnsupportedVersion,
		CodeRelationNotFound, CodeProfileMismatch, CodeRelationEnded, CodeStaleEpoch,
		CodeStaleHeadA, CodeStaleHeadB, CodeNonceAMismatch, CodeNonceBMismatch,
		CodeResolutionInvalid, CodeOutputRuleViolation, CodeHandoffRuleViolation,
		CodeSignatureAInvalid, CodeSignatureBInvalid,
	}
	wantSuccessor := []Code{
		CodeMalformedChainState, CodeMalformedSuccessor, CodeUnsupportedVersion,
		CodePredecessorClosureNotFound, CodePredecessorNotRenew,
		CodePredecessorAlreadyConsumed, CodeRelationAlreadyExists, CodeProfileMismatch,
		CodePredecessorCommitmentMismatch, CodeHandoffCommitmentMismatch,
		CodeFreshInputMismatch, CodeSuccessorAdmissionMismatch,
		CodeVerifierScopesNotDistinct, CodeFreshSignatureAInvalid, CodeFreshSignatureBInvalid,
	}
	wantWithdrawal := []Code{
		CodeMalformedChainState, CodeMalformedWithdrawal, CodeUnsupportedVersion,
		CodeRelationNotFound, CodeProfileMismatch, CodeRelationEnded, CodeStaleEpoch,
		CodeStaleHeadA, CodeStaleHeadB, CodeSideInvalid, CodeWithdrawNonceMismatch,
		CodeWithdrawSignatureInvalid,
	}
	for name, pair := range map[string]struct{ got, want []Code }{
		"open":      {OpenPrecedence(), wantOpen},
		"crossing":  {TransitionPrecedence(), wantCrossing},
		"successor": {SuccessorPrecedence(), wantSuccessor},
		"withdraw":  {WithdrawalPrecedence(), wantWithdrawal},
	} {
		if !reflect.DeepEqual(pair.got, pair.want) {
			t.Fatalf("%s precedence = %v, want %v", name, pair.got, pair.want)
		}
	}
	mutated := TransitionPrecedence()
	mutated[0] = CodeAccepted
	if TransitionPrecedence()[0] != CodeMalformedChainState {
		t.Fatal("caller mutated the frozen transition precedence")
	}
}

func TestCollisionPrecedence(t *testing.T) {
	suite := loadVectorSuite(t)
	byID := make(map[string]vectorCase, len(suite.Operations))
	for _, vector := range suite.Operations {
		byID[vector.ID] = vector
	}
	expect := func(name string, got, want Code) {
		t.Helper()
		if got != want {
			t.Fatalf("%s collision = %s, want %s", name, got, want)
		}
	}
	zeroHash := "0000000000000000000000000000000000000000000000000000000000000000"
	badSignature := base64.StdEncoding.EncodeToString(make([]byte, 64))

	var crossing CrossingProposal
	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.NonceA = 0
	crossing.Resolution = "X"
	expect("crossing nonce_a/resolution", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeNonceAMismatch)

	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.Version = ""
	expect("crossing empty version", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.Version = "999"
	crossing.RelationID = zeroHash
	expect("crossing version/relation", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.PrevHeadA = zeroHash
	crossing.PrevHeadB = zeroHash
	expect("crossing head_a/head_b", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeStaleHeadA)

	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.NonceB = 0
	crossing.Resolution = "X"
	expect("crossing nonce_b/resolution", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeNonceBMismatch)

	var renew CrossingProposal
	if err := json.Unmarshal(byID["C005"].Operation.Value, &renew); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.NextHeadA = nil
	crossing.HandoffManifest = renew.HandoffManifest
	expect("crossing output/handoff", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeOutputRuleViolation)

	if err := json.Unmarshal(byID["C001"].Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	crossing.SignatureA = badSignature
	crossing.SignatureB = badSignature
	expect("crossing signature_a/signature_b", ApplyCrossing(byID["C001"].InitialState, crossing).Code, CodeSignatureAInvalid)

	var withdrawal WithdrawChamber
	if err := json.Unmarshal(byID["W001"].Operation.Value, &withdrawal); err != nil {
		t.Fatal(err)
	}
	withdrawal.Version = ""
	expect("withdrawal empty version", ApplyWithdrawal(byID["W001"].InitialState, withdrawal).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["W001"].Operation.Value, &withdrawal); err != nil {
		t.Fatal(err)
	}
	withdrawal.Side = "X"
	withdrawal.Nonce = 0
	expect("withdrawal side/nonce", ApplyWithdrawal(byID["W001"].InitialState, withdrawal).Code, CodeSideInvalid)

	if err := json.Unmarshal(byID["W001"].Operation.Value, &withdrawal); err != nil {
		t.Fatal(err)
	}
	withdrawal.PrevHeadB = zeroHash
	withdrawal.Side = "X"
	expect("withdrawal head_b/side", ApplyWithdrawal(byID["W001"].InitialState, withdrawal).Code, CodeStaleHeadB)

	if err := json.Unmarshal(byID["W001"].Operation.Value, &withdrawal); err != nil {
		t.Fatal(err)
	}
	withdrawal.Nonce = 0
	withdrawal.Signature = badSignature
	expect("withdrawal nonce/signature", ApplyWithdrawal(byID["W001"].InitialState, withdrawal).Code, CodeWithdrawNonceMismatch)

	var open OpenRelation
	if err := json.Unmarshal(byID["O002"].Operation.Value, &open); err != nil {
		t.Fatal(err)
	}
	open.ProfileCommitment = zeroHash
	expect("open profile/relation", ApplyOpen(byID["O002"].InitialState, open).Code, CodeProfileMismatch)

	if err := json.Unmarshal(byID["O001"].Operation.Value, &open); err != nil {
		t.Fatal(err)
	}
	open.Version = ""
	expect("open empty version", ApplyOpen(byID["O001"].InitialState, open).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["O001"].Operation.Value, &open); err != nil {
		t.Fatal(err)
	}
	open.Version = "999"
	open.ProfileCommitment = zeroHash
	expect("open version/profile", ApplyOpen(byID["O001"].InitialState, open).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["O002"].Operation.Value, &open); err != nil {
		t.Fatal(err)
	}
	open.VerifierB = open.VerifierA
	expect("open relation/verifiers", ApplyOpen(byID["O002"].InitialState, open).Code, CodeRelationAlreadyExists)

	var successor SuccessorAdmission
	if err := json.Unmarshal(byID["S003"].Operation.Value, &successor); err != nil {
		t.Fatal(err)
	}
	successor.Admission.Version = ""
	expect("successor empty version", ApplySuccessor(byID["S003"].InitialState, successor).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["S003"].Operation.Value, &successor); err != nil {
		t.Fatal(err)
	}
	successor.Admission.Version = "999"
	expect("successor version/predecessor", ApplySuccessor(byID["S003"].InitialState, successor).Code, CodeUnsupportedVersion)

	if err := json.Unmarshal(byID["S002"].Operation.Value, &successor); err != nil {
		t.Fatal(err)
	}
	expect("successor consumed/relation", ApplySuccessor(byID["S002"].InitialState, successor).Code, CodePredecessorAlreadyConsumed)

	if err := json.Unmarshal(byID["S001"].Operation.Value, &successor); err != nil {
		t.Fatal(err)
	}
	successor.Admission.InputCommitment = zeroHash
	successor.Admission.BoundsCommitment = zeroHash
	expect("successor input/surface", ApplySuccessor(byID["S001"].InitialState, successor).Code, CodeFreshInputMismatch)

	if err := json.Unmarshal(byID["S001"].Operation.Value, &successor); err != nil {
		t.Fatal(err)
	}
	successor.Admission.VerifierB = successor.Admission.VerifierA
	successor.Admission.SignatureA = badSignature
	expect("successor surface/verifiers", ApplySuccessor(byID["S001"].InitialState, successor).Code, CodeSuccessorAdmissionMismatch)
}

func TestCanonicalSPKIDER(t *testing.T) {
	const verifier = "MCowBQYDK2VwAyEANvb9A8jk+bXSut6YGGpk0604nDS+wAOXfKTEnJjpN7k="
	key, ok := parseVerifier(verifier)
	if !ok || len(key) != 32 || !sameVerifierKey(verifier, verifier) {
		t.Fatal("canonical Ed25519 SPKI was not accepted")
	}
	der, err := base64.StdEncoding.DecodeString(verifier)
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string][]byte{
		"trailing":   append(append([]byte(nil), der...), 0),
		"nonminimal": append([]byte{0x30, 0x81, 0x2a}, der[2:]...),
		"null_params": append(
			[]byte{0x30, 0x2c, 0x30, 0x07, 0x06, 0x03, 0x2b, 0x65, 0x70, 0x05, 0x00, 0x03, 0x21, 0x00},
			der[12:]...,
		),
	}
	for name, alternate := range cases {
		alternate := base64.StdEncoding.EncodeToString(alternate)
		if _, accepted := parseVerifier(alternate); accepted {
			t.Fatalf("%s alternate DER was accepted", name)
		}
		if sameVerifierKey(verifier, alternate) {
			t.Fatalf("%s alternate encoded the same admitted verifier", name)
		}
	}
}

func TestCrossChainReplayRefusedByEffectiveProfile(t *testing.T) {
	suite := loadVectorSuite(t)
	var vector vectorCase
	for _, candidate := range suite.Operations {
		if candidate.ID == "O001" {
			vector = candidate
			break
		}
	}
	if vector.ID == "" {
		t.Fatal("missing O001 replay fixture")
	}
	var operation OpenRelation
	if err := json.Unmarshal(vector.Operation.Value, &operation); err != nil {
		t.Fatal(err)
	}
	chainB := cloneChainState(vector.InitialState)
	chainB.ProfileCommitment = "0000000000000000000000000000000000000000000000000000000000000001"
	before := cloneChainState(chainB)
	result := ApplyOpen(chainB, operation)
	if result.Accepted || result.Code != CodeProfileMismatch {
		t.Fatalf("cross-chain replay = (%t, %s), want PROFILE_MISMATCH", result.Accepted, result.Code)
	}
	if result.OperationHash != "" || result.RelationID != "" {
		t.Fatal("cross-chain refusal exposed a post-profile operation identity")
	}
	if !reflect.DeepEqual(result.State, before) || !reflect.DeepEqual(chainB, before) {
		t.Fatal("cross-chain refusal changed the target chain")
	}
	encoded, err := CanonicalOpenRelationBytes(operation)
	if err != nil {
		t.Fatal(err)
	}
	profileOffset := len(openRelationDomain) + 4 + len(operation.Version) + 32
	boundProfile, err := decodeHash(operation.ProfileCommitment)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(encoded[profileOffset:profileOffset+32], boundProfile) {
		t.Fatal("signed open bytes do not bind the effective chain profile")
	}
}

func TestSuccessorRevalidatesClassifiedCarryMap(t *testing.T) {
	suite := loadVectorSuite(t)
	var vector vectorCase
	for _, candidate := range suite.Operations {
		if candidate.ID == "S001" {
			vector = candidate
			break
		}
	}
	if vector.ID == "" {
		t.Fatal("missing S001 carry-map fixture")
	}
	var operation SuccessorAdmission
	if err := json.Unmarshal(vector.Operation.Value, &operation); err != nil {
		t.Fatal(err)
	}
	state := cloneChainState(vector.InitialState)
	operation.HandoffManifest.Preserved = []string{}
	commitment, err := HandoffHash(operation.HandoffManifest)
	if err != nil {
		t.Fatal(err)
	}
	var renewal *ClosureRecord
	for index := range state.Closures {
		if state.Closures[index].Kind == "↻" {
			renewal = &state.Closures[index]
			break
		}
	}
	if renewal == nil {
		t.Fatal("S001 fixture lacks renewal closure")
	}
	renewal.HandoffCommitment = &commitment
	closureHash, err := CrossingClosureHash(renewal.PriorStateHash, renewal.ProposalHash, renewal.Kind, renewal.NextStateHash, renewal.HandoffCommitment)
	if err != nil {
		t.Fatal(err)
	}
	renewal.ClosureHash = closureHash
	sort.Slice(state.Closures, func(left, right int) bool {
		return state.Closures[left].ClosureHash < state.Closures[right].ClosureHash
	})
	operation.PredecessorClosureHash = closureHash
	operation.Admission.PredecessorCommitment = closureHash
	operation.Admission.HandoffCommitment = commitment
	if err := ValidateChainState(state); err != nil {
		t.Fatalf("recommitted adversarial predecessor is invalid: %v", err)
	}
	before := cloneChainState(state)
	result := ApplySuccessor(state, operation)
	if result.Accepted || result.Code != CodeHandoffCommitmentMismatch {
		t.Fatalf("incomplete classified carry map = (%t, %s), want HANDOFF_COMMITMENT_MISMATCH", result.Accepted, result.Code)
	}
	if !reflect.DeepEqual(result.State, before) || !reflect.DeepEqual(state, before) {
		t.Fatal("carry-map refusal changed predecessor state")
	}
}

func TestUint64CanonicalAndTransitionBoundaries(t *testing.T) {
	suite := loadVectorSuite(t)
	var vector vectorCase
	for _, candidate := range suite.Operations {
		if candidate.ID == "C001" {
			vector = candidate
			break
		}
	}
	if vector.ID == "" {
		t.Fatal("missing C001 boundary fixture")
	}
	var crossing CrossingProposal
	if err := json.Unmarshal(vector.Operation.Value, &crossing); err != nil {
		t.Fatal(err)
	}
	fullMax := ^uint64(0)
	portableMax := MaxCounterValue
	firstBeyondPortable := portableMax + 1
	if !isNext(portableMax-1, portableMax) {
		t.Fatal("last portable successor was refused")
	}
	if isNext(portableMax, firstBeyondPortable) || isNext(fullMax, 0) {
		t.Fatal("portable counter exhaustion wrapped or admitted a successor")
	}

	crossing.Epoch = portableMax
	crossing.Resolution = "0"
	crossing.NextHeadA = nil
	crossing.NextHeadB = nil
	crossing.NonceA = firstBeyondPortable
	crossing.NonceB = fullMax
	crossing.HandoffManifest = nil
	encoded, err := CanonicalProposalBytes(crossing)
	if err != nil {
		t.Fatalf("canonical boundary proposal: %v", err)
	}
	epochOffset := len(proposalDomain) + 4 + len(crossing.Version) + 32 + 32
	if got := binary.BigEndian.Uint64(encoded[epochOffset : epochOffset+8]); got != portableMax {
		t.Fatalf("canonical epoch = %d, want %d", got, portableMax)
	}
	if got := binary.BigEndian.Uint64(encoded[len(encoded)-17 : len(encoded)-9]); got != firstBeyondPortable {
		t.Fatalf("canonical nonce_a = %d, want %d", got, firstBeyondPortable)
	}
	if got := binary.BigEndian.Uint64(encoded[len(encoded)-9 : len(encoded)-1]); got != fullMax {
		t.Fatalf("canonical nonce_b = %d, want %d", got, fullMax)
	}
	if err := ValidateCrossingProposal(crossing); err == nil {
		t.Fatal("proposal above the cross-runtime portable range was admitted")
	}

	var complete CrossingProposal
	if err := json.Unmarshal(vector.Operation.Value, &complete); err != nil {
		t.Fatal(err)
	}
	relation := cloneRelation(vector.InitialState.Relations[0])
	relation.Epoch = portableMax - 1
	relation.NonceA = portableMax - 1
	relation.NonceB = portableMax - 1
	complete.Epoch = portableMax - 1
	complete.NonceA = portableMax
	complete.NonceB = portableMax
	next, advanced := transitionCrossingRelation(relation, complete)
	if !advanced || next.Epoch != portableMax || next.NonceA != portableMax || next.NonceB != portableMax {
		t.Fatal("crossing transition refused the final portable counters")
	}
	relation.Epoch = portableMax
	if _, advanced = transitionCrossingRelation(relation, complete); advanced {
		t.Fatal("crossing transition advanced an exhausted portable epoch")
	}

	withdrawal := WithdrawChamber{Side: "B", Nonce: portableMax}
	relation.Epoch = portableMax - 1
	relation.NonceA = portableMax - 2
	relation.NonceB = portableMax - 1
	next, advanced = transitionWithdrawalRelation(relation, withdrawal)
	if !advanced || next.Epoch != portableMax || next.NonceA != portableMax-2 || next.NonceB != portableMax || next.Status != "ENDED" {
		t.Fatal("withdrawal transition failed at the final portable counters")
	}
	relation.Epoch = portableMax
	if _, advanced = transitionWithdrawalRelation(relation, withdrawal); advanced {
		t.Fatal("withdrawal transition advanced an exhausted portable epoch")
	}
}

func TestRefusalOwnsDeepStateCopy(t *testing.T) {
	suite := loadVectorSuite(t)
	var vector vectorCase
	for _, candidate := range suite.Operations {
		if candidate.ID == "C006" {
			vector = candidate
			break
		}
	}
	original := cloneChainState(vector.InitialState)
	result, _ := decodeAndApply(t, vector.InitialState, vector.Operation)
	if result.Accepted {
		t.Fatal("refusal fixture unexpectedly accepted")
	}
	result.State.Relations[0].HandoffDomain[0] = "0000000000000000000000000000000000000000000000000000000000000000"
	for index := range result.State.Closures {
		if result.State.Closures[index].HandoffCommitment != nil {
			*result.State.Closures[index].HandoffCommitment = "0000000000000000000000000000000000000000000000000000000000000000"
			break
		}
	}
	if !reflect.DeepEqual(vector.InitialState, original) {
		t.Fatal("mutating a refused result aliased its input state")
	}

	var proposal CrossingProposal
	if err := json.Unmarshal(vector.Operation.Value, &proposal); err != nil {
		t.Fatal(err)
	}
	malformed := cloneChainState(vector.InitialState)
	malformed.Relations = nil
	beforeMalformed := cloneChainState(malformed)
	malformedResult := ApplyCrossing(malformed, proposal)
	if malformedResult.Code != CodeMalformedChainState || !reflect.DeepEqual(malformedResult.State, beforeMalformed) {
		t.Fatal("malformed-state refusal normalized or changed the prior state representation")
	}
}

func TestAcceptedHistoryReplay(t *testing.T) {
	suite := loadVectorSuite(t)
	state := suite.Lifecycle.InitialState
	wantReplay := []Code{
		CodeRelationAlreadyExists,
		CodeStaleEpoch,
		CodeRelationEnded,
		CodePredecessorAlreadyConsumed,
		CodeRelationEnded,
	}
	for index, step := range suite.Lifecycle.Steps {
		accepted, _ := decodeAndApply(t, state, step.Operation)
		if !accepted.Accepted {
			t.Fatalf("step %d setup refused with %s", index+1, accepted.Code)
		}
		beforeReplay := cloneChainState(accepted.State)
		replay, _ := decodeAndApply(t, accepted.State, step.Operation)
		if replay.Accepted || replay.Code != wantReplay[index] {
			t.Fatalf("step %d replay = (%t, %s), want refusal %s", index+1, replay.Accepted, replay.Code, wantReplay[index])
		}
		if !reflect.DeepEqual(replay.State, beforeReplay) || !reflect.DeepEqual(accepted.State, beforeReplay) {
			t.Fatalf("step %d replay mutated accepted history", index+1)
		}
		state = accepted.State
	}
	firstHash, err := ChainStateHash(state)
	if err != nil {
		t.Fatalf("hash first lifecycle: %v", err)
	}
	const finalChainHash = "cddbd4f92ca9122d62d99da6c83e70df5d27e10b4a0f2652fac35fb3280d95a9"
	if firstHash != finalChainHash {
		t.Fatalf("first lifecycle hash = %s, want %s", firstHash, finalChainHash)
	}

	replayed := cloneChainState(suite.Lifecycle.InitialState)
	for index, step := range suite.Lifecycle.Steps {
		accepted, _ := decodeAndApply(t, replayed, step.Operation)
		if !accepted.Accepted {
			t.Fatalf("independent replay step %d refused with %s", index+1, accepted.Code)
		}
		replayed = accepted.State
	}
	secondHash, err := ChainStateHash(replayed)
	if err != nil {
		t.Fatalf("hash independent replay: %v", err)
	}
	if secondHash != firstHash || !reflect.DeepEqual(replayed, state) {
		t.Fatalf("independent replay diverged: hash %s, want %s", secondHash, firstHash)
	}
}

func TestValidateChainStateRejectsCorruptHistoryAndLineage(t *testing.T) {
	suite := loadVectorSuite(t)
	byID := make(map[string]vectorCase, len(suite.Operations))
	for _, vector := range suite.Operations {
		byID[vector.ID] = vector
	}

	renewed := cloneChainState(byID["S001"].InitialState)
	if err := ValidateChainState(renewed); err != nil {
		t.Fatalf("renewed fixture is invalid: %v", err)
	}

	hashCorrupt := cloneChainState(renewed)
	hashCorrupt.Closures[0].ProposalHash = hashBytes([]byte("hash-corrupt"))
	if err := ValidateChainState(hashCorrupt); err == nil {
		t.Fatal("closure with a corrupted proposal hash was admitted")
	}

	dangling := cloneChainState(renewed)
	dangling.Closures[0].RelationID = hashBytes([]byte("missing-relation"))
	if err := ValidateChainState(dangling); err == nil {
		t.Fatal("closure referencing a missing relation was admitted")
	}

	disconnected := cloneChainState(renewed)
	var later *ClosureRecord
	for index := range disconnected.Closures {
		if disconnected.Closures[index].Epoch == 1 {
			later = &disconnected.Closures[index]
			break
		}
	}
	if later == nil {
		t.Fatal("renewed fixture lacks its second closure")
	}
	later.PriorStateHash = hashBytes([]byte("detached-prior-state"))
	recomputed, err := CrossingClosureHash(later.PriorStateHash, later.ProposalHash, later.Kind, later.NextStateHash, later.HandoffCommitment)
	if err != nil {
		t.Fatal(err)
	}
	later.ClosureHash = recomputed
	sort.Slice(disconnected.Closures, func(left, right int) bool {
		return disconnected.Closures[left].ClosureHash < disconnected.Closures[right].ClosureHash
	})
	if err := ValidateChainState(disconnected); err == nil {
		t.Fatal("disconnected but self-consistent closure history was admitted")
	}

	succeeded := cloneChainState(byID["S002"].InitialState)
	if err := ValidateChainState(succeeded); err != nil {
		t.Fatalf("succeeded fixture is invalid: %v", err)
	}
	targetIndex, found := relationIndex(succeeded.Relations, succeeded.Successors[0].RelationID)
	if !found {
		t.Fatal("succeeded fixture lacks its successor relation")
	}
	target := succeeded.Relations[targetIndex]
	predecessor := cloneRelation(target)
	predecessor.RelationID = hashBytes([]byte("second-predecessor-relation"))
	ended := cloneRelation(predecessor)
	ended.Epoch = 1
	ended.Status = "ENDED"
	ended.NonceA = 1
	ended.NonceB = 1
	priorHash, err := RelationStateHash(predecessor)
	if err != nil {
		t.Fatal(err)
	}
	nextHash, err := RelationStateHash(ended)
	if err != nil {
		t.Fatal(err)
	}
	handoffCommitment := hashBytes([]byte("second-handoff"))
	proposalHash := hashBytes([]byte("second-renew-proposal"))
	closureHash, err := CrossingClosureHash(priorHash, proposalHash, "↻", nextHash, &handoffCommitment)
	if err != nil {
		t.Fatal(err)
	}
	fabricatedRenew := ClosureRecord{
		ClosureHash: closureHash, RelationID: ended.RelationID, Epoch: 0,
		PriorStateHash: priorHash, ProposalHash: proposalHash, Kind: "↻",
		NextStateHash: nextHash, HandoffCommitment: &handoffCommitment,
	}
	duplicateOwner := cloneChainState(succeeded)
	duplicateOwner.Relations = insertRelation(duplicateOwner.Relations, ended)
	duplicateOwner.Closures = insertClosure(duplicateOwner.Closures, fabricatedRenew)
	duplicateOwner.Successors = insertSuccessor(duplicateOwner.Successors, SuccessorLink{
		PredecessorClosureHash: fabricatedRenew.ClosureHash,
		AdmissionHash:          hashBytes([]byte("second-successor-admission")),
		RelationID:             target.RelationID,
	})
	if err := ValidateChainState(duplicateOwner); err == nil {
		t.Fatal("successor relation with two predecessor owners was admitted")
	}

	selfLineage := cloneChainState(succeeded)
	predecessorClosure, found := closureIndex(selfLineage.Closures, selfLineage.Successors[0].PredecessorClosureHash)
	if !found {
		t.Fatal("succeeded fixture lacks its predecessor closure")
	}
	selfLineage.Successors[0].RelationID = selfLineage.Closures[predecessorClosure].RelationID
	if err := ValidateChainState(selfLineage); err == nil {
		t.Fatal("self-successor lineage was admitted")
	}

	firstPredecessor := succeeded.Closures[predecessorClosure].RelationID
	cyclic := cloneChainState(succeeded)
	cyclicTargetIndex, found := relationIndex(cyclic.Relations, target.RelationID)
	if !found {
		t.Fatal("cyclic fixture lacks its successor relation")
	}
	cyclicInitial := cloneRelation(cyclic.Relations[cyclicTargetIndex])
	cyclicEnded := cloneRelation(cyclicInitial)
	cyclicEnded.Epoch = 1
	cyclicEnded.Status = "ENDED"
	cyclicEnded.NonceA = 1
	cyclicEnded.NonceB = 1
	cyclicPriorHash, err := RelationStateHash(cyclicInitial)
	if err != nil {
		t.Fatal(err)
	}
	cyclicNextHash, err := RelationStateHash(cyclicEnded)
	if err != nil {
		t.Fatal(err)
	}
	cycleHandoff := hashBytes([]byte("cycle-handoff"))
	cycleProposal := hashBytes([]byte("cycle-renew-proposal"))
	cycleClosureHash, err := CrossingClosureHash(cyclicPriorHash, cycleProposal, "↻", cyclicNextHash, &cycleHandoff)
	if err != nil {
		t.Fatal(err)
	}
	cycleClosure := ClosureRecord{
		ClosureHash: cycleClosureHash, RelationID: cyclicEnded.RelationID, Epoch: 0,
		PriorStateHash: cyclicPriorHash, ProposalHash: cycleProposal, Kind: "↻",
		NextStateHash: cyclicNextHash, HandoffCommitment: &cycleHandoff,
	}
	cyclic.Relations[cyclicTargetIndex] = cyclicEnded
	cyclic.Closures = insertClosure(cyclic.Closures, cycleClosure)
	cyclic.Successors = insertSuccessor(cyclic.Successors, SuccessorLink{
		PredecessorClosureHash: cycleClosure.ClosureHash,
		AdmissionHash:          hashBytes([]byte("cycle-admission")),
		RelationID:             firstPredecessor,
	})
	if err := ValidateChainState(cyclic); err == nil {
		t.Fatal("cyclic successor lineage was admitted")
	}
}

func TestValidateChainStateTreatsAdmissionHashAsReplayIndex(t *testing.T) {
	suite := loadVectorSuite(t)
	var succeeded ChainState
	for _, vector := range suite.Operations {
		if vector.ID == "S002" {
			succeeded = cloneChainState(vector.InitialState)
			break
		}
	}
	if len(succeeded.Successors) != 1 {
		t.Fatal("missing succeeded fixture")
	}
	succeeded.Successors[0].AdmissionHash = hashBytes([]byte("opaque-altered-admission"))
	if err := ValidateChainState(succeeded); err != nil {
		t.Fatalf("structural validation rejected an opaque admission replay index: %v", err)
	}
}
