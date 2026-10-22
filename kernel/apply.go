package kernel

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"sort"
)

func NewChainState(config ChainConfig) (ChainState, error) {
	state := ChainState{
		ProtocolVersion:   config.ProtocolVersion,
		KernelVersion:     config.KernelVersion,
		ProfileCommitment: config.ProfileCommitment,
		Relations:         []RelationState{},
		Closures:          []ClosureRecord{},
		Successors:        []SuccessorLink{},
	}
	if err := ValidateChainState(state); err != nil {
		return ChainState{}, err
	}
	return state, nil
}

func cloneStrings(values []string) []string {
	if values == nil {
		return nil
	}
	cloned := make([]string, len(values))
	copy(cloned, values)
	return cloned
}

func cloneOptionalHash(value *string) *string {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}

func cloneRelation(value RelationState) RelationState {
	value.HandoffDomain = cloneStrings(value.HandoffDomain)
	return value
}

func cloneChainState(value ChainState) ChainState {
	cloned := ChainState{
		ProtocolVersion:   value.ProtocolVersion,
		KernelVersion:     value.KernelVersion,
		ProfileCommitment: value.ProfileCommitment,
	}
	if value.Relations != nil {
		cloned.Relations = make([]RelationState, len(value.Relations))
		for index, relation := range value.Relations {
			cloned.Relations[index] = cloneRelation(relation)
		}
	}
	if value.Closures != nil {
		cloned.Closures = make([]ClosureRecord, len(value.Closures))
		for index, closure := range value.Closures {
			cloned.Closures[index] = closure
			cloned.Closures[index].HandoffCommitment = cloneOptionalHash(closure.HandoffCommitment)
		}
	}
	if value.Successors != nil {
		cloned.Successors = make([]SuccessorLink, len(value.Successors))
		copy(cloned.Successors, value.Successors)
	}
	return cloned
}

func refusal(code Code, state ChainState, relationID string) ApplyResult {
	return ApplyResult{
		Accepted:   false,
		Code:       code,
		State:      cloneChainState(state),
		RelationID: relationID,
	}
}

func verifySignature(verifier string, payload []byte, signature string) bool {
	key, ok := parseVerifier(verifier)
	if !ok {
		return false
	}
	rawSignature, ok := validBase64(signature)
	return ok && ed25519.Verify(key, payload, rawSignature)
}

func insertRelation(values []RelationState, value RelationState) []RelationState {
	index := sort.Search(len(values), func(index int) bool {
		return values[index].RelationID >= value.RelationID
	})
	values = append(values, RelationState{})
	copy(values[index+1:], values[index:])
	values[index] = cloneRelation(value)
	return values
}

func insertClosure(values []ClosureRecord, value ClosureRecord) []ClosureRecord {
	index := sort.Search(len(values), func(index int) bool {
		return values[index].ClosureHash >= value.ClosureHash
	})
	values = append(values, ClosureRecord{})
	copy(values[index+1:], values[index:])
	value.HandoffCommitment = cloneOptionalHash(value.HandoffCommitment)
	values[index] = value
	return values
}

func insertSuccessor(values []SuccessorLink, value SuccessorLink) []SuccessorLink {
	index := sort.Search(len(values), func(index int) bool {
		return values[index].PredecessorClosureHash >= value.PredecessorClosureHash
	})
	values = append(values, SuccessorLink{})
	copy(values[index+1:], values[index:])
	values[index] = value
	return values
}

func transitionCrossingRelation(relation RelationState, proposal CrossingProposal) (RelationState, bool) {
	if relation.Epoch >= MaxCounterValue {
		return RelationState{}, false
	}
	next := cloneRelation(relation)
	next.Epoch = relation.Epoch + 1
	next.NonceA = proposal.NonceA
	next.NonceB = proposal.NonceB
	if proposal.Resolution == "1" {
		next.HeadA = *proposal.NextHeadA
		next.HeadB = *proposal.NextHeadB
	} else {
		next.Status = "ENDED"
	}
	return next, true
}

func transitionWithdrawalRelation(relation RelationState, operation WithdrawChamber) (RelationState, bool) {
	if relation.Epoch >= MaxCounterValue {
		return RelationState{}, false
	}
	next := cloneRelation(relation)
	next.Epoch = relation.Epoch + 1
	next.Status = "ENDED"
	if operation.Side == "A" {
		next.NonceA = operation.Nonce
	} else {
		next.NonceB = operation.Nonce
	}
	return next, true
}

