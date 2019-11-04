package mainnet_v0_9_0

import "github.com/protolambda/zssz"

type Version [4]byte
type Epoch uint64
type Slot uint64
type Hash [32]byte
type BLSPubkey [48]byte
type BLSSignature [96]byte
type Gwei uint64
type Shard uint64
type ValidatorIndex uint64
type CommitteeIndex uint64

type Fork struct {
	PreviousVersion Version
	CurrentVersion  Version
	Epoch           Epoch // Epoch of latest fork
}

type Checkpoint struct {
	Epoch Epoch
	Root  Hash
}

type Validator struct {
	Pubkey                BLSPubkey
	WithdrawalCredentials Hash // Commitment to pubkey for withdrawals and transfers
	EffectiveBalance      Gwei // Balance at stake
	Slashed               bool
	// Status epochs
	ActivationEligibilityEpoch Epoch // When criteria for activation were met
	ActivationEpoch            Epoch
	ExitEpoch                  Epoch
	WithdrawableEpoch          Epoch // When validator can withdraw or transfer funds
}

type AttestationData struct {
	slot  Slot
	index CommitteeIndex
	// LMD GHOST vote
	BeaconBlockRoot Hash
	// FFG vote
	Source Checkpoint
	Target Checkpoint
}

type AttestationDataAndCustodyBit struct {
	Data       AttestationData
	CustodyBit bool // Challengeable bit (SSZ-bool, 1 byte) for the custody of shard data
}

type CommitteeIndices []ValidatorIndex

func (*CommitteeIndices) Limit() uint64 {
	return MAX_VALIDATORS_PER_COMMITTEE
}

type IndexedAttestation struct {
	CustodyBit_0Indices CommitteeIndices // Indices with custody bit equal to 0
	CustodyBit_1Indices CommitteeIndices // Indices with custody bit equal to 1
	Data                AttestationData
	Signature           BLSSignature
}

type CommitteeBits []byte

func (*CommitteeBits) BitLimit() uint64 {
	return MAX_VALIDATORS_PER_COMMITTEE
}

type PendingAttestation struct {
	AggregationBits CommitteeBits
	Data            AttestationData
	InclusionDelay  Slot
	ProposerIndex   ValidatorIndex
}

type Eth1Data struct {
	DepositRoot  Hash
	DepositCount uint64
	BlockHash    Hash
}

type SlotRoots []Hash

func (*SlotRoots) Limit() uint64 {
	return SLOTS_PER_HISTORICAL_ROOT
}

type HistoricalBatch struct {
	BlockRoots SlotRoots
	StateRoots SlotRoots
}

type DepositData struct {
	Pubkey                BLSPubkey
	WithdrawalCredentials Hash
	Amount                Gwei
	Signature             BLSSignature
}

type BeaconBlockHeader struct {
	Slot       Slot
	ParentRoot Hash
	StateRoot  Hash
	BodyRoot   Hash
	Signature  BLSSignature
}

type ProposerSlashing struct {
	ProposerIndex ValidatorIndex
	Header_1      BeaconBlockHeader
	Header_2      BeaconBlockHeader
}

type AttesterSlashing struct {
	Attestation_1 IndexedAttestation
	Attestation_2 IndexedAttestation
}

type Attestation struct {
	AggregationBits CommitteeBits
	Data            AttestationData
	CustodyBits     CommitteeBits
	Signature       BLSSignature
}

type Deposit struct {
	Proof [DEPOSIT_CONTRACT_TREE_DEPTH + 1]Hash // Merkle path to deposit data list root
	Data  DepositData
}

type VoluntaryExit struct {
	Epoch          Epoch // Earliest epoch when voluntary exit can be processed
	ValidatorIndex ValidatorIndex
	Signature      BLSSignature
}

type ProposerSlashings []ProposerSlashing

func (*ProposerSlashings) Limit() uint64 {
	return MAX_PROPOSER_SLASHINGS
}

type AttesterSlashings []AttesterSlashing

func (*AttesterSlashings) Limit() uint64 {
	return MAX_ATTESTER_SLASHINGS
}

type Attestations []Attestation

func (*Attestations) Limit() uint64 {
	return MAX_ATTESTATIONS
}

type Deposits []Deposit

func (*Deposits) Limit() uint64 {
	return MAX_DEPOSITS
}

type VoluntaryExits []VoluntaryExit

func (*VoluntaryExits) Limit() uint64 {
	return MAX_VOLUNTARY_EXITS
}

type BeaconBlockBody struct {
	RandaoReveal BLSSignature
	Eth1Data     Eth1Data // Eth1 data vote
	Graffiti     [32]byte // Arbitrary data
	// Operations
	ProposerSlashings ProposerSlashings
	AttesterSlashings AttesterSlashings
	Attestations      Attestations
	Deposits          Deposits
	VoluntaryExits    VoluntaryExits
}

type BeaconBlock struct {
	Slot       Slot
	ParentRoot Hash
	StateRoot  Hash
	Body       BeaconBlockBody
	Signature  BLSSignature
}

type HistoricalRoots []Hash

func (*HistoricalRoots) Limit() uint64 {
	return HISTORICAL_ROOTS_LIMIT
}

type Eth1Votes []Eth1Data

func (*Eth1Votes) Limit() uint64 {
	return SLOTS_PER_ETH1_VOTING_PERIOD
}

type Validators []Validator

func (*Validators) Limit() uint64 {
	return VALIDATOR_REGISTRY_LIMIT
}

type ValidatorBalances []Gwei

func (*ValidatorBalances) Limit() uint64 {
	return VALIDATOR_REGISTRY_LIMIT
}

type PendingAttestations []PendingAttestation

func (*PendingAttestations) Limit() uint64 {
	return MAX_ATTESTATIONS * SLOTS_PER_EPOCH
}

type JustificationBits [1]byte

func (jb *JustificationBits) BitLen() uint64 {
	return JUSTIFICATION_BITS_LENGTH
}

type BeaconState struct {
	// Versioning
	GenesisTime uint64
	Slot        Slot
	Fork        Fork
	// History
	LatestBlockHeader BeaconBlockHeader
	BlockRoots        SlotRoots
	StateRoots        SlotRoots
	HistoricalRoots   HistoricalRoots
	// Eth1
	Eth1Data         Eth1Data
	Eth1DataVotes    Eth1Votes
	Eth1DepositIndex uint64
	// Registry
	Validators Validators
	Balances   ValidatorBalances
	// Randomness
	RandaoMixes [EPOCHS_PER_HISTORICAL_VECTOR]Hash
	// Slashings
	Slashings [EPOCHS_PER_SLASHINGS_VECTOR]Gwei // Per-epoch sums of slashed effective balances
	// Attestations
	PreviousEpochAttestations PendingAttestations
	CurrentEpochAttestations  PendingAttestations
	// Finality
	JustificationBits           JustificationBits // Bit set for every recent justified epoch
	PreviousJustifiedCheckpoint Checkpoint        // Previous epoch snapshot
	CurrentJustifiedCheckpoint  Checkpoint
	FinalizedCheckpoint         Checkpoint
}

var BeaconBlockSSZ = zssz.GetSSZ((*BeaconBlock)(nil))
var BeaconStateSSZ = zssz.GetSSZ((*BeaconState)(nil))
