package dids

import (
	"github.com/glcohen/ceramic-client-golang/internal"
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

	t.Run("bad did method", func(tt *testing.T) {
		resolvedDID, err := resolver.Resolve("did:test:abcd")
		assert.Error(tt, err)
		assert.Empty(tt, resolvedDID)
		assert.Contains(tt, err.Error(), "unknown did method: 'test'")
	})

	t.Run("unknown did", func(tt *testing.T) {
		resolvedDID, err := resolver.Resolve("did:3:bad")
		assert.Error(tt, err)
		assert.Empty(tt, resolvedDID)
		assert.Contains(tt, err.Error(), "invalid docid")
	})

	t.Run("bad did:key", func(tt *testing.T) {
		resolvedDID, err := resolver.Resolve("did:key:bad")
		assert.Error(tt, err)
		assert.Empty(tt, resolvedDID)
		assert.Contains(tt, err.Error(), "error parsing varint")
	})

	t.Run("known did:key", func(tt *testing.T) {
		did := "did:key:z6MktvqCyLxTsXUH1tUZncNdVeEZ7hNh7npPRbUU27GTrYb8"
		resolvedDID, err := resolver.Resolve(did)
		assert.NoError(tt, err)
		assert.NotEmpty(tt, resolvedDID)

		assert.Empty(tt, resolvedDID.ResolutionMetadata)
		assert.NotEmpty(tt, resolvedDID.Document)
		assert.Empty(tt, resolvedDID.DocumentMetadata)
		assert.Equal(tt, did, resolvedDID.Document.ID)

		assert.True(tt, len(resolvedDID.Document.Authentication) == 1)
		assert.Equal(tt, "Ed25519VerificationKey2018", resolvedDID.Document.Authentication[0].Type)
		assert.Equal(tt, "zFUaAP6i2XyyouPds73QneYgZJ86qhua2jaZYBqJSwKok", resolvedDID.Document.Authentication[0].PublicKeyMultibase)
	})

	t.Run("new did:key", func(tt *testing.T) {
		pk, _, err := internal.GenerateEd25519Key()
		assert.NoError(tt, err)
		did, err := CreateDIDKey(pk)
		assert.NoError(tt, err)

		resolvedDID, err := resolver.Resolve(*did)
		assert.NoError(tt, err)
		assert.NotEmpty(tt, resolvedDID)

		assert.Empty(tt, resolvedDID.ResolutionMetadata)
		assert.NotEmpty(tt, resolvedDID.Document)
		assert.Empty(tt, resolvedDID.DocumentMetadata)
		assert.Equal(tt, *did, resolvedDID.Document.ID)

		assert.True(tt, len(resolvedDID.Document.Authentication) == 1)
		assert.Equal(tt, "Ed25519VerificationKey2018", resolvedDID.Document.Authentication[0].Type)
	})
}
