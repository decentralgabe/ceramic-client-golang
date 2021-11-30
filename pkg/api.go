package pkg

type CeramicAPI interface {
	// Streams //

	GetStreamState(streamID string) (*StreamStateResponse, error)
	CreateStream(req CreateStreamRequest) (*CreateStreamResponse, error)

	// Multiqueries //

	QueryStreams(req QueryStreamsRequest) (*QueryStreamsResponse, error)

	// Commits //
	
	Commit(req CommitRequest) (*CommitResponse, error)

	// Pins //

	AddToPinset(streamID string) (*AddToPinsetResponse, error)
	RemoveFromPinset(streamID string) (*RemoveFromPinsetResponse, error)
	ListStreamsInPinset() (*ListStreamsInPinsetResponse, error)
	ConfirmStreamInPinset(streamID string) (*ConfirmStreamInPinsetResponse, error)

	// Node Info //

	GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error)
	HealthCheck() (*HealthCheckResponse, error)
}

type JSONResponse map[string]interface{}

// Streams API //

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

// Multiqueries API //

type QueryStreamsRequest struct {
	Queries []QueryStreamRequest `json:"queries"`
}

type QueryStreamRequest struct {
	StreamID string   `json:"streamId"`
	Paths    []string `json:"paths"`
}

type QueryStreamsResponse struct {
	Responses map[string]State
}

// Commit API //

type UpdateOpts struct {
	Anchor  bool `json:"anchor,omitempty"`
	Pin     bool `json:"pin,omitempty"`
	Publish bool `json:"publish,omitempty"`
}

type CommitRequest struct {
	StreamID string      `json:"streamId"`
	Commit   interface{} `json:"commit"`
	Opts     UpdateOpts  `json:"opts,omitempty"`
}

type CommitResponse struct {
	Response     StreamState `json:"response,omitempty"`
	ResponseCode int         `json:"code"`
}

func Commit(req CommitRequest) (*CommitResponse, error) {
	return nil, nil
}

// Pins API //

type AddToPinsetResponse struct {
	StreamID     string `json:"streamId"`
	ResponseCode int    `json:"code,omitempty"`
}

func AddToPinset(streamID string) (*AddToPinsetResponse, error) {
	return nil, nil
}

type RemoveFromPinsetResponse struct {
	StreamID     string `json:"streamId"`
	ResponseCode int    `json:"code"`
}

func RemoveFromPinset(streamID string) (*RemoveFromPinsetResponse, error) {
	return nil, nil
}

type ListStreamsInPinsetResponse struct {
	PinnedStreamIDs []string `json:"pinnedStreamIds"`
	ResponseCode    int      `json:"code"`
}

func ListStreamsInPinset() (*ListStreamsInPinsetResponse, error) {
	return nil, nil
}

type ConfirmStreamInPinsetResponse struct {
	PinnedStreamIDs []string `json:"pinnedStreamIds"`
	ResponseCode    int      `json:"code"`
}

func ConfirmStreamInPinset(streamID string) (*ConfirmStreamInPinsetResponse, error) {
	return nil, nil
}

// Node Info API //

type GetSupportedBlockchainsResponse struct {
	SupportedChains []string `json:"supportedChains"`
	ResponseCode    int      `json:"code"`
}

func GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error) {
	return nil, nil
}

type HealthCheckResponse struct {
	HealthStatus string `json:"healthStatus"`
	ResponseCode int    `json:"code"`
}

func HealthCheck() (*HealthCheckResponse, error) {
	return nil, nil
}
