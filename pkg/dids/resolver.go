package dids

import (
	"github.com/textileio/go-did-resolver/keys"
	"github.com/textileio/go-did-resolver/resolver"
	"github.com/textileio/go-did-resolver/threeid"
)

type ResolvedDID struct {
	resolver.ResolutionMetadata
	resolver.Document
	resolver.DocumentMetadata
}

type Resolver struct {
	client   threeid.HTTPClient
	registry resolver.Registry
}

// CreateDefaultResolver The default resolver contains both a did:key and did:3 resolver
func CreateDefaultResolver(baseURL string) Resolver {
	return CreateDIDResolver(baseURL, keys.New(), threeid.New())
}

func CreateDIDResolver(baseURL string, resolvers ...resolver.Resolver) Resolver {
	client := threeid.HTTPClient{APIURL: baseURL}
	registry := resolver.New(resolvers, true)
	return Resolver{
		client:   client,
		registry: registry,
	}
}

func (r Resolver) Resolve(did string) (*ResolvedDID, error) {
	resolvedMetadata, document, documentMetadata, err := r.registry.Resolve(did, nil)
	if err != nil {
		return nil, err
	}
	return &ResolvedDID{
		ResolutionMetadata: resolvedMetadata,
		Document:           *document,
		DocumentMetadata:   documentMetadata,
	}, nil
}