func ApplyOpen(state ChainState, operation OpenRelation) ApplyResult {
	if err := ValidateChainState(state); err != nil {
		return refusal(CodeMalformedChainState, state, "")
	}
	if err := ValidateOpenRelation(operation); err != nil {
		return refusal(CodeMalformedOpen, state, "")
	}
	if operation.Version != Version {
		return refusal(CodeUnsupportedVersion, state, "")
	}
	if operation.ProfileCommitment != state.ProfileCommitment {
		return refusal(CodeProfileMismatch, state, "")
	}
	if _, exists := relationIndex(state.Relations, operation.RelationID); exists {
		return refusal(CodeRelationAlreadyExists, state, operation.RelationID)
	}
	if sameVerifierKey(operation.VerifierA, operation.VerifierB) {
		return refusal(CodeVerifierScopesNotDistinct, state, operation.RelationID)
	}
	payload, err := CanonicalOpenRelationBytes(operation)
	if err != nil {
		return refusal(CodeMalformedOpen, state, operation.RelationID)
	}
	digest := hashBytes(payload)
	if !verifySignature(operation.VerifierA, payload, operation.SignatureA) {
		result := refusal(CodeOpenSignatureAInvalid, state, operation.RelationID)
		result.OperationHash = digest
		return result
	}
	if !verifySignature(operation.VerifierB, payload, operation.SignatureB) {
		result := refusal(CodeOpenSignatureBInvalid, state, operation.RelationID)
		result.OperationHash = digest
		return result
	}
	relation := RelationState{
		Version:           Version,
		RelationID:        operation.RelationID,
		ProfileCommitment: operation.ProfileCommitment,
		Epoch:             0,
		Status:            "ACTIVE",
		HeadA:             operation.HeadA,
		VerifierA:         operation.VerifierA,
		NonceA:            0,
		HeadB:             operation.HeadB,
		VerifierB:         operation.VerifierB,
		NonceB:            0,
		HandoffDomain:     cloneStrings(operation.HandoffDomain),
	}
	next := cloneChainState(state)
	next.Relations = insertRelation(next.Relations, relation)
	if err := ValidateChainState(next); err != nil {
		return refusal(CodeMalformedChainState, state, operation.RelationID)
	}
	relationDigest, _ := RelationStateHash(relation)
	return ApplyResult{
		Accepted:          true,
		Code:              CodeOpenAccepted,
		State:             next,
		RelationID:        operation.RelationID,
		OperationHash:     digest,
		RelationStateHash: relationDigest,
	}
}

