package dids

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	ClayTestnet = "https://ceramic-clay.3boxlabs.com"
	V0Path      = "api/v0"
)

func TestResolver(t *testing.T) {
	resolver := CreateDefaultResolver(strings.Join([]string{ClayTestnet, V0Path}, "/"))
	assert.NotEmpty(t, resolver)
	resolvedDID, err := resolver.Resolve("did:test")
	assert.NoError(t, err)
	assert.NotEmpty(t, resolvedDID)
	fmt.Printf("%+v", resolvedDID)
}
