package kernel

import "encoding/json"

const (
	ProtocolVersion = "CONTINUITY_L1_001"
	Version         = "002"
	MaxCounterValue = uint64(1<<53 - 1)
)

type Code string

const (
	CodeAccepted                      Code = "ACCEPTED"
	CodeOpenAccepted                  Code = "OPEN_ACCEPTED"
	CodeSuccessorAccepted             Code = "SUCCESSOR_ACCEPTED"
	CodeWithdrawalAccepted            Code = "WITHDRAW_ACCEPTED"
	CodeMalformedChainState           Code = "MALFORMED_CHAIN_STATE"
	CodeMalformedProposal             Code = "MALFORMED_PROPOSAL"
	CodeMalformedOpen                 Code = "MALFORMED_OPEN"
	CodeMalformedSuccessor            Code = "MALFORMED_SUCCESSOR"
	CodeMalformedWithdrawal           Code = "MALFORMED_WITHDRAWAL"
	CodeUnsupportedVersion            Code = "UNSUPPORTED_VERSION"
	CodeProfileMismatch               Code = "PROFILE_MISMATCH"
	CodeRelationNotFound              Code = "RELATION_NOT_FOUND"
	CodeRelationAlreadyExists         Code = "RELATION_EXISTS"
	CodeRelationEnded                 Code = "RELATION_ENDED"
	CodeVerifierScopesNotDistinct     Code = "VERIFIERS_NOT_DISTINCT"
	CodeStaleEpoch                    Code = "STALE_EPOCH"
	CodeStaleHeadA                    Code = "STALE_HEAD_A"
	CodeStaleHeadB                    Code = "STALE_HEAD_B"
	CodeNonceAMismatch                Code = "NONCE_A_MISMATCH"
	CodeNonceBMismatch                Code = "NONCE_B_MISMATCH"
	CodeResolutionInvalid             Code = "RESOLUTION_INVALID"
	CodeOutputRuleViolation           Code = "OUTPUT_RULE_VIOLATION"
	CodeHandoffRuleViolation          Code = "HANDOFF_RULE_VIOLATION"
	CodeSignatureAInvalid             Code = "SIGNATURE_A_INVALID"
	CodeSignatureBInvalid             Code = "SIGNATURE_B_INVALID"
	CodeOpenSignatureAInvalid         Code = "OPEN_SIGNATURE_A_INVALID"
	CodeOpenSignatureBInvalid         Code = "OPEN_SIGNATURE_B_INVALID"
	CodePredecessorClosureNotFound    Code = "PREDECESSOR_CLOSURE_NOT_FOUND"
	CodePredecessorNotRenew           Code = "PREDECESSOR_NOT_RENEW"
	CodePredecessorAlreadyConsumed    Code = "PREDECESSOR_ALREADY_CONSUMED"
	CodePredecessorCommitmentMismatch Code = "PREDECESSOR_COMMITMENT_MISMATCH"
	CodeHandoffCommitmentMismatch     Code = "HANDOFF_COMMITMENT_MISMATCH"
	CodeFreshInputMismatch            Code = "FRESH_INPUT_MISMATCH"
	CodeSuccessorAdmissionMismatch    Code = "SUCCESSOR_ADMISSION_COMMITMENT_MISMATCH"
	CodeFreshSignatureAInvalid        Code = "FRESH_SIGNATURE_A_INVALID"
	CodeFreshSignatureBInvalid        Code = "FRESH_SIGNATURE_B_INVALID"
	CodeSideInvalid                   Code = "SIDE_INVALID"
	CodeWithdrawNonceMismatch         Code = "WITHDRAW_NONCE_MISMATCH"
	CodeWithdrawSignatureInvalid      Code = "WITHDRAW_SIGNATURE_INVALID"
)

var transitionPrecedence = [...]Code{
	CodeMalformedChainState,
	CodeMalformedProposal,
	CodeUnsupportedVersion,
	CodeRelationNotFound,
	CodeProfileMismatch,
	CodeRelationEnded,
	CodeStaleEpoch,
	CodeStaleHeadA,
	CodeStaleHeadB,
	CodeNonceAMismatch,
	CodeNonceBMismatch,
	CodeResolutionInvalid,
	CodeOutputRuleViolation,
	CodeHandoffRuleViolation,
	CodeSignatureAInvalid,
	CodeSignatureBInvalid,
}