func ApplyCrossing(state ChainState, proposal CrossingProposal) ApplyResult {
	if err := ValidateChainState(state); err != nil {
		return refusal(CodeMalformedChainState, state, "")
	}
	if err := ValidateCrossingProposal(proposal); err != nil {
		return refusal(CodeMalformedProposal, state, proposal.RelationID)
	}
	if proposal.Version != Version {
		return refusal(CodeUnsupportedVersion, state, proposal.RelationID)
	}
	index, found := relationIndex(state.Relations, proposal.RelationID)
	if !found {
		return refusal(CodeRelationNotFound, state, proposal.RelationID)
	}
	relation := state.Relations[index]
	if proposal.ProfileCommitment != relation.ProfileCommitment {
		return refusal(CodeProfileMismatch, state, proposal.RelationID)
	}
	if relation.Status == "ENDED" {
		return refusal(CodeRelationEnded, state, proposal.RelationID)
	}
	if proposal.Epoch != relation.Epoch {
		return refusal(CodeStaleEpoch, state, proposal.RelationID)
	}
	if proposal.PrevHeadA != relation.HeadA {
		return refusal(CodeStaleHeadA, state, proposal.RelationID)
	}
	if proposal.PrevHeadB != relation.HeadB {
		return refusal(CodeStaleHeadB, state, proposal.RelationID)
	}
	if !isNext(relation.NonceA, proposal.NonceA) {
		return refusal(CodeNonceAMismatch, state, proposal.RelationID)
	}
	if !isNext(relation.NonceB, proposal.NonceB) {
		return refusal(CodeNonceBMismatch, state, proposal.RelationID)
	}
	if proposal.Resolution != "1" && proposal.Resolution != "0" && proposal.Resolution != "↻" {
		return refusal(CodeResolutionInvalid, state, proposal.RelationID)
	}
	payload, err := CanonicalProposalBytes(proposal)
	if err != nil {
		return refusal(CodeMalformedProposal, state, proposal.RelationID)
	}
	proposalDigest := hashBytes(payload)
	hasA := proposal.NextHeadA != nil
	hasB := proposal.NextHeadB != nil
	if (proposal.Resolution == "1" && (!hasA || !hasB)) || (proposal.Resolution != "1" && (hasA || hasB)) {
		result := refusal(CodeOutputRuleViolation, state, proposal.RelationID)
		result.OperationHash = proposalDigest
		return result
	}
	if proposal.Resolution == "↻" {
		if proposal.HandoffManifest == nil || !ManifestRule(relation.HandoffDomain, *proposal.HandoffManifest) {
			result := refusal(CodeHandoffRuleViolation, state, proposal.RelationID)
			result.OperationHash = proposalDigest
			return result
		}
	} else if proposal.HandoffManifest != nil {
		result := refusal(CodeHandoffRuleViolation, state, proposal.RelationID)
		result.OperationHash = proposalDigest
		return result
	}
	if !verifySignature(relation.VerifierA, payload, proposal.SignatureA) {
		result := refusal(CodeSignatureAInvalid, state, proposal.RelationID)
		result.OperationHash = proposalDigest
		return result
	}
	if !verifySignature(relation.VerifierB, payload, proposal.SignatureB) {
		result := refusal(CodeSignatureBInvalid, state, proposal.RelationID)
		result.OperationHash = proposalDigest
		return result
	}
	nextRelation, advanced := transitionCrossingRelation(relation, proposal)
	if !advanced {
		return refusal(CodeMalformedChainState, state, proposal.RelationID)
	}
	priorStateHash, _ := RelationStateHash(relation)
	nextStateHash, _ := RelationStateHash(nextRelation)
	var handoffCommitment *string
	if proposal.Resolution == "↻" {
		digest, hashErr := HandoffHash(*proposal.HandoffManifest)
		if hashErr != nil {
			return refusal(CodeHandoffRuleViolation, state, proposal.RelationID)
		}
		handoffCommitment = &digest
	}
	closureHash, hashErr := CrossingClosureHash(priorStateHash, proposalDigest, proposal.Resolution, nextStateHash, handoffCommitment)
	if hashErr != nil {
		return refusal(CodeMalformedProposal, state, proposal.RelationID)
	}
	closure := ClosureRecord{
		ClosureHash:       closureHash,
		RelationID:        relation.RelationID,
		Epoch:             relation.Epoch,
		PriorStateHash:    priorStateHash,
		ProposalHash:      proposalDigest,
		Kind:              proposal.Resolution,
		NextStateHash:     nextStateHash,
		HandoffCommitment: handoffCommitment,
	}
	next := cloneChainState(state)
	next.Relations[index] = nextRelation
	next.Closures = insertClosure(next.Closures, closure)
	if err := ValidateChainState(next); err != nil {
		return refusal(CodeMalformedChainState, state, proposal.RelationID)
	}
	return ApplyResult{
		Accepted:          true,
		Code:              CodeAccepted,
		State:             next,
		RelationID:        proposal.RelationID,
		OperationHash:     proposalDigest,
		RelationStateHash: nextStateHash,
		ClosureHash:       closureHash,
	}
}

