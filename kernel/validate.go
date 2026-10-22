package kernel

import (
	"bytes"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
)

func validHash(value string) bool {
	if len(value) != 64 {
		return false
	}
	decoded, err := hex.DecodeString(value)
	return err == nil && len(decoded) == 32 && hex.EncodeToString(decoded) == value
}

func validBase64(value string) ([]byte, bool) {
	if value == "" {
		return nil, false
	}
	decoded, err := base64.StdEncoding.DecodeString(value)
	return decoded, err == nil && base64.StdEncoding.EncodeToString(decoded) == value
}

func parseVerifier(value string) (ed25519.PublicKey, bool) {
	der, ok := validBase64(value)
	if !ok {
		return nil, false
	}
	parsed, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, false
	}
	key, ok := parsed.(ed25519.PublicKey)
	if !ok || len(key) != ed25519.PublicKeySize {
		return nil, false
	}
	canonical, err := x509.MarshalPKIXPublicKey(key)
	if err != nil || !bytes.Equal(canonical, der) {
		return nil, false
	}
	return key, true
}

func sameVerifierKey(left, right string) bool {
	leftKey, leftOK := parseVerifier(left)
	rightKey, rightOK := parseVerifier(right)
	return leftOK && rightOK && bytes.Equal(leftKey, rightKey)
}

func validSignature(value string) bool {
	decoded, ok := validBase64(value)
	return ok && len(decoded) == ed25519.SignatureSize
}

func validCounter(value uint64) bool {
	return value <= MaxCounterValue
}

func sortedUniqueHashes(values []string) bool {
	if values == nil {
		return false
	}
	for index, value := range values {
		if !validHash(value) || (index > 0 && values[index-1] >= value) {
			return false
		}
	}
	return true
}

func validManifestShape(value HandoffManifest) bool {
	if !validHash(value.SuccessorInputCommitment) || !validHash(value.SuccessorAdmissionCommitment) {
		return false
	}
	if value.Preserved == nil || value.Transformed == nil || value.Ended == nil || value.Refused == nil || value.Unresolved == nil {
		return false
	}
	for _, values := range [][]string{value.Preserved, value.Ended, value.Refused, value.Unresolved} {
		for _, hash := range values {
			if !validHash(hash) {
				return false
			}
		}
	}
	for _, transform := range value.Transformed {
		if !validHash(transform.From) || !validHash(transform.To) {
			return false
		}
	}
	return true
}

func ManifestRule(domain []string, manifest HandoffManifest) bool {
	if !validManifestShape(manifest) || !sortedUniqueHashes(manifest.Preserved) || !sortedUniqueHashes(manifest.Ended) || !sortedUniqueHashes(manifest.Refused) || !sortedUniqueHashes(manifest.Unresolved) {
		return false
	}
	for index, transform := range manifest.Transformed {
		if index > 0 {
			previous := manifest.Transformed[index-1]
			if previous.From > transform.From || (previous.From == transform.From && previous.To >= transform.To) {
				return false
			}
		}
	}
	sources := make([]string, 0, len(manifest.Preserved)+len(manifest.Transformed)+len(manifest.Ended)+len(manifest.Refused)+len(manifest.Unresolved))
	sources = append(sources, manifest.Preserved...)
	for _, transform := range manifest.Transformed {
		sources = append(sources, transform.From)
	}
	sources = append(sources, manifest.Ended...)
	sources = append(sources, manifest.Refused...)
	sources = append(sources, manifest.Unresolved...)
	sort.Strings(sources)
	if len(sources) != len(domain) {
		return false
	}
	for index := range sources {
		if sources[index] != domain[index] || (index > 0 && sources[index-1] == sources[index]) {
			return false
		}
	}
	carried := make([]string, 0, len(manifest.Preserved)+len(manifest.Transformed)+len(manifest.Unresolved))
	carried = append(carried, manifest.Preserved...)
	for _, transform := range manifest.Transformed {
		carried = append(carried, transform.To)
	}
	carried = append(carried, manifest.Unresolved...)
	sort.Strings(carried)
	for index, hash := range carried {
		if index > 0 && carried[index-1] == hash {
			return false
		}
		if containsSorted(manifest.Ended, hash) || containsSorted(manifest.Refused, hash) {
			return false
		}
	}
	return true
}

func SuccessorDomain(manifest HandoffManifest) []string {
	values := make([]string, 0, len(manifest.Preserved)+len(manifest.Transformed)+len(manifest.Unresolved))
	values = append(values, manifest.Preserved...)
	for _, transform := range manifest.Transformed {
		values = append(values, transform.To)
	}
	values = append(values, manifest.Unresolved...)
	sort.Strings(values)
	return values
}

