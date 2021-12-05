package api

import "github.com/glcohen/ceramic-client-golang/pkg/streams"

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
	Response     streams.StreamState `json:"response"`
	ResponseCode int                 `json:"code"`
}

type Metadata struct {
	Family      string   `json:"family"`
	Controllers []string `json:"controllers"`
}

type CreateStreamRequest struct {
	// https://github.com/ceramicnetwork/CIP/blob/main/CIPs/CIP-59/tables/streamtypes.csv
	Type    int         `json:"type"`
	Genesis interface{} `json:"genesis"`
	Opts    CreateOpts  `json:"opts,omitempty"`
}

type CreateStreamResponse struct {
	Response     streams.StreamState `json:"response"`
	ResponseCode int                 `json:"code"`
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
	Responses    map[string]streams.StreamState `json:"responses"`
	ResponseCode int                            `json:"code"`
}

type QueryStreamResponse struct {
	Response     streams.StreamState `json:"state"`
	ResponseCode int                 `json:"code"`
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

type ApplyCommitRequest struct {
	StreamID string      `json:"streamId"`
	Commit   interface{} `json:"commit"`
	Opts     UpdateOpts  `json:"opts,omitempty"`
}

type ApplyCommitResponse struct {
	Response     streams.StreamState `json:"response"`
	ResponseCode int                 `json:"code"`
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
