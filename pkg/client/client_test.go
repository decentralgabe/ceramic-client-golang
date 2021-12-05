package client

import (
	"fmt"
	"github.com/glcohen/ceramic-client-golang/pkg/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStreams(t *testing.T) {
	client := NewCeramicClient(ClayTestnet, V0Path)
	assert.NotEmpty(t, client)

	createReq := api.CreateStreamRequest{
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

	//{"streamId":"k2t6wyfsu4pg2qvoorchoj23e8hf3eiis4w7bucllxkmlk91sjgluuag5syphl","state":{"type":0,"content":{},"metadata":{"family":"test","controllers":["did:key:z6MkfZ6S4NVVTEuts8o5xFzRMR8eC6Y1bngoBQNnXiCvhH8H"]},"signature":0,"anchorStatus":"ANCHORED","log":[{"cid":"bafyreihtdxfb6cpcvomm2c2elm3re2onqaix6frq4nbg45eaqszh5mifre","type":0},{"cid":"bafyreic6vh3eiuuzwztyjxl4tjw2gkmb5ypco7zcdcnkkjzaicxdllt33e","type":2,"timestamp":1611680505}],"anchorProof":{"root":"bafyreiastagccuwzhjtrvmpxhx62ykswj2377tixkfxhpobze4y3iehjba","txHash":"bagjqcgzabm3nr5eme65yefacrnkaihlhdxooxa6rqfgabbanc5qacarkcxqa","chainId":"eip155:3","blockNumber":9542177,"blockTimestamp":1611680505},"doctype":"tile"}}
	// get it back
	streamResp, err := client.GetStreamState(api.StreamStateRequest{StreamID: createResp.Response.ID})
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
