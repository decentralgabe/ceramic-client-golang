package pkg

type CeramicAPI interface {
	GetStreamState(streamID string) (*StreamStateResponse, error)
	CreateStream(req CreateStreamRequest) (*CreateStreamResponse, error)
}

type JSONResponse map[string]interface{}

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
