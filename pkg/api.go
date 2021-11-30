package pkg

type CeramicAPI interface {
	// Streams

	GetStreamState(streamID string) (*StreamStateResponse, error)
	CreateStream(req CreateStreamRequest) (*CreateStreamResponse, error)

	// Multiqueries

	QueryStreams(req QueryStreamsRequest) (*QueryStreamsResponse, error)

	// Commits

	Commit(req CommitRequest) (*CommitResponse, error)

	// Pins

	AddToPinset(streamID string) (*AddToPinsetResponse, error)
	RemoveFromPinset(streamID string) (*RemoveFromPinsetResponse, error)
	ListStreamsInPinset() (*ListStreamsInPinsetResponse, error)
	ConfirmStreamInPinset(streamID string) (*ConfirmStreamInPinsetResponse, error)

	// Node Info

	GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error)
	HealthCheck() (*HealthCheckResponse, error)
}

type JSONResponse map[string]interface{}

// Streams API //

type StreamStateResponse struct {
	Response     StreamState `json:"response,omitempty"`
	ResponseCode int         `json:"code,omitempty"`
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
	Genesis interface{} `json:"genesis,omitempty"`
	Opts    CreateOpts  `json:"opts,omitempty"`
}

type CreateStreamResponse struct {
	Response     StreamState `json:"response,omitempty"`
	ResponseCode int         `json:"code,omitempty"`
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

type CommitRequest struct {
}

type CommitResponse struct {
}

func Commit(req CommitRequest) (*CommitResponse, error) {
	return nil, nil
}

// Pins API //

type AddToPinsetResponse struct {
}

func AddToPinset(streamID string) (*AddToPinsetResponse, error) {
	return nil, nil
}

type RemoveFromPinsetResponse struct {
	
}

func RemoveFromPinset(streamID string) (*RemoveFromPinsetResponse, error) {
	return nil, nil
}

type ListStreamsInPinsetResponse struct {
	
}

func ListStreamsInPinset() (*ListStreamsInPinsetResponse, error) {
	return nil, nil
}

type ConfirmStreamInPinsetResponse struct {
	
}

func ConfirmStreamInPinset(streamID string) (*ConfirmStreamInPinsetResponse, error) {
	return nil, nil
}

// Node Info API //

type GetSupportedBlockchainsResponse struct {
	
}

func GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error) {
	return nil, nil
}

type HealthCheckResponse struct {
	
}

func HealthCheck() (*HealthCheckResponse, error) {
	return nil, nil
}
