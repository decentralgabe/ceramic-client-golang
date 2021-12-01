package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	ClayTestnet      = "https://ceramic-clay.3boxlabs.com"
	V0Path           = "api/v0"
	StreamsPath      = "streams"
	MultiqueriesPath = "multiqueries"
	PinsPath         = "pins"
	NodePath         = "node"

	ContentTypeHeader = "Content-Type"
	ContentTypeJSON   = "application/json"
)

type CeramicClient struct {
	Host     string
	BasePath string
	Path     string
	*http.Client
}

func NewCeramicClient(host, base string) *CeramicClient {
	return &CeramicClient{
		Host:     host,
		BasePath: base,
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (c CeramicClient) GetStreamState(streamID string) (*StreamStateResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, StreamsPath, streamID}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	state := StreamStateResponse{ResponseCode: resp.StatusCode}

	var data StreamState
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}
	state.Response = data

	return &state, nil
}

func (c CeramicClient) CreateStream(req CreateStreamRequest) (*CreateStreamResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, StreamsPath}, "/")

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set(ContentTypeHeader, ContentTypeJSON)

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	state := CreateStreamResponse{ResponseCode: resp.StatusCode}
	var data StreamState
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}
	state.Response = data
	return &state, nil
}

func (c CeramicClient) QueryStream(req QueryStreamRequest) (*QueryStreamResponse, error) {
	resp, err := c.QueryStreams(QueryStreamsRequest{Queries: []QueryStreamRequest{req}})
	if err != nil {
		return nil, err
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("no stream found for stream<%s> with paths: %s", req.StreamID, strings.Join(req.Paths, ", "))
	}
	if len(resp.Responses) > 1 {
		return nil, fmt.Errorf("multiple responses returned for stream<%s> with paths: %s", req.StreamID, strings.Join(req.Paths, ", "))
	}
	return &QueryStreamResponse{
		State:        resp.Responses[req.StreamID],
		ResponseCode: resp.ResponseCode,
	}, nil
}

func (c CeramicClient) QueryStreams(req QueryStreamsRequest) (*QueryStreamsResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, MultiqueriesPath}, "/")

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set(ContentTypeHeader, ContentTypeJSON)

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := QueryStreamsResponse{ResponseCode: resp.StatusCode}
	data := make(map[string]State)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}
	response.Responses = data
	return &response, nil
}

func (c CeramicClient) Commit(req CommitRequest) (*CommitResponse, error) {
	return nil, nil
}

func (c CeramicClient) AddToPinset(streamID string) (*AddToPinsetResponse, error) {
	return nil, nil
}

func (c CeramicClient) RemoveFromPinset(streamID string) (*RemoveFromPinsetResponse, error) {
	return nil, nil
}

func (c CeramicClient) ListStreamsInPinset() (*ListStreamsInPinsetResponse, error) {
	return nil, nil
}

func (c CeramicClient) ConfirmStreamInPinset(streamID string) (*ConfirmStreamInPinsetResponse, error) {
	return nil, nil
}

func (c CeramicClient) GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error) {
	return nil, nil
}

func (c CeramicClient) HealthCheck() (*HealthCheckResponse, error) {
	return nil, nil
}
