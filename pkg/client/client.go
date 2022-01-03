package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/decentralgabe/ceramic-client-golang/pkg/api"
	"github.com/decentralgabe/ceramic-client-golang/pkg/streams"
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
	CommitsPath      = "streams"
	PinsPath         = "pins"
	NodePath         = "node"
	ChainsPath       = "chains"
	HealthcheckPath  = "healthcheck"

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

func (c CeramicClient) GetStreamState(req api.StreamStateRequest) (*api.StreamStateResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, StreamsPath, req.StreamID}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data streams.StreamState
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &api.StreamStateResponse{
		Response:     data,
		ResponseCode: resp.StatusCode,
	}, nil
}

func (c CeramicClient) CreateStream(req api.CreateStreamRequest) (*api.CreateStreamResponse, error) {
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

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data streams.StreamStateHolder
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &api.CreateStreamResponse{
		Response:     data,
		ResponseCode: resp.StatusCode,
	}, nil
}

func (c CeramicClient) QueryStream(req api.QueryStreamRequest) (*api.QueryStreamResponse, error) {
	resp, err := c.QueryStreams(api.QueryStreamsRequest{Queries: []api.QueryStreamRequest{req}})
	if err != nil {
		return nil, err
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("stream not found for stream<%s> with paths: %s", req.StreamID, strings.Join(req.Paths, ", "))
	}
	if len(resp.Responses) > 1 {
		return nil, fmt.Errorf("multiple responses returned for stream<%s> with paths: %s", req.StreamID, strings.Join(req.Paths, ", "))
	}
	return &api.QueryStreamResponse{
		Response:     resp.Responses[req.StreamID],
		ResponseCode: resp.ResponseCode,
	}, nil
}

func (c CeramicClient) QueryStreams(req api.QueryStreamsRequest) (*api.QueryStreamsResponse, error) {
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

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]streams.StreamState)
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &api.QueryStreamsResponse{
		Responses:    data,
		ResponseCode: resp.StatusCode,
	}, nil
}

func (c CeramicClient) GetCommits(req api.GetCommitsRequest) (*api.GetCommitsResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, CommitsPath, req.StreamID}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	getCommitsResp := api.GetCommitsResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &getCommitsResp); err != nil {
		return nil, err
	}

	return &getCommitsResp, nil
}

func (c CeramicClient) ApplyCommit(req api.ApplyCommitRequest) (*api.ApplyCommitResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, CommitsPath}, "/")

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

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data streams.StreamState
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &api.ApplyCommitResponse{
		Response:     data,
		ResponseCode: resp.StatusCode,
	}, nil
}

func (c CeramicClient) AddToPinset(req api.AddToPinsetRequest) (*api.AddToPinsetResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, PinsPath, req.StreamID}, "/")

	// no body, no content type https://stackoverflow.com/a/29784642
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	addToPinsetResp := api.AddToPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &addToPinsetResp); err != nil {
		return nil, err
	}

	return &addToPinsetResp, nil
}

func (c CeramicClient) RemoveFromPinset(req api.RemoveFromPinsetRequest) (*api.RemoveFromPinsetResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, PinsPath, req.StreamID}, "/")

	httpReq, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set(ContentTypeHeader, ContentTypeJSON)

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	removeFromPinsetResp := api.RemoveFromPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &removeFromPinsetResp); err != nil {
		return nil, err
	}

	return &removeFromPinsetResp, nil
}

func (c CeramicClient) ListStreamsInPinset() (*api.ListStreamsInPinsetResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, PinsPath}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	streamsInPinsetResp := api.ListStreamsInPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &streamsInPinsetResp); err != nil {
		return nil, err
	}

	return &streamsInPinsetResp, nil
}

func (c CeramicClient) ConfirmStreamInPinset(req api.ConfirmStreamInPinsetRequest) (*api.ConfirmStreamInPinsetResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, PinsPath, req.StreamID}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	streamInPinsetResp := api.ConfirmStreamInPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &streamInPinsetResp); err != nil {
		return nil, err
	}

	return &streamInPinsetResp, nil
}

func (c CeramicClient) GetSupportedBlockchains() (*api.GetSupportedBlockchainsResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, NodePath, ChainsPath}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	supportedChains := api.GetSupportedBlockchainsResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &supportedChains); err != nil {
		return nil, err
	}

	return &supportedChains, nil
}

func (c CeramicClient) HealthCheck() (*api.HealthCheckResponse, error) {
	url := strings.Join([]string{c.Host, c.BasePath, NodePath, HealthcheckPath}, "/")
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &api.HealthCheckResponse{
		HealthStatus: string(respBytes),
		ResponseCode: resp.StatusCode,
	}, nil
}
