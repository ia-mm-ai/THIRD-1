package avalanchevm

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

const testNetworkID = uint32(999)

var testSubnetID = ids.ID(sha256.Sum256([]byte("continuity-test-subnet")))
var testChainID = ids.ID(sha256.Sum256([]byte("continuity-test-chain")))

func testHash(label string) string {
	digest := sha256.Sum256([]byte(label))
	return hex.EncodeToString(digest[:])
}

func testSigner(t *testing.T) (string, ed25519.PrivateKey) {
	t.Helper()
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKIXPublicKey(privateKey.Public())
	if err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(der), privateKey
}

func signedOpen(t *testing.T, relationLabel, profile string) kernel.OpenRelation {
	t.Helper()
	verifierA, privateA := testSigner(t)
	verifierB, privateB := testSigner(t)
	operation := kernel.OpenRelation{
		Version: kernel.Version, RelationID: testHash(relationLabel), ProfileCommitment: profile,
		InvocationID: testHash(relationLabel + ":invoke"), InputCommitment: testHash(relationLabel + ":input"),
		BoundsCommitment: testHash(relationLabel + ":bounds"), HeadA: testHash(relationLabel + ":a"),
		VerifierA: verifierA, HeadB: testHash(relationLabel + ":b"), VerifierB: verifierB,
		HandoffDomain: []string{},
	}
	payload, err := kernel.CanonicalOpenRelationBytes(operation)
	if err != nil {
		t.Fatal(err)
	}
	operation.SignatureA = base64.StdEncoding.EncodeToString(ed25519.Sign(privateA, payload))
	operation.SignatureB = base64.StdEncoding.EncodeToString(ed25519.Sign(privateB, payload))
	return operation
}

func testChainState(t *testing.T) kernel.ChainState {
	t.Helper()
	profileCommitment, err := deriveChainProfileCommitment(testHash("profile"), testNetworkID, testSubnetID, testChainID, kernel.Version)
	if err != nil {
		t.Fatal(err)
	}
	state, err := kernel.NewChainState(kernel.ChainConfig{
		ProtocolVersion: kernel.ProtocolVersion,
		KernelVersion: kernel.Version,
		ProfileCommitment: profileCommitment,
	})
	if err != nil {
		t.Fatal(err)
	}
	return state
}

func TestOperationRoundTripAndApply(t *testing.T) {
	state := testChainState(t)
	first, err := NewOperation(ActionOpenRelation, signedOpen(t, "z-relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	parsed, operationID, err := parseOperation(first.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if operationID != first.ID() || parsed.Action != ActionOpenRelation {
		t.Fatal("operation identity or action changed across round trip")
	}
	result, err := applyOperation(state, parsed)
	if err != nil {
		t.Fatal(err)
	}
	second, err := NewOperation(ActionOpenRelation, signedOpen(t, "a-relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	result, err = applyOperation(result.State, second)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.State.Relations) != 2 || result.State.Relations[0].RelationID >= result.State.Relations[1].RelationID {
		t.Fatal("multi-relation registry was not kept strictly sorted")
	}
}

func TestNullableCrossingFieldsAreExplicit(t *testing.T) {
	zeroSignature := base64.StdEncoding.EncodeToString(make([]byte, ed25519.SignatureSize))
	proposal := kernel.CrossingProposal{
		Version: kernel.Version, RelationID: testHash("relation"), ProfileCommitment: testHash("profile"),
		PrevHeadA: testHash("a"), PrevHeadB: testHash("b"), Resolution: "0",
		NextHeadA: nil, NextHeadB: nil, NonceA: 1, NonceB: 1, HandoffManifest: nil,
		SignatureA: zeroSignature, SignatureB: zeroSignature,
	}
	operation, err := NewOperation(ActionApplyCrossing, proposal)
	if err != nil {
		t.Fatal(err)
	}
	data := operation.Bytes()
	for _, field := range [][]byte{[]byte(`"next_head_a":null`), []byte(`"next_head_b":null`), []byte(`"handoff_manifest":null`)} {
		if !bytes.Contains(data, field) {
			t.Fatalf("canonical operation omitted %s", field)
		}
	}
	omitted := bytes.Replace(data, []byte(`"next_head_a":null,`), nil, 1)
	if _, _, err := parseOperation(omitted); err == nil {
		t.Fatal("omitted nullable field was accepted")
	}
}

func TestReceiptPreservesNullableKernelResultFields(t *testing.T) {
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "receipt-null", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	result, err := applyOperation(state, operation)
	if err != nil {
		t.Fatal(err)
	}
	receipt, err := receiptFor(operation, result, testHash("block"), 1, 0)
	if err != nil {
		t.Fatal(err)
	}
	data, err := canonicalJSON(receipt)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte(`"closure_hash":null`)) {
		t.Fatal("receipt collapsed an absent kernel result hash into a string")
	}
	decoded := new(Receipt)
	if err := decodeCanonical(data, decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.ClosureHash != nil {
		t.Fatal("receipt null did not survive canonical round trip")
	}
}

func TestOperationRejectsNoncanonicalAndNestedSmuggling(t *testing.T) {
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	data := operation.Bytes()
	unknown := bytes.Replace(data, []byte(`"version":"002"`), []byte(`"version":"002","unknown":0`), 1)
	if _, _, err := parseOperation(unknown); err == nil {
		t.Fatal("nested unknown field was accepted")
	}
	duplicate := bytes.Replace(data, []byte(`"relation_id":`), []byte(`"version":"002","relation_id":`), 1)
	if _, _, err := parseOperation(duplicate); err == nil {
		t.Fatal("nested duplicate field was accepted")
	}
	if _, _, err := parseOperation(append(data, '\n')); err == nil {
		t.Fatal("noncanonical trailing whitespace was accepted")
	}
}

func TestDurablePendingRoundTrip(t *testing.T) {
	state := testChainState(t)
	operation, err := NewOperation(ActionOpenRelation, signedOpen(t, "relation", state.ProfileCommitment))
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := encodeDurablePendingOperation(7, operation.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	decoded, operationID, sequence, err := decodeDurablePendingOperation(encoded, maxOperationBytes)
	if err != nil {
		t.Fatal(err)
	}
	if sequence != 7 || operationID != operation.ID() || !bytes.Equal(decoded.Bytes(), operation.Bytes()) {
		t.Fatal("durable pending operation changed identity, order, or bytes")
	}
}
