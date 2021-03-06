package streams

import (
	"encoding/json"
	"github.com/textileio/go-did-resolver/threeid"
)

type StreamType int

const (
	// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-59/tables/streamtypes.csv

	Tile StreamType = iota
	CAIP10Link
)

type SignatureStatus int

const (
	GenesisSigStatus SignatureStatus = iota
	PartialSigStatus
	SignedSigStatus
)

type AnchorStatus string

const (
	NotRequested AnchorStatus = "NOT_REQUESTED"
	Pending                   = "PENDING"
	Processing                = "PROCESSING"
	Anchored                  = "ANCHORED"
	Failed                    = "FAILED"
)

type CommitType int

const (
	GenesisCommitType CommitType = iota
	SignedCommitType
	AnchorCommitType
)

type CommitHeader struct {
	Controllers []string         `json:"controllers,omitempty"`
	Family      string           `json:"family,omitempty"`
	Schema      string           `json:"schema,omitempty"`
	Tags        []string         `json:"tags,omitempty"`
	Index       *json.RawMessage `json:"index,omitempty"`
}

type GenesisHeader struct {
	CommitHeader           `json:"commitHeader,omitempty"`
	Unique                 string `json:"unique,omitempty"`
	ForbidControllerChange bool   `json:"forbidControllerChange,omitempty"`
}

type GenesisCommit struct {
	Header GenesisHeader    `json:"header,omitempty"`
	Data   *json.RawMessage `json:"data,omitempty"`
}

type RawCommit struct {
	ID     string           `json:"id,omitempty"`
	Header CommitHeader     `json:"header,omitempty"`
	Data   *json.RawMessage `json:"data,omitempty"`
	Prev   string           `json:"prev,omitempty"`
}

type AnchorProof struct {
	ChainID        string `json:"chainId,omitempty"`
	BlockNumber    uint64 `json:"blockNumber,omitempty"`
	BlockTimestamp uint64 `json:"blockTimestamp,omitempty"`
	TxHash         string `json:"txHash,omitempty"`
	Root           string `json:"root,omitempty"`
}

type AnchorCommit struct {
	ID    string `json:"id,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Proof string `json:"proof,omitempty"`
	Path  string `json:"path,omitempty"`
}

type StreamMetadata struct {
	Controllers            []string         `json:"controllers,omitempty"`
	Family                 string           `json:"family,omitempty"`
	Schema                 string           `json:"schema,omitempty"`
	Tags                   []string         `json:"tags,omitempty"`
	ForbidControllerChange bool             `json:"forbidControllerChange,omitempty"`
	Index                  *json.RawMessage `json:"index,omitempty"`
}

type StreamNext struct {
	Content     *json.RawMessage `json:"content,omitempty"`
	Controllers []string         `json:"controllers,omitempty"`
	Metadata    StreamMetadata   `json:"metadata"`
}

type LogEntry struct {
	CID       string     `json:"cid,omitempty"`
	Type      CommitType `json:"type,omitempty"`
	Timestamp uint64     `json:"timestamp,omitempty"`
}

type JWSSignature struct {
	Protected string `json:"protected,omitempty"`
	Signature string `json:"signature,omitempty"`
}

type DAGJWS struct {
	Payload    string         `json:"payload,omitempty"`
	Signatures []JWSSignature `json:"signatures,omitempty"`
	Link       string         `json:"link,omitempty"`
}

type CommitData struct {
	LogEntry         `json:"logEntry,omitempty"`
	Commit           *json.RawMessage `json:"commit,omitempty"`
	Envelope         DAGJWS           `json:"envelope,omitempty"`
	Proof            AnchorProof      `json:"proof,omitempty"`
	DisableTimeCheck bool             `json:"disableTimeCheck,omitempty"`
}

type StreamState struct {
	Type               uint64           `json:"type,omitempty"`
	Content            *json.RawMessage `json:"content,omitempty"`
	Next               StreamNext       `json:"next,omitempty"`
	Metadata           StreamMetadata   `json:"metadata,omitempty"`
	Signature          SignatureStatus  `json:"signature,omitempty"`
	AnchorStatus       string           `json:"anchorStatus,omitempty"`
	AnchorScheduledFor uint64           `json:"anchorScheduledFor,omitempty"`
	AnchorProof        AnchorProof      `json:"anchorProof"`
	Log                []LogEntry       `json:"log,omitempty"`
	DocType            string           `json:"doctype,omitempty"`
}

type StreamStateHolder struct {
	ID    string      `json:"streamId,omitempty"`
	State StreamState `json:"state,omitempty"`
}

type Stream interface {
	ID() string
	// API() api.CeramicAPI
	Metadata() StreamMetadata
	Content() interface{}
	Controllers() []string
	Tip() string
	CommitID() threeid.CommitID
	AllCommitIDs() []threeid.CommitID
	AnchorCommitIDs() []threeid.CommitID
	State() StreamState
	Sync()
	RequestAnchor() AnchorStatus
	MakeReadOnly()
	IsReadOnly() bool
}