var openPrecedence = [...]Code{
	CodeMalformedChainState,
	CodeMalformedOpen,
	CodeUnsupportedVersion,
	CodeProfileMismatch,
	CodeRelationAlreadyExists,
	CodeVerifierScopesNotDistinct,
	CodeOpenSignatureAInvalid,
	CodeOpenSignatureBInvalid,
}

var successorPrecedence = [...]Code{
	CodeMalformedChainState,
	CodeMalformedSuccessor,
	CodeUnsupportedVersion,
	CodePredecessorClosureNotFound,
	CodePredecessorNotRenew,
	CodePredecessorAlreadyConsumed,
	CodeRelationAlreadyExists,
	CodeProfileMismatch,
	CodePredecessorCommitmentMismatch,
	CodeHandoffCommitmentMismatch,
	CodeFreshInputMismatch,
	CodeSuccessorAdmissionMismatch,
	CodeVerifierScopesNotDistinct,
	CodeFreshSignatureAInvalid,
	CodeFreshSignatureBInvalid,
}

var withdrawalPrecedence = [...]Code{
	CodeMalformedChainState,
	CodeMalformedWithdrawal,
	CodeUnsupportedVersion,
	CodeRelationNotFound,
	CodeProfileMismatch,
	CodeRelationEnded,
	CodeStaleEpoch,
	CodeStaleHeadA,
	CodeStaleHeadB,
	CodeSideInvalid,
	CodeWithdrawNonceMismatch,
	CodeWithdrawSignatureInvalid,
}

func copyPrecedence(values []Code) []Code {
	copyOfValues := make([]Code, len(values))
	copy(copyOfValues, values)
	return copyOfValues
}

func TransitionPrecedence() []Code {
	return copyPrecedence(transitionPrecedence[:])
}

func OpenPrecedence() []Code {
	return copyPrecedence(openPrecedence[:])
}

func SuccessorPrecedence() []Code {
	return copyPrecedence(successorPrecedence[:])
}

func WithdrawalPrecedence() []Code {
	return copyPrecedence(withdrawalPrecedence[:])
}

type ChainConfig struct {
	ProtocolVersion   string `json:"protocol_version"`
	KernelVersion     string `json:"kernel_version"`
	ProfileCommitment string `json:"profile_commitment"`
}

type ChainState struct {
	ProtocolVersion   string          `json:"protocol_version"`
	KernelVersion     string          `json:"kernel_version"`
	ProfileCommitment string          `json:"profile_commitment"`
	Relations         []RelationState `json:"relations"`
	Closures          []ClosureRecord `json:"closures"`
	Successors        []SuccessorLink `json:"successors"`
}

type RelationState struct {
	Version           string   `json:"version"`
	RelationID        string   `json:"relation_id"`
	ProfileCommitment string   `json:"profile_commitment"`
	Epoch             uint64   `json:"epoch"`
	Status            string   `json:"status"`
	HeadA             string   `json:"head_a"`
	VerifierA         string   `json:"verifier_a"`
	NonceA            uint64   `json:"nonce_a"`
	HeadB             string   `json:"head_b"`
	VerifierB         string   `json:"verifier_b"`
	NonceB            uint64   `json:"nonce_b"`
	HandoffDomain     []string `json:"handoff_domain"`
}

type Transform struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type HandoffManifest struct {
	SuccessorInputCommitment     string      `json:"successor_input_commitment"`
	SuccessorAdmissionCommitment string      `json:"successor_admission_commitment"`
	Preserved                    []string    `json:"preserved"`
	Transformed                  []Transform `json:"transformed"`
	Ended                        []string    `json:"ended"`
	Refused                      []string    `json:"refused"`
	Unresolved                   []string    `json:"unresolved"`
}

type CrossingProposal struct {
	Version           string           `json:"version"`
	RelationID        string           `json:"relation_id"`
	ProfileCommitment string           `json:"profile_commitment"`
	Epoch             uint64           `json:"epoch"`
	PrevHeadA         string           `json:"prev_head_a"`
	PrevHeadB         string           `json:"prev_head_b"`
	Resolution        string           `json:"resolution"`
	NextHeadA         *string          `json:"next_head_a"`
	NextHeadB         *string          `json:"next_head_b"`
	NonceA            uint64           `json:"nonce_a"`
	NonceB            uint64           `json:"nonce_b"`
	HandoffManifest   *HandoffManifest `json:"handoff_manifest"`
	SignatureA        string           `json:"signature_a"`
	SignatureB        string           `json:"signature_b"`
}