func containsSorted(values []string, target string) bool {
	index := sort.SearchStrings(values, target)
	return index < len(values) && values[index] == target
}

func ValidateRelationState(value RelationState) error {
	if value.Version == "" || !validHash(value.RelationID) || !validHash(value.ProfileCommitment) {
		return errors.New("invalid relation identity")
	}
	if value.Status != "ACTIVE" && value.Status != "ENDED" {
		return errors.New("invalid relation status")
	}
	if !validCounter(value.Epoch) || !validCounter(value.NonceA) || !validCounter(value.NonceB) {
		return errors.New("relation counter exceeds portable integer range")
	}
	if !validHash(value.HeadA) || !validHash(value.HeadB) {
		return errors.New("invalid relation head")
	}
	if _, ok := parseVerifier(value.VerifierA); !ok {
		return errors.New("invalid verifier_a")
	}
	if _, ok := parseVerifier(value.VerifierB); !ok {
		return errors.New("invalid verifier_b")
	}
	if sameVerifierKey(value.VerifierA, value.VerifierB) {
		return errors.New("verifier scopes are not distinct")
	}
	if !sortedUniqueHashes(value.HandoffDomain) {
		return errors.New("handoff domain must be sorted and unique")
	}
	return nil
}

func ValidateCrossingProposal(value CrossingProposal) error {
	if !validHash(value.RelationID) || !validHash(value.ProfileCommitment) || !validHash(value.PrevHeadA) || !validHash(value.PrevHeadB) {
		return errors.New("invalid proposal identity")
	}
	if value.NextHeadA != nil && !validHash(*value.NextHeadA) {
		return errors.New("invalid next_head_a")
	}
	if value.NextHeadB != nil && !validHash(*value.NextHeadB) {
		return errors.New("invalid next_head_b")
	}
	if !validCounter(value.Epoch) || !validCounter(value.NonceA) || !validCounter(value.NonceB) {
		return errors.New("proposal counter exceeds portable integer range")
	}
	if value.HandoffManifest != nil && !validManifestShape(*value.HandoffManifest) {
		return errors.New("invalid handoff manifest")
	}
	if !validSignature(value.SignatureA) || !validSignature(value.SignatureB) {
		return errors.New("invalid proposal signature shape")
	}
	return nil
}

func ValidateOpenRelation(value OpenRelation) error {
	for name, hash := range map[string]string{
		"relation_id": value.RelationID,
		"profile_commitment": value.ProfileCommitment,
		"invocation_id": value.InvocationID,
		"input_commitment": value.InputCommitment,
		"bounds_commitment": value.BoundsCommitment,
		"head_a": value.HeadA,
		"head_b": value.HeadB,
	} {
		if !validHash(hash) {
			return fmt.Errorf("invalid %s", name)
		}
	}
	if _, ok := parseVerifier(value.VerifierA); !ok {
		return errors.New("invalid verifier_a")
	}
	if _, ok := parseVerifier(value.VerifierB); !ok {
		return errors.New("invalid verifier_b")
	}
	if !sortedUniqueHashes(value.HandoffDomain) {
		return errors.New("handoff domain must be sorted and unique")
	}
	if !validSignature(value.SignatureA) || !validSignature(value.SignatureB) {
		return errors.New("invalid open signature shape")
	}
	return nil
}

func ValidateFreshAdmission(value FreshAdmission) error {
	for name, hash := range map[string]string{
		"relation_id": value.RelationID,
		"profile_commitment": value.ProfileCommitment,
		"invocation_id": value.InvocationID,
		"predecessor_commitment": value.PredecessorCommitment,
		"handoff_commitment": value.HandoffCommitment,
		"input_commitment": value.InputCommitment,
		"bounds_commitment": value.BoundsCommitment,
		"head_a": value.HeadA,
		"head_b": value.HeadB,
	} {
		if !validHash(hash) {
			return fmt.Errorf("invalid %s", name)
		}
	}
	if _, ok := parseVerifier(value.VerifierA); !ok {
		return errors.New("invalid verifier_a")
	}
	if _, ok := parseVerifier(value.VerifierB); !ok {
		return errors.New("invalid verifier_b")
	}
	if !validSignature(value.SignatureA) || !validSignature(value.SignatureB) {
		return errors.New("invalid fresh signature shape")
	}
	return nil
}

