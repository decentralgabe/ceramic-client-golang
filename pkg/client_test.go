package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStreams(t *testing.T) {
	client := NewCeramicClient(ClayTestnet, V0Path)
	assert.NotEmpty(t, client)

	createReq := CreateStreamRequest{
		Type: 0,
		Genesis: map[string]interface{}{
			"header": map[string]interface{}{
				"family":      "test",
				"controllers": []string{"did:key:z6MkfZ6S4NVVTEuts8o5xFzRMR8eC6Y1bngoBQNnXiCvhH8H"},
			},
		},
	}

	// create stream
	createResp, err := client.CreateStream(createReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, createResp)

	fmt.Printf("%+v\n", createResp)

	// get it back
	streamResp, err := client.GetStreamState(StreamStateRequest{StreamID: createResp.Response.StreamID})
	assert.NoError(t, err)
	assert.NotEmpty(t, streamResp)

	fmt.Printf("%+v\n", streamResp)
}

func TestMultiqueries(t *testing.T) {

}

func TestCommits(t *testing.T) {

}

func TestPins(t *testing.T) {

}

func TestNodeInfo(t *testing.T) {
	client := NewCeramicClient(ClayTestnet, V0Path)
	assert.NotEmpty(t, client)

	t.Run("test get supported blockchains", func(tt *testing.T) {
		resp, err := client.GetSupportedBlockchains()
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, resp.ResponseCode)
		assert.NotEmpty(tt, resp.SupportedChains)
		assert.Contains(tt, resp.SupportedChains, "eip155:3")
	})

	t.Run("test health check", func(tt *testing.T) {
		resp, err := client.HealthCheck()
		assert.NoError(tt, err)
		assert.Equal(tt, http.StatusOK, resp.ResponseCode)
		assert.NotEmpty(tt, resp.HealthStatus)
		assert.Equal(tt, "Alive!", resp.HealthStatus)
	})
}