func ApplySuccessor(state ChainState, operation SuccessorAdmission) ApplyResult {
	relationID := operation.Admission.RelationID
	if err := ValidateChainState(state); err != nil {
		return refusal(CodeMalformedChainState, state, "")
	}
	if err := ValidateSuccessorAdmission(operation); err != nil {
		return refusal(CodeMalformedSuccessor, state, "")
	}
	if operation.Admission.Version != Version {
		return refusal(CodeUnsupportedVersion, state, relationID)
	}
	closurePosition, found := closureIndex(state.Closures, operation.PredecessorClosureHash)
	if !found {
		return refusal(CodePredecessorClosureNotFound, state, relationID)
	}
	closure := state.Closures[closurePosition]
	if closure.Kind != "↻" || closure.HandoffCommitment == nil {
		return refusal(CodePredecessorNotRenew, state, relationID)
	}
	if _, consumed := successorIndex(state.Successors, closure.ClosureHash); consumed {
		return refusal(CodePredecessorAlreadyConsumed, state, relationID)
	}
	if _, exists := relationIndex(state.Relations, relationID); exists {
		return refusal(CodeRelationAlreadyExists, state, relationID)
	}
	predecessorPosition, found := relationIndex(state.Relations, closure.RelationID)
	if !found {
		return refusal(CodeMalformedChainState, state, relationID)
	}
	predecessor := state.Relations[predecessorPosition]
	predecessorHash, err := RelationStateHash(predecessor)
	if err != nil || predecessor.Status != "ENDED" || predecessorHash != closure.NextStateHash {
		return refusal(CodeMalformedChainState, state, relationID)
	}
	admission := operation.Admission
	if admission.ProfileCommitment != predecessor.ProfileCommitment || admission.ProfileCommitment != state.ProfileCommitment {
		return refusal(CodeProfileMismatch, state, relationID)
	}
	if admission.PredecessorCommitment != closure.ClosureHash || operation.PredecessorClosureHash != closure.ClosureHash {
		return refusal(CodePredecessorCommitmentMismatch, state, relationID)
	}
	manifestHash, err := HandoffHash(operation.HandoffManifest)
	if err != nil || !ManifestRule(predecessor.HandoffDomain, operation.HandoffManifest) || admission.HandoffCommitment != manifestHash || *closure.HandoffCommitment != manifestHash {
		return refusal(CodeHandoffCommitmentMismatch, state, relationID)
	}
	if admission.InputCommitment != operation.HandoffManifest.SuccessorInputCommitment {
		return refusal(CodeFreshInputMismatch, state, relationID)
	}
	surfaceHash, err := FreshSurfaceHash(admission)
	if err != nil || surfaceHash != operation.HandoffManifest.SuccessorAdmissionCommitment {
		return refusal(CodeSuccessorAdmissionMismatch, state, relationID)
	}
	if sameVerifierKey(admission.VerifierA, admission.VerifierB) {
		return refusal(CodeVerifierScopesNotDistinct, state, relationID)
	}
	payload, err := CanonicalFreshAdmissionBytes(admission)
	if err != nil {
		return refusal(CodeMalformedSuccessor, state, relationID)
	}
	if !verifySignature(admission.VerifierA, payload, admission.SignatureA) {
		result := refusal(CodeFreshSignatureAInvalid, state, relationID)
		result.OperationHash = hashBytes(payload)
		return result
	}
	if !verifySignature(admission.VerifierB, payload, admission.SignatureB) {
		result := refusal(CodeFreshSignatureBInvalid, state, relationID)
		result.OperationHash = hashBytes(payload)
		return result
	}
	admissionHash := hashBytes(payload)
	successor := RelationState{
		Version:           Version,
		RelationID:        admission.RelationID,
		ProfileCommitment: admission.ProfileCommitment,
		Epoch:             0,
		Status:            "ACTIVE",
		HeadA:             admission.HeadA,
		VerifierA:         admission.VerifierA,
		NonceA:            0,
		HeadB:             admission.HeadB,
		VerifierB:         admission.VerifierB,
		NonceB:            0,
		HandoffDomain:     SuccessorDomain(operation.HandoffManifest),
	}
	link := SuccessorLink{
		PredecessorClosureHash: closure.ClosureHash,
		AdmissionHash:          admissionHash,
		RelationID:             successor.RelationID,
	}
	successorStateHash, err := RelationStateHash(successor)
	if err != nil {
		return refusal(CodeMalformedSuccessor, state, relationID)
	}
	next := cloneChainState(state)
	next.Relations = insertRelation(next.Relations, successor)
	next.Successors = insertSuccessor(next.Successors, link)
	if err := ValidateChainState(next); err != nil {
		return refusal(CodeMalformedChainState, state, relationID)
	}
	return ApplyResult{
		Accepted:          true,
		Code:              CodeSuccessorAccepted,
		State:             next,
		RelationID:        relationID,
		OperationHash:     admissionHash,
		RelationStateHash: successorStateHash,
	}
}