func ValidateSuccessorAdmission(value SuccessorAdmission) error {
	if !validHash(value.PredecessorClosureHash) || !validManifestShape(value.HandoffManifest) {
		return errors.New("invalid predecessor or handoff")
	}
	return ValidateFreshAdmission(value.Admission)
}

func ValidateWithdrawal(value WithdrawChamber) error {
	if !validHash(value.RelationID) || !validHash(value.ProfileCommitment) || !validHash(value.PrevHeadA) || !validHash(value.PrevHeadB) {
		return errors.New("invalid withdrawal identity")
	}
	if !validSignature(value.Signature) {
		return errors.New("invalid withdrawal signature shape")
	}
	if !validCounter(value.Epoch) || !validCounter(value.Nonce) {
		return errors.New("withdrawal counter exceeds portable integer range")
	}
	return nil
}

func ValidateChainState(value ChainState) error {
	if value.ProtocolVersion != ProtocolVersion || value.KernelVersion != Version || !validHash(value.ProfileCommitment) {
		return errors.New("invalid chain config")
	}
	if value.Relations == nil || value.Closures == nil || value.Successors == nil {
		return errors.New("nil chain registry")
	}
	for index, relation := range value.Relations {
		if err := ValidateRelationState(relation); err != nil {
			return fmt.Errorf("relation %d: %w", index, err)
		}
		if relation.Version != Version || relation.ProfileCommitment != value.ProfileCommitment {
			return errors.New("relation differs from pinned chain profile")
		}
		if index > 0 && value.Relations[index-1].RelationID >= relation.RelationID {
			return errors.New("relations are not strictly sorted")
		}
	}
	closuresByRelation := make(map[string][]ClosureRecord, len(value.Relations))
	seenEpochs := make(map[string]struct{}, len(value.Closures))
	closuresByHash := make(map[string]ClosureRecord, len(value.Closures))
	for index, closure := range value.Closures {
		if err := validateClosureRecord(closure); err != nil {
			return fmt.Errorf("closure %d: %w", index, err)
		}
		if index > 0 && value.Closures[index-1].ClosureHash >= closure.ClosureHash {
			return errors.New("closures are not strictly sorted")
		}
		if _, found := relationIndex(value.Relations, closure.RelationID); !found {
			return errors.New("closure references missing relation")
		}
		epochKey := fmt.Sprintf("%s:%d", closure.RelationID, closure.Epoch)
		if _, found := seenEpochs[epochKey]; found {
			return errors.New("duplicate relation closure epoch")
		}
		seenEpochs[epochKey] = struct{}{}
		closuresByHash[closure.ClosureHash] = closure
		closuresByRelation[closure.RelationID] = append(closuresByRelation[closure.RelationID], closure)
	}
	for _, relation := range value.Relations {
		history := closuresByRelation[relation.RelationID]
		if uint64(len(history)) != relation.Epoch {
			return errors.New("relation epoch differs from closure history")
		}
		if len(history) == 0 {
			if relation.Status != "ACTIVE" || relation.NonceA != 0 || relation.NonceB != 0 {
				return errors.New("fresh relation state is not initial")
			}
			continue
		}
		sort.Slice(history, func(left, right int) bool {
			return history[left].Epoch < history[right].Epoch
		})
		var expectedNonceA uint64
		var expectedNonceB uint64
		for index, closure := range history {
			if closure.Epoch != uint64(index) {
				return errors.New("relation closure epochs are not contiguous")
			}
			if index > 0 && history[index-1].NextStateHash != closure.PriorStateHash {
				return errors.New("relation closure history is disconnected")
			}
			terminal := closure.Kind != "1"
			if terminal && index != len(history)-1 {
				return errors.New("terminal closure is followed by history")
			}
			if closure.Kind == "WITHDRAW_A" {
				if expectedNonceA == ^uint64(0) {
					return errors.New("nonce_a exhausted")
				}
				expectedNonceA++
			} else if closure.Kind == "WITHDRAW_B" {
				if expectedNonceB == ^uint64(0) {
					return errors.New("nonce_b exhausted")
				}
				expectedNonceB++
			} else {
				if expectedNonceA == ^uint64(0) || expectedNonceB == ^uint64(0) {
					return errors.New("crossing nonce exhausted")
				}
				expectedNonceA++
				expectedNonceB++
			}
		}
		finalClosure := history[len(history)-1]
		expectedStatus := "ACTIVE"
		if finalClosure.Kind != "1" {
			expectedStatus = "ENDED"
		}
		currentHash, err := RelationStateHash(relation)
		if err != nil || relation.Status != expectedStatus || relation.NonceA != expectedNonceA || relation.NonceB != expectedNonceB || finalClosure.NextStateHash != currentHash {
			return errors.New("relation state differs from closure history")
		}
	}
	ownedRelations := make(map[string]struct{}, len(value.Successors))
	nextByRelation := make(map[string]string, len(value.Successors))
	for index, successor := range value.Successors {
		if !validHash(successor.PredecessorClosureHash) || !validHash(successor.AdmissionHash) || !validHash(successor.RelationID) {
			return errors.New("invalid successor link")
		}
		if index > 0 && value.Successors[index-1].PredecessorClosureHash >= successor.PredecessorClosureHash {
			return errors.New("successors are not strictly sorted")
		}
		predecessor, found := closuresByHash[successor.PredecessorClosureHash]
		if !found || predecessor.Kind != "↻" {
			return errors.New("successor does not reference a renewal closure")
		}
		if _, found := relationIndex(value.Relations, successor.RelationID); !found {
			return errors.New("successor references missing relation")
		}
		if _, found := ownedRelations[successor.RelationID]; found {
			return errors.New("successor relation has multiple predecessors")
		}
		ownedRelations[successor.RelationID] = struct{}{}
		if predecessor.RelationID == successor.RelationID {
			return errors.New("self succession")
		}
		if _, found := nextByRelation[predecessor.RelationID]; found {
			return errors.New("predecessor relation has multiple successors")
		}
		nextByRelation[predecessor.RelationID] = successor.RelationID
	}
	for start := range nextByRelation {
		visited := make(map[string]struct{}, len(nextByRelation))
		cursor := start
		for {
			next, found := nextByRelation[cursor]
			if !found {
				break
			}
			if _, seen := visited[cursor]; seen {
				return errors.New("cyclic successor lineage")
			}
			visited[cursor] = struct{}{}
			cursor = next
		}
	}
	return nil
}

