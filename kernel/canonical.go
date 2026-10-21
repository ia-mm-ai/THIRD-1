package kernel

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

var (
	proposalDomain          = []byte("IK:CROSSING:002\x00")
	relationStateDomain     = []byte("IK:RELATION_STATE:002\x00")
	handoffDomain           = []byte("IK:HANDOFF:002\x00")
	closureDomain           = []byte("IK:CLOSURE:002\x00")
	freshAdmissionDomain    = []byte("IK:FRESH_ADMISSION:002\x00")
	freshSurfaceDomain      = []byte("IK:FRESH_SURFACE:002\x00")
	openRelationDomain      = []byte("IK:OPEN_RELATION:001\x00")
	withdrawalDomain        = []byte("IK:WITHDRAW_CHAMBER:001\x00")
	withdrawalClosureDomain = []byte("IK:WITHDRAW_CLOSURE:001\x00")
	chainStateDomain        = []byte("IK:CHAIN_STATE:001\x00")
	closureRecordDomain     = []byte("IK:CLOSURE_RECORD:001\x00")
	successorLinkDomain     = []byte("IK:SUCCESSOR_LINK:001\x00")
)

func appendU32(dst []byte, n uint32) []byte {
	var value [4]byte
	binary.BigEndian.PutUint32(value[:], n)
	return append(dst, value[:]...)
}

func appendU64(dst []byte, n uint64) []byte {
	var value [8]byte
	binary.BigEndian.PutUint64(value[:], n)
	return append(dst, value[:]...)
}

func appendBytes(dst, value []byte) ([]byte, error) {
	if uint64(len(value)) > uint64(^uint32(0)) {
		return nil, errors.New("byte string exceeds u32 length")
	}
	dst = appendU32(dst, uint32(len(value)))
	return append(dst, value...), nil
}

func appendText(dst []byte, value string) ([]byte, error) {
	return appendBytes(dst, []byte(value))
}

func decodeHash(value string) ([]byte, error) {
	if !validHash(value) {
		return nil, errors.New("expected 32-byte lowercase hexadecimal hash")
	}
	return hex.DecodeString(value)
}

func appendHash(dst []byte, value string) ([]byte, error) {
	decoded, err := decodeHash(value)
	if err != nil {
		return nil, err
	}
	return append(dst, decoded...), nil
}

func appendOptionalHash(dst []byte, value *string) ([]byte, error) {
	if value == nil {
		return append(dst, 0), nil
	}
	dst = append(dst, 1)
	return appendHash(dst, *value)
}

func appendHashList(dst []byte, values []string) ([]byte, error) {
	if uint64(len(values)) > uint64(^uint32(0)) {
		return nil, errors.New("hash list exceeds u32 length")
	}
	dst = appendU32(dst, uint32(len(values)))
	var err error
	for _, value := range values {
		dst, err = appendHash(dst, value)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func appendVerifier(dst []byte, value string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil || base64.StdEncoding.EncodeToString(decoded) != value {
		return nil, errors.New("verifier is not canonical base64")
	}
	return appendBytes(dst, decoded)
}

func resolutionByte(value string) (byte, error) {
	switch value {
	case "1":
		return 1, nil
	case "0":
		return 0, nil
	case "↻":
		return 2, nil
	default:
		return 0, errors.New("invalid resolution")
	}
}

func sideByte(value string) (byte, error) {
	switch value {
	case "A":
		return 0, nil
	case "B":
		return 1, nil
	default:
		return 0, errors.New("invalid chamber side")
	}
}

func hashBytes(value []byte) string {
	digest := sha256.Sum256(value)
	return hex.EncodeToString(digest[:])
}

func CanonicalHandoffBytes(value HandoffManifest) ([]byte, error) {
	dst := append([]byte(nil), handoffDomain...)
	var err error
	if dst, err = appendHash(dst, value.SuccessorInputCommitment); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.SuccessorAdmissionCommitment); err != nil {
		return nil, err
	}
	if dst, err = appendHashList(dst, value.Preserved); err != nil {
		return nil, err
	}
	if uint64(len(value.Transformed)) > uint64(^uint32(0)) {
		return nil, errors.New("transform list exceeds u32 length")
	}
	dst = appendU32(dst, uint32(len(value.Transformed)))
	for _, transform := range value.Transformed {
		if dst, err = appendHash(dst, transform.From); err != nil {
			return nil, err
		}
		if dst, err = appendHash(dst, transform.To); err != nil {
			return nil, err
		}
	}
	if dst, err = appendHashList(dst, value.Ended); err != nil {
		return nil, err
	}
	if dst, err = appendHashList(dst, value.Refused); err != nil {
		return nil, err
	}
	return appendHashList(dst, value.Unresolved)
}

func HandoffHash(value HandoffManifest) (string, error) {
	encoded, err := CanonicalHandoffBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CanonicalRelationStateBytes(value RelationState) ([]byte, error) {
	dst := append([]byte(nil), relationStateDomain...)
	var err error
	if dst, err = appendText(dst, value.Version); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.RelationID); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.ProfileCommitment); err != nil {
		return nil, err
	}
	dst = appendU64(dst, value.Epoch)
	switch value.Status {
	case "ACTIVE":
		dst = append(dst, 1)
	case "ENDED":
		dst = append(dst, 0)
	default:
		return nil, errors.New("invalid relation status")
	}
	if dst, err = appendHash(dst, value.HeadA); err != nil {
		return nil, err
	}
	if dst, err = appendVerifier(dst, value.VerifierA); err != nil {
		return nil, err
	}
	dst = appendU64(dst, value.NonceA)
	if dst, err = appendHash(dst, value.HeadB); err != nil {
		return nil, err
	}
	if dst, err = appendVerifier(dst, value.VerifierB); err != nil {
		return nil, err
	}
	dst = appendU64(dst, value.NonceB)
	return appendHashList(dst, value.HandoffDomain)
}

