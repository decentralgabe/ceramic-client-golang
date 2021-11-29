package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewCeramicClient(ClayTestnet, V0)
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
	streamResp, err := client.GetStreamState(createResp.Response.StreamID)
	assert.NoError(t, err)
	assert.NotEmpty(t, streamResp)

	fmt.Printf("%+v\n", streamResp)
}
