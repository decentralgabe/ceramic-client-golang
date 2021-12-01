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
	CommitsPath      = "commits"
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

func (c CeramicClient) GetStreamState(req StreamStateRequest) (*StreamStateResponse, error) {
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

	var data StreamState
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &StreamStateResponse{
		Response:     data,
		ResponseCode: resp.StatusCode,
	}, nil
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

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data StreamState
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &CreateStreamResponse{
		Response:     data,
		ResponseCode: resp.StatusCode,
	}, nil
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
		Response:     resp.Responses[req.StreamID],
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

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]State)
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &QueryStreamsResponse{
		Responses:    data,
		ResponseCode: resp.StatusCode,
	}, nil
}

func (c CeramicClient) GetCommits(req GetCommitsRequest) (*GetCommitsResponse, error) {
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

	getCommitsResp := GetCommitsResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &getCommitsResp); err != nil {
		return nil, err
	}

	return &getCommitsResp, nil
}

func (c CeramicClient) ApplyCommit(req ApplyCommitRequest) (*ApplyCommitResponse, error) {
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

	var data StreamState
	if err := json.Unmarshal(respBytes, &data); err != nil {
		return nil, err
	}

	return &ApplyCommitResponse{
		Response:     data,
		ResponseCode: resp.StatusCode,
	}, nil
}

func (c CeramicClient) AddToPinset(req AddToPinsetRequest) (*AddToPinsetResponse, error) {
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

	addToPinsetResp := AddToPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &addToPinsetResp); err != nil {
		return nil, err
	}

	return &addToPinsetResp, nil
}

func (c CeramicClient) RemoveFromPinset(req RemoveFromPinsetRequest) (*RemoveFromPinsetResponse, error) {
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

	removeFromPinsetResp := RemoveFromPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &removeFromPinsetResp); err != nil {
		return nil, err
	}

	return &removeFromPinsetResp, nil
}

func (c CeramicClient) ListStreamsInPinset() (*ListStreamsInPinsetResponse, error) {
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

	streamsInPinsetResp := ListStreamsInPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &streamsInPinsetResp); err != nil {
		return nil, err
	}

	return &streamsInPinsetResp, nil
}

func (c CeramicClient) ConfirmStreamInPinset(req ConfirmStreamInPinsetRequest) (*ConfirmStreamInPinsetResponse, error) {
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

	streamInPinsetResp := ConfirmStreamInPinsetResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &streamInPinsetResp); err != nil {
		return nil, err
	}

	return &streamInPinsetResp, nil
}

func (c CeramicClient) GetSupportedBlockchains() (*GetSupportedBlockchainsResponse, error) {
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

	supportedChains := GetSupportedBlockchainsResponse{ResponseCode: resp.StatusCode}
	if err := json.Unmarshal(respBytes, &supportedChains); err != nil {
		return nil, err
	}

	return &supportedChains, nil
}

func (c CeramicClient) HealthCheck() (*HealthCheckResponse, error) {
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

	var healthCheckStatus string
	if err := json.Unmarshal(respBytes, &healthCheckStatus); err != nil {
		return nil, err
	}

	return &HealthCheckResponse{
		HealthStatus: healthCheckStatus,
		ResponseCode: resp.StatusCode,
	}, nil
}