func RelationStateHash(value RelationState) (string, error) {
	encoded, err := CanonicalRelationStateBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CanonicalProposalBytes(value CrossingProposal) ([]byte, error) {
	dst := append([]byte(nil), proposalDomain...)
	var err error
	if dst, err = appendText(dst, value.Version); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.RelationID); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.ProfileCommitment); err != nil {
		return nil, err
	}
	dst = appendU64(dst, value.Epoch)
	if dst, err = appendHash(dst, value.PrevHeadA); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.PrevHeadB); err != nil {
		return nil, err
	}
	resolution, err := resolutionByte(value.Resolution)
	if err != nil {
		return nil, err
	}
	dst = append(dst, resolution)
	if dst, err = appendOptionalHash(dst, value.NextHeadA); err != nil {
		return nil, err
	}
	if dst, err = appendOptionalHash(dst, value.NextHeadB); err != nil {
		return nil, err
	}
	dst = appendU64(dst, value.NonceA)
	dst = appendU64(dst, value.NonceB)
	if value.HandoffManifest == nil {
		return append(dst, 0), nil
	}
	dst = append(dst, 1)
	manifest, err := CanonicalHandoffBytes(*value.HandoffManifest)
	if err != nil {
		return nil, err
	}
	return append(dst, manifest...), nil
}