type OpenRelation struct {
	Version           string   `json:"version"`
	RelationID        string   `json:"relation_id"`
	ProfileCommitment string   `json:"profile_commitment"`
	InvocationID      string   `json:"invocation_id"`
	InputCommitment   string   `json:"input_commitment"`
	BoundsCommitment  string   `json:"bounds_commitment"`
	HeadA             string   `json:"head_a"`
	VerifierA         string   `json:"verifier_a"`
	HeadB             string   `json:"head_b"`
	VerifierB         string   `json:"verifier_b"`
	HandoffDomain     []string `json:"handoff_domain"`
	SignatureA        string   `json:"signature_a"`
	SignatureB        string   `json:"signature_b"`
}

type FreshAdmission struct {
	Version               string `json:"version"`
	RelationID            string `json:"relation_id"`
	ProfileCommitment     string `json:"profile_commitment"`
	InvocationID          string `json:"invocation_id"`
	PredecessorCommitment string `json:"predecessor_commitment"`
	HandoffCommitment     string `json:"handoff_commitment"`
	InputCommitment       string `json:"input_commitment"`
	BoundsCommitment      string `json:"bounds_commitment"`
	HeadA                 string `json:"head_a"`
	VerifierA             string `json:"verifier_a"`
	HeadB                 string `json:"head_b"`
	VerifierB             string `json:"verifier_b"`
	SignatureA            string `json:"signature_a"`
	SignatureB            string `json:"signature_b"`
}

type SuccessorAdmission struct {
	PredecessorClosureHash string          `json:"predecessor_closure_hash"`
	HandoffManifest       HandoffManifest `json:"handoff_manifest"`
	Admission             FreshAdmission  `json:"fresh_admission"`
}

type WithdrawChamber struct {
	Version           string `json:"version"`
	RelationID        string `json:"relation_id"`
	ProfileCommitment string `json:"profile_commitment"`
	Epoch             uint64 `json:"epoch"`
	PrevHeadA         string `json:"head_a"`
	PrevHeadB         string `json:"head_b"`
	Side              string `json:"side"`
	Nonce             uint64 `json:"nonce"`
	Signature         string `json:"signature"`
}

type ClosureRecord struct {
	ClosureHash       string  `json:"closure_hash"`
	RelationID        string  `json:"relation_id"`
	Epoch             uint64  `json:"epoch"`
	PriorStateHash    string  `json:"prior_state_hash"`
	ProposalHash      string  `json:"proposal_hash"`
	Kind              string  `json:"resolution"`
	NextStateHash     string  `json:"next_state_hash"`
	HandoffCommitment *string `json:"handoff_commitment"`
}

type SuccessorLink struct {
	PredecessorClosureHash string `json:"predecessor_closure_hash"`
	AdmissionHash          string `json:"admission_hash"`
	RelationID             string `json:"relation_id"`
}

type ApplyResult struct {
	Accepted          bool       `json:"accepted"`
	Code              Code       `json:"code"`
	State             ChainState `json:"state"`
	RelationID        string     `json:"-"`
	OperationHash     string     `json:"-"`
	RelationStateHash string     `json:"-"`
	ClosureHash       string     `json:"-"`
}

func nullableResultString(value string) *string {
	if value == "" {
		return nil
	}
	copy := value
	return &copy
}

func (value ApplyResult) MarshalJSON() ([]byte, error) {
	wire := struct {
		Accepted          bool        `json:"accepted"`
		Code              Code        `json:"code"`
		OperationHash     *string     `json:"operation_hash"`
		RelationID        *string     `json:"relation_id"`
		RelationStateHash *string     `json:"relation_state_hash"`
		ClosureHash       *string     `json:"closure_hash"`
		State             ChainState  `json:"state"`
	}{
		Accepted:          value.Accepted,
		Code:              value.Code,
		OperationHash:     nullableResultString(value.OperationHash),
		RelationID:        nullableResultString(value.RelationID),
		RelationStateHash: nullableResultString(value.RelationStateHash),
		ClosureHash:       nullableResultString(value.ClosureHash),
		State:             value.State,
	}
	return json.Marshal(wire)
}