func ApplyWithdrawal(state ChainState, operation WithdrawChamber) ApplyResult {
	if err := ValidateChainState(state); err != nil {
		return refusal(CodeMalformedChainState, state, "")
	}
	if err := ValidateWithdrawal(operation); err != nil {
		return refusal(CodeMalformedWithdrawal, state, "")
	}
	if operation.Version != Version {
		return refusal(CodeUnsupportedVersion, state, operation.RelationID)
	}
	index, found := relationIndex(state.Relations, operation.RelationID)
	if !found {
		return refusal(CodeRelationNotFound, state, operation.RelationID)
	}
	relation := state.Relations[index]
	if operation.ProfileCommitment != relation.ProfileCommitment {
		return refusal(CodeProfileMismatch, state, operation.RelationID)
	}
	if relation.Status == "ENDED" {
		return refusal(CodeRelationEnded, state, operation.RelationID)
	}
	if operation.Epoch != relation.Epoch {
		return refusal(CodeStaleEpoch, state, operation.RelationID)
	}
	if operation.PrevHeadA != relation.HeadA {
		return refusal(CodeStaleHeadA, state, operation.RelationID)
	}
	if operation.PrevHeadB != relation.HeadB {
		return refusal(CodeStaleHeadB, state, operation.RelationID)
	}
	if operation.Side != "A" && operation.Side != "B" {
		return refusal(CodeSideInvalid, state, operation.RelationID)
	}
	if operation.Side == "A" && !isNext(relation.NonceA, operation.Nonce) {
		return refusal(CodeWithdrawNonceMismatch, state, operation.RelationID)
	}
	if operation.Side == "B" && !isNext(relation.NonceB, operation.Nonce) {
		return refusal(CodeWithdrawNonceMismatch, state, operation.RelationID)
	}
	payload, err := CanonicalWithdrawalBytes(operation)
	if err != nil {
		return refusal(CodeMalformedWithdrawal, state, operation.RelationID)
	}
	withdrawalHash := hashBytes(payload)
	verifier := relation.VerifierA
	if operation.Side == "B" {
		verifier = relation.VerifierB
	}
	if !verifySignature(verifier, payload, operation.Signature) {
		result := refusal(CodeWithdrawSignatureInvalid, state, operation.RelationID)
		result.OperationHash = withdrawalHash
		return result
	}
	nextRelation, advanced := transitionWithdrawalRelation(relation, operation)
	if !advanced {
		return refusal(CodeMalformedChainState, state, operation.RelationID)
	}
	priorStateHash, _ := RelationStateHash(relation)
	nextStateHash, _ := RelationStateHash(nextRelation)
	closureHash, err := WithdrawalClosureHash(priorStateHash, withdrawalHash, operation.Side, nextStateHash)
	if err != nil {
		return refusal(CodeMalformedWithdrawal, state, operation.RelationID)
	}
	closure := ClosureRecord{
		ClosureHash:    closureHash,
		RelationID:     relation.RelationID,
		Epoch:          relation.Epoch,
		PriorStateHash: priorStateHash,
		ProposalHash:   withdrawalHash,
		Kind:           "WITHDRAW_" + operation.Side,
		NextStateHash:  nextStateHash,
	}
	next := cloneChainState(state)
	next.Relations[index] = nextRelation
	next.Closures = insertClosure(next.Closures, closure)
	if err := ValidateChainState(next); err != nil {
		return refusal(CodeMalformedChainState, state, operation.RelationID)
	}
	return ApplyResult{
		Accepted:          true,
		Code:              CodeWithdrawalAccepted,
		State:             next,
		RelationID:        operation.RelationID,
		OperationHash:     withdrawalHash,
		RelationStateHash: nextStateHash,
		ClosureHash:       closureHash,
	}
}

func isNext(current, candidate uint64) bool {
	return current < MaxCounterValue && candidate == current+1
}

func DecodeSignature(value string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil || len(decoded) != ed25519.SignatureSize || base64.StdEncoding.EncodeToString(decoded) != value {
		return nil, errors.New("invalid canonical Ed25519 signature")
	}
	return decoded, nil
}
