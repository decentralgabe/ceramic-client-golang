package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	ClayTestnet  = "https://ceramic-clay.3boxlabs.com"
	V0           = "api/v0"
	Streams      = "streams"
	Multiqueries = "multiqueries"
	Commits      = "commits"
	Pins         = "pins"
	Node         = "node"

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
	url := strings.Join([]string{c.Host, c.BasePath, "streams", streamID}, "/")
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
	url := strings.Join([]string{c.Host, c.BasePath, "streams"}, "/")

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
