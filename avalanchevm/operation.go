package avalanchevm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ia-mm-ai/THIRD-1/kernel"
)

const operationIDDomain = "CONTINUITY:OPERATION:001"

const codeStateBoundExceeded kernel.Code = "ADAPTER_STATE_BOUND_EXCEEDED"

type operationEnvelope struct {
	Schema  string          `json:"schema"`
	Version uint16          `json:"version"`
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

type Operation struct {
	Action     string
	Open       *kernel.OpenRelation
	Crossing   *kernel.CrossingProposal
	Successor  *kernel.SuccessorAdmission
	Withdrawal *kernel.WithdrawChamber
	raw        []byte
}

func (operation *Operation) Bytes() []byte {
	if operation == nil {
		return nil
	}
	return bytes.Clone(operation.raw)
}

func (operation *Operation) ID() string { return domainHash(operationIDDomain, operation.raw) }

func (operation *Operation) MarshalJSON() ([]byte, error) {
	if operation == nil || len(operation.raw) == 0 {
		return nil, errors.New("cannot encode an empty Continuity operation")
	}
	return bytes.Clone(operation.raw), nil
}

func (operation *Operation) UnmarshalJSON(data []byte) error {
	parsed, _, err := parseOperation(data)
	if err != nil {
		return err
	}
	*operation = *parsed
	return nil
}

func NewOperation(action string, payload any) (*Operation, error) {
	payloadBytes, err := canonicalJSON(payload)
	if err != nil {
		return nil, err
	}
	data, err := canonicalJSON(operationEnvelope{
		Schema: OperationSchema, Version: 1, Action: action, Payload: payloadBytes,
	})
	if err != nil {
		return nil, err
	}
	parsed, _, err := parseOperation(data)
	return parsed, err
}

func parseOperation(data []byte) (*Operation, string, error) {
	if len(data) == 0 || uint64(len(data)) > uint64(maxOperationBytes) {
		return nil, "", errors.New("operation byte length is outside compiled bounds")
	}
	envelope := new(operationEnvelope)
	if err := decodeCanonical(data, envelope); err != nil {
		return nil, "", err
	}
	if envelope.Schema != OperationSchema || envelope.Version != 1 || len(envelope.Payload) == 0 {
		return nil, "", errors.New("unsupported Continuity operation envelope")
	}
	operation := &Operation{Action: envelope.Action, raw: bytes.Clone(data)}
	switch envelope.Action {
	case ActionOpenRelation:
		value := new(kernel.OpenRelation)
		if err := decodeCanonical(envelope.Payload, value); err != nil {
			return nil, "", fmt.Errorf("open relation payload: %w", err)
		}
		if err := kernel.ValidateOpenRelation(*value); err != nil {
			return nil, "", err
		}
		operation.Open = value
	case ActionApplyCrossing:
		value := new(kernel.CrossingProposal)
		if err := decodeCanonical(envelope.Payload, value); err != nil {
			return nil, "", fmt.Errorf("crossing payload: %w", err)
		}
		if err := kernel.ValidateCrossingProposal(*value); err != nil {
			return nil, "", err
		}
		operation.Crossing = value
	case ActionAdmitSuccessor:
		value := new(kernel.SuccessorAdmission)
		if err := decodeCanonical(envelope.Payload, value); err != nil {
			return nil, "", fmt.Errorf("successor payload: %w", err)
		}
		if err := kernel.ValidateSuccessorAdmission(*value); err != nil {
			return nil, "", err
		}
		operation.Successor = value
	case ActionWithdraw:
		value := new(kernel.WithdrawChamber)
		if err := decodeCanonical(envelope.Payload, value); err != nil {
			return nil, "", fmt.Errorf("withdrawal payload: %w", err)
		}
		if err := kernel.ValidateWithdrawal(*value); err != nil {
			return nil, "", err
		}
		operation.Withdrawal = value
	default:
		return nil, "", fmt.Errorf("unsupported Continuity action %q", envelope.Action)
	}
	return operation, operation.ID(), nil
}

type refusalError struct{ code kernel.Code }

func (err *refusalError) Error() string { return "kernel refused operation: " + string(err.code) }

func refusalCode(err error) (kernel.Code, bool) {
	target := new(refusalError)
	if !errors.As(err, &target) {
		return "", false
	}
	return target.code, true
}

func applyOperation(state kernel.ChainState, operation *Operation) (kernel.ApplyResult, error) {
	var result kernel.ApplyResult
	switch operation.Action {
	case ActionOpenRelation:
		result = kernel.ApplyOpen(state, *operation.Open)
	case ActionApplyCrossing:
		result = kernel.ApplyCrossing(state, *operation.Crossing)
	case ActionAdmitSuccessor:
		result = kernel.ApplySuccessor(state, *operation.Successor)
	case ActionWithdraw:
		result = kernel.ApplyWithdrawal(state, *operation.Withdrawal)
	default:
		return kernel.ApplyResult{}, errors.New("operation has no executable action")
	}
	if !result.Accepted {
		return result, &refusalError{code: result.Code}
	}
	if err := kernel.ValidateChainState(result.State); err != nil {
		return kernel.ApplyResult{}, fmt.Errorf("kernel emitted invalid state: %w", err)
	}
	return result, nil
}

func (vm *VM) applyBoundedOperation(state kernel.ChainState, operation *Operation) (kernel.ApplyResult, error) {
	result, err := applyOperation(state, operation)
	if err != nil {
		return result, err
	}
	if _, err := encodeBoundedState(result.State, vm.genesis.Limits); err != nil {
		return result, &refusalError{code: codeStateBoundExceeded}
	}
	return result, nil
}

type Receipt struct {
	Schema            string      `json:"schema"`
	Version           uint16      `json:"version"`
	OperationID       string      `json:"operation_id"`
	Action            string      `json:"action"`
	Code              kernel.Code `json:"code"`
	RelationID        *string     `json:"relation_id"`
	OperationHash     *string     `json:"operation_hash"`
	RelationStateHash *string     `json:"relation_state_hash"`
	ClosureHash       *string     `json:"closure_hash"`
	StateHash         string      `json:"state_hash"`
	BlockID           string      `json:"block_id"`
	BlockHeight       uint64      `json:"block_height"`
	OperationIndex    uint16      `json:"operation_index"`
}

func nullableReceiptString(value string) *string {
	if value == "" {
		return nil
	}
	copy := value
	return &copy
}

func receiptFor(operation *Operation, result kernel.ApplyResult, blockID string, height uint64, index uint16) (*Receipt, error) {
	stateHash, err := kernel.ChainStateHash(result.State)
	if err != nil {
		return nil, err
	}
	return &Receipt{
		Schema: ReceiptSchema, Version: 1, OperationID: operation.ID(), Action: operation.Action,
		Code: result.Code, RelationID: nullableReceiptString(result.RelationID),
		OperationHash: nullableReceiptString(result.OperationHash),
		RelationStateHash: nullableReceiptString(result.RelationStateHash),
		ClosureHash: nullableReceiptString(result.ClosureHash), StateHash: stateHash,
		BlockID: blockID, BlockHeight: height, OperationIndex: index,
	}, nil
}
