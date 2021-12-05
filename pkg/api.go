package pkg

type CeramicAPI interface {
	// StreamsPath //

	GetStreamState(rqe StreamStateRequest) (*StreamStateResponse, error)
	CreateStream(req CreateStreamRequest) (*CreateStreamResponse, error)

	// MultiqueriesPath //

	QueryStream(req QueryStreamRequest) (*QueryStreamResponse, error)
	QueryStreams(req QueryStreamsRequest) (*QueryStreamsResponse, error)

	// Commits //

	GetCommits(req GetCommitsRequest) (*GetCommitsResponse, error)
	ApplyCommit(req ApplyCommitRequest) (*ApplyCommitResponse, error)

	// Pins //

	AddToPinset(req AddToPinsetRequest) (*AddToPinsetResponse, error)
	RemoveFromPinset(req RemoveFromPinsetRequest) (*RemoveFromPinsetResponse, error)
	ListStreamsInPinset() (*ListStreamsInPinsetResponse, error)
	ConfirmStreamInPinset(req ConfirmStreamInPinsetRequest) (*ConfirmStreamInPinsetResponse, error)

	// Node Info //

	GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error)
	HealthCheck() (*HealthCheckResponse, error)
}

type JSONResponse map[string]interface{}

// StreamsPath API //

type StreamStateRequest struct {
	StreamID string `json:"streamId"`
}

type StreamStateResponse struct {
	Response     StreamState `json:"response"`
	ResponseCode int         `json:"code"`
}

type StreamState struct {
	StreamID string `json:"streamId"`
	State    State  `json:"state"`
}

type State struct {
	Type               int         `json:"type"`
	Content            interface{} `json:"content"`
	Metadata           Metadata    `json:"metadata"`
	Signature          int         `json:"signature"`
	AnchorStatus       string      `json:"anchorStatus"`
	Log                []LogEntry  `json:"log"`
	AnchorScheduledFor string      `json:"anchorScheduledFor"`
}

type Metadata struct {
	Family      string   `json:"family"`
	Controllers []string `json:"controllers"`
}

type LogEntry struct {
	CID  string `json:"cid"`
	Type int    `json:"type"`
}

type CreateOpts struct {
	Anchor             bool        `json:"anchor,omitempty"`
	Pin                bool        `json:"pin,omitempty"`
	Publish            bool        `json:"publish,omitempty"`
	Sync               SyncOptions `json:"sync,omitempty"`
	SyncTimeoutSeconds int         `json:"syncTimeoutSeconds,omitempty"`
}

type SyncOptions int

const (
	NeverSync SyncOptions = iota
	PreferCache
	SyncAlways
)

type CreateStreamRequest struct {
	// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-59/tables/streamtypes.csv
	Type    int         `json:"type"`
	Genesis interface{} `json:"genesis"`
	Opts    CreateOpts  `json:"opts,omitempty"`
}

type CreateStreamResponse struct {
	Response     StreamState `json:"response"`
	ResponseCode int         `json:"code"`
}

// MultiqueriesPath API //

type QueryStreamsRequest struct {
	Queries []QueryStreamRequest `json:"queries"`
}

type QueryStreamRequest struct {
	StreamID string   `json:"streamId"`
	Paths    []string `json:"paths"`
}

type QueryStreamsResponse struct {
	Responses    map[string]State `json:"responses"`
	ResponseCode int              `json:"code"`
}

type QueryStreamResponse struct {
	Response     State `json:"state"`
	ResponseCode int   `json:"code"`
}

// GetCommits API //

type GetCommitsRequest struct {
	StreamID string `json:"streamId"`
}

type GetCommitsResponse struct {
	StreamID     string        `json:"streamId"`
	Commits      []interface{} `json:"streams"`
	ResponseCode int           `json:"code"`
}

type UpdateOpts struct {
	Anchor  bool `json:"anchor,omitempty"`
	Pin     bool `json:"pin,omitempty"`
	Publish bool `json:"publish,omitempty"`
}

type ApplyCommitRequest struct {
	StreamID string      `json:"streamId"`
	Commit   interface{} `json:"commit"`
	Opts     UpdateOpts  `json:"opts,omitempty"`
}

type ApplyCommitResponse struct {
	Response     StreamState `json:"response"`
	ResponseCode int         `json:"code"`
}

// Pins API //

type AddToPinsetRequest struct {
	StreamID string `json:"streamId"`
}

type AddToPinsetResponse struct {
	StreamID     string `json:"streamId"`
	ResponseCode int    `json:"code"`
}

type RemoveFromPinsetRequest struct {
	StreamID string `json:"streamId"`
}

type RemoveFromPinsetResponse struct {
	StreamID     string `json:"streamId"`
	ResponseCode int    `json:"code"`
}

type ListStreamsInPinsetResponse struct {
	PinnedStreamIDs []string `json:"pinnedStreamIds"`
	ResponseCode    int      `json:"code"`
}

type ConfirmStreamInPinsetRequest struct {
	StreamID string `json:"streamId"`
}

type ConfirmStreamInPinsetResponse struct {
	PinnedStreamIDs []string `json:"pinnedStreamIds"`
	ResponseCode    int      `json:"code"`
}

// Node Info API //

type GetSupportedBlockchainsResponse struct {
	SupportedChains []string `json:"supportedChains"`
	ResponseCode    int      `json:"code"`
}

type HealthCheckResponse struct {
	HealthStatus string `json:"healthStatus"`
	ResponseCode int    `json:"code"`
}