func validateClosureRecord(value ClosureRecord) error {
	if !validHash(value.ClosureHash) || !validHash(value.RelationID) || !validHash(value.PriorStateHash) || !validHash(value.ProposalHash) || !validHash(value.NextStateHash) {
		return errors.New("invalid closure hash field")
	}
	if !validCounter(value.Epoch) {
		return errors.New("closure epoch exceeds portable integer range")
	}
	if value.HandoffCommitment != nil && !validHash(*value.HandoffCommitment) {
		return errors.New("invalid handoff commitment")
	}
	var expected string
	var err error
	switch value.Kind {
	case "1", "0", "↻":
		expected, err = CrossingClosureHash(value.PriorStateHash, value.ProposalHash, value.Kind, value.NextStateHash, value.HandoffCommitment)
	case "WITHDRAW_A":
		if value.HandoffCommitment != nil {
			return errors.New("withdrawal closure has a handoff")
		}
		expected, err = WithdrawalClosureHash(value.PriorStateHash, value.ProposalHash, "A", value.NextStateHash)
	case "WITHDRAW_B":
		if value.HandoffCommitment != nil {
			return errors.New("withdrawal closure has a handoff")
		}
		expected, err = WithdrawalClosureHash(value.PriorStateHash, value.ProposalHash, "B", value.NextStateHash)
	default:
		return errors.New("invalid closure kind")
	}
	if err != nil || expected != value.ClosureHash {
		return errors.New("closure hash mismatch")
	}
	if value.Kind == "↻" && value.HandoffCommitment == nil {
		return errors.New("renewal closure lacks handoff")
	}
	if value.Kind != "↻" && value.HandoffCommitment != nil {
		return errors.New("non-renewal closure carries handoff")
	}
	return nil
}

func relationIndex(values []RelationState, relationID string) (int, bool) {
	index := sort.Search(len(values), func(index int) bool {
		return values[index].RelationID >= relationID
	})
	return index, index < len(values) && values[index].RelationID == relationID
}

func closureIndex(values []ClosureRecord, closureHash string) (int, bool) {
	index := sort.Search(len(values), func(index int) bool {
		return values[index].ClosureHash >= closureHash
	})
	return index, index < len(values) && values[index].ClosureHash == closureHash
}

func successorIndex(values []SuccessorLink, closureHash string) (int, bool) {
	index := sort.Search(len(values), func(index int) bool {
		return values[index].PredecessorClosureHash >= closureHash
	})
	return index, index < len(values) && values[index].PredecessorClosureHash == closureHash
}
