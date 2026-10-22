package avalanchevm

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ia-mm-ai/THIRD-1/kernel"
)

type AvalancheSecurityDomain struct {
	NetworkID uint32 `json:"network_id"`
	SubnetID  string `json:"subnet_id"`
}

type GenesisKernelConfig struct {
	ProtocolVersion           string `json:"protocol_version"`
	KernelVersion             string `json:"kernel_version"`
	SemanticProfileCommitment string `json:"semantic_profile_commitment"`
}

type Genesis struct {
	Schema         string                  `json:"schema"`
	Version        uint16                  `json:"version"`
	SecurityDomain AvalancheSecurityDomain `json:"security_domain"`
	Config         GenesisKernelConfig     `json:"config"`
	Limits         Limits                  `json:"limits"`
}

func deriveChainProfileCommitment(
	semanticProfileCommitment string,
	networkID uint32,
	subnetID ids.ID,
	chainID ids.ID,
	kernelVersion string,
) (string, error) {
	if err := requireHash("semantic profile commitment", semanticProfileCommitment); err != nil {
		return "", err
	}
	if len(kernelVersion) == 0 || len(kernelVersion) > math.MaxUint16 {
		return "", errors.New("kernel version length is outside the profile-domain encoding")
	}
	semanticProfile, err := hex.DecodeString(semanticProfileCommitment)
	if err != nil {
		return "", err
	}
	preimage := make([]byte, 0, len(ProfileDomain)+1+sha256.Size+4+(2*ids.IDLen)+2+len(kernelVersion))
	preimage = append(preimage, ProfileDomain...)
	preimage = append(preimage, 0)
	preimage = append(preimage, semanticProfile...)
	var networkBytes [4]byte
	binary.BigEndian.PutUint32(networkBytes[:], networkID)
	preimage = append(preimage, networkBytes[:]...)
	preimage = append(preimage, subnetID[:]...)
	preimage = append(preimage, chainID[:]...)
	var versionLength [2]byte
	binary.BigEndian.PutUint16(versionLength[:], uint16(len(kernelVersion)))
	preimage = append(preimage, versionLength[:]...)
	preimage = append(preimage, kernelVersion...)
	digest := sha256.Sum256(preimage)
	return hex.EncodeToString(digest[:]), nil
}

func (genesis *Genesis) validateSecurityDomain() (ids.ID, error) {
	if genesis.SecurityDomain.SubnetID == "" {
		return ids.Empty, errors.New("security domain subnet ID is empty")
	}
	subnetID, err := ids.FromString(genesis.SecurityDomain.SubnetID)
	if err != nil {
		return ids.Empty, fmt.Errorf("parse security domain subnet ID: %w", err)
	}
	if err := requireHash("semantic profile commitment", genesis.Config.SemanticProfileCommitment); err != nil {
		return ids.Empty, err
	}
	return subnetID, nil
}

func (genesis *Genesis) chainConfigForContext(chainContext *snow.Context) (kernel.ChainConfig, error) {
	if chainContext == nil {
		return kernel.ChainConfig{}, errors.New("Avalanche chain context is required")
	}
	subnetID, err := genesis.validateSecurityDomain()
	if err != nil {
		return kernel.ChainConfig{}, err
	}
	if genesis.SecurityDomain.NetworkID != chainContext.NetworkID {
		return kernel.ChainConfig{}, fmt.Errorf("genesis network ID %d does not match Avalanche context %d", genesis.SecurityDomain.NetworkID, chainContext.NetworkID)
	}
	if subnetID != chainContext.SubnetID {
		return kernel.ChainConfig{}, errors.New("genesis subnet/L1 ID does not match Avalanche context")
	}
	if chainContext.ChainID == ids.Empty {
		return kernel.ChainConfig{}, errors.New("Avalanche chain ID is required for the security domain")
	}
	effectiveProfile, err := deriveChainProfileCommitment(
		genesis.Config.SemanticProfileCommitment,
		chainContext.NetworkID,
		chainContext.SubnetID,
		chainContext.ChainID,
		genesis.Config.KernelVersion,
	)
	if err != nil {
		return kernel.ChainConfig{}, fmt.Errorf("derive security-domain profile commitment: %w", err)
	}
	return kernel.ChainConfig{
		ProtocolVersion:   genesis.Config.ProtocolVersion,
		KernelVersion:     genesis.Config.KernelVersion,
		ProfileCommitment: effectiveProfile,
	}, nil
}

func ParseGenesis(data []byte) (*Genesis, error) {
	genesis := new(Genesis)
	if err := decodeCanonical(data, genesis); err != nil {
		return nil, err
	}
	if genesis.Schema != GenesisSchema || genesis.Version != 1 {
		return nil, errors.New("unsupported Continuity genesis")
	}
	if err := genesis.Limits.validate(); err != nil {
		return nil, err
	}
	if _, err := genesis.validateSecurityDomain(); err != nil {
		return nil, err
	}
	return genesis, nil
}

func encodeBoundedState(state kernel.ChainState, limits Limits) ([]byte, error) {
	if uint64(len(state.Relations)) > uint64(limits.MaxRelations) {
		return nil, errors.New("relation registry exceeds genesis bound")
	}
	if uint64(len(state.Closures)) > uint64(limits.MaxClosures) {
		return nil, errors.New("closure registry exceeds genesis bound")
	}
	if uint64(len(state.Successors)) > uint64(limits.MaxSuccessors) {
		return nil, errors.New("successor registry exceeds genesis bound")
	}
	data, err := canonicalJSON(state)
	if err != nil {
		return nil, err
	}
	if uint64(len(data)) > uint64(limits.MaxAcceptedStateBytes) {
		return nil, errors.New("accepted state exceeds genesis byte bound")
	}
	return data, nil
}
