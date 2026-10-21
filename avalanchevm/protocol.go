package avalanchevm

import "errors"

const (
	ImplementationID      = "CONTINUITY_L1_CANDIDATE"
	ImplementationVersion = "0.1.0"
	VMVersion             = "continuity-l1/0.1.0"
	AvalancheGoTarget     = "v1.15.0"
	RequiredRPCProtocol   = uint(46)

	GenesisSchema   = "CONTINUITY_GENESIS_001"
	ProfileDomain   = "CONTINUITY:AVALANCHE_PROFILE_DOMAIN:002"
	OperationSchema = "CONTINUITY_OPERATION_001"
	BlockSchema     = "CONTINUITY_BLOCK_001"
	ReceiptSchema   = "CONTINUITY_RECEIPT_001"

	ActionOpenRelation   = "OPEN_RELATION"
	ActionApplyCrossing  = "APPLY_CROSSING"
	ActionAdmitSuccessor = "ADMIT_SUCCESSOR"
	ActionWithdraw       = "WITHDRAW_CHAMBER"

	maxOperationBytes     = uint32(1 << 20)
	maxBlockBytes         = uint32(2 << 20)
	maxOperationsPerBlock = uint16(1)
	maxAcceptedStateBytes = uint32(16 << 20)
	maxRelations          = uint32(4096)
	maxClosures           = uint32(8192)
	maxSuccessors         = uint32(4096)
	maxPendingOperations  = uint32(1024)
	maxPendingBytes       = uint64(64 << 20)
)

var (
	errNotInitialized      = errors.New("Continuity VM is not initialized")
	errNoPendingOperations = errors.New("no valid pending operations")
	errInvalidPreference   = errors.New("preferred block is unavailable or unverified")
	errNonCanonicalBlock   = errors.New("block object is not the canonical reference for its ID")
	errCannotRejectGenesis = errors.New("the definitionally accepted genesis block cannot be rejected")
)

type Limits struct {
	MaxOperationBytes     uint32 `json:"max_operation_bytes"`
	MaxBlockBytes         uint32 `json:"max_block_bytes"`
	MaxOperationsPerBlock uint16 `json:"max_operations_per_block"`
	MaxAcceptedStateBytes uint32 `json:"max_accepted_state_bytes"`
	MaxRelations          uint32 `json:"max_relations"`
	MaxClosures           uint32 `json:"max_closures"`
	MaxSuccessors         uint32 `json:"max_successors"`
	MaxPendingOperations  uint32 `json:"max_pending_operations"`
	MaxPendingBytes       uint64 `json:"max_pending_bytes"`
}

func fixedLimits() Limits {
	return Limits{
		MaxOperationBytes:     maxOperationBytes,
		MaxBlockBytes:         maxBlockBytes,
		MaxOperationsPerBlock: maxOperationsPerBlock,
		MaxAcceptedStateBytes: maxAcceptedStateBytes,
		MaxRelations:          maxRelations,
		MaxClosures:           maxClosures,
		MaxSuccessors:         maxSuccessors,
		MaxPendingOperations:  maxPendingOperations,
		MaxPendingBytes:       maxPendingBytes,
	}
}

func (limits Limits) validate() error {
	if limits != fixedLimits() {
		return errors.New("genesis limits differ from compiled Continuity limits")
	}
	return nil
}