func ProposalHash(value CrossingProposal) (string, error) {
	encoded, err := CanonicalProposalBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CanonicalFreshSurfaceBytes(value FreshAdmission) ([]byte, error) {
	dst := append([]byte(nil), freshSurfaceDomain...)
	var err error
	if dst, err = appendText(dst, value.Version); err != nil {
		return nil, err
	}
	for _, hash := range []string{
		value.RelationID,
		value.ProfileCommitment,
		value.InvocationID,
		value.InputCommitment,
		value.BoundsCommitment,
		value.HeadA,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	if dst, err = appendVerifier(dst, value.VerifierA); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.HeadB); err != nil {
		return nil, err
	}
	return appendVerifier(dst, value.VerifierB)
}

func FreshSurfaceHash(value FreshAdmission) (string, error) {
	encoded, err := CanonicalFreshSurfaceBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CanonicalFreshAdmissionBytes(value FreshAdmission) ([]byte, error) {
	dst := append([]byte(nil), freshAdmissionDomain...)
	var err error
	if dst, err = appendText(dst, value.Version); err != nil {
		return nil, err
	}
	for _, hash := range []string{
		value.RelationID,
		value.ProfileCommitment,
		value.InvocationID,
		value.PredecessorCommitment,
		value.HandoffCommitment,
		value.InputCommitment,
		value.BoundsCommitment,
		value.HeadA,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	if dst, err = appendVerifier(dst, value.VerifierA); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.HeadB); err != nil {
		return nil, err
	}
	return appendVerifier(dst, value.VerifierB)
}

func FreshAdmissionHash(value FreshAdmission) (string, error) {
	encoded, err := CanonicalFreshAdmissionBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CanonicalOpenRelationBytes(value OpenRelation) ([]byte, error) {
	dst := append([]byte(nil), openRelationDomain...)
	var err error
	if dst, err = appendText(dst, value.Version); err != nil {
		return nil, err
	}
	for _, hash := range []string{
		value.RelationID,
		value.ProfileCommitment,
		value.InvocationID,
		value.InputCommitment,
		value.BoundsCommitment,
		value.HeadA,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	if dst, err = appendVerifier(dst, value.VerifierA); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.HeadB); err != nil {
		return nil, err
	}
	if dst, err = appendVerifier(dst, value.VerifierB); err != nil {
		return nil, err
	}
	return appendHashList(dst, value.HandoffDomain)
}

func OpenRelationHash(value OpenRelation) (string, error) {
	encoded, err := CanonicalOpenRelationBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CanonicalWithdrawalBytes(value WithdrawChamber) ([]byte, error) {
	dst := append([]byte(nil), withdrawalDomain...)
	var err error
	if dst, err = appendText(dst, value.Version); err != nil {
		return nil, err
	}
	for _, hash := range []string{
		value.RelationID,
		value.ProfileCommitment,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	dst = appendU64(dst, value.Epoch)
	if dst, err = appendHash(dst, value.PrevHeadA); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.PrevHeadB); err != nil {
		return nil, err
	}
	side, err := sideByte(value.Side)
	if err != nil {
		return nil, err
	}
	dst = append(dst, side)
	return appendU64(dst, value.Nonce), nil
}

func WithdrawalHash(value WithdrawChamber) (string, error) {
	encoded, err := CanonicalWithdrawalBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

func CrossingClosureHash(priorStateHash, proposalHash, resolution, nextStateHash string, handoffCommitment *string) (string, error) {
	dst := append([]byte(nil), closureDomain...)
	var err error
	if dst, err = appendHash(dst, priorStateHash); err != nil {
		return "", err
	}
	if dst, err = appendHash(dst, proposalHash); err != nil {
		return "", err
	}
	resolutionCode, err := resolutionByte(resolution)
	if err != nil {
		return "", err
	}
	dst = append(dst, resolutionCode)
	if dst, err = appendHash(dst, nextStateHash); err != nil {
		return "", err
	}
	if dst, err = appendOptionalHash(dst, handoffCommitment); err != nil {
		return "", err
	}
	return hashBytes(dst), nil
}

func WithdrawalClosureHash(priorStateHash, withdrawalHash, side, nextStateHash string) (string, error) {
	dst := append([]byte(nil), withdrawalClosureDomain...)
	var err error
	if dst, err = appendHash(dst, priorStateHash); err != nil {
		return "", err
	}
	if dst, err = appendHash(dst, withdrawalHash); err != nil {
		return "", err
	}
	sideCode, err := sideByte(side)
	if err != nil {
		return "", err
	}
	dst = append(dst, sideCode)
	if dst, err = appendHash(dst, nextStateHash); err != nil {
		return "", err
	}
	return hashBytes(dst), nil
}

func canonicalClosureRecordBytes(value ClosureRecord) ([]byte, error) {
	dst := append([]byte(nil), closureRecordDomain...)
	var err error
	for _, hash := range []string{
		value.ClosureHash,
		value.RelationID,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	dst = appendU64(dst, value.Epoch)
	for _, hash := range []string{
		value.PriorStateHash,
		value.ProposalHash,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	if dst, err = appendText(dst, value.Kind); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.NextStateHash); err != nil {
		return nil, err
	}
	return appendOptionalHash(dst, value.HandoffCommitment)
}

func canonicalSuccessorLinkBytes(value SuccessorLink) ([]byte, error) {
	dst := append([]byte(nil), successorLinkDomain...)
	var err error
	for _, hash := range []string{
		value.PredecessorClosureHash,
		value.AdmissionHash,
		value.RelationID,
	} {
		if dst, err = appendHash(dst, hash); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func appendFramedValues(dst []byte, values [][]byte) ([]byte, error) {
	if uint64(len(values)) > uint64(^uint32(0)) {
		return nil, errors.New("value list exceeds u32 length")
	}
	dst = appendU32(dst, uint32(len(values)))
	var err error
	for _, value := range values {
		dst, err = appendBytes(dst, value)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func CanonicalChainStateBytes(value ChainState) ([]byte, error) {
	if err := ValidateChainState(value); err != nil {
		return nil, err
	}
	dst := append([]byte(nil), chainStateDomain...)
	var err error
	if dst, err = appendText(dst, value.ProtocolVersion); err != nil {
		return nil, err
	}
	if dst, err = appendText(dst, value.KernelVersion); err != nil {
		return nil, err
	}
	if dst, err = appendHash(dst, value.ProfileCommitment); err != nil {
		return nil, err
	}
	relations := make([][]byte, len(value.Relations))
	for index, relation := range value.Relations {
		relations[index], err = CanonicalRelationStateBytes(relation)
		if err != nil {
			return nil, err
		}
	}
	if dst, err = appendFramedValues(dst, relations); err != nil {
		return nil, err
	}
	closures := make([][]byte, len(value.Closures))
	for index, closure := range value.Closures {
		closures[index], err = canonicalClosureRecordBytes(closure)
		if err != nil {
			return nil, err
		}
	}
	if dst, err = appendFramedValues(dst, closures); err != nil {
		return nil, err
	}
	successors := make([][]byte, len(value.Successors))
	for index, successor := range value.Successors {
		successors[index], err = canonicalSuccessorLinkBytes(successor)
		if err != nil {
			return nil, err
		}
	}
	return appendFramedValues(dst, successors)
}

func ChainStateHash(value ChainState) (string, error) {
	encoded, err := CanonicalChainStateBytes(value)
	if err != nil {
		return "", err
	}
	return hashBytes(encoded), nil
}

