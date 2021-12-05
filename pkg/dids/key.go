package dids

import (
	"crypto/ed25519"
	"fmt"
	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

const (
	// Base58BTCMultiBase Base58BTC https://github.com/multiformats/go-multibase/blob/master/multibase.go
	Base58BTCMultiBase = multibase.Base58BTC

	// Ed25519MultiCodec ed25519-pub https://github.com/multiformats/multicodec/blob/master/table.csv
	Ed25519MultiCodec = multicodec.Ed25519Pub

	// DIDPrefix did:key prefix
	DIDPrefix = "did:key"
)

func CreateDIDKey(key ed25519.PublicKey) (*string, error) {
	// did:key:<multibase encoded, multicodec identified, public key>
	prefix := varint.ToUvarint(uint64(Ed25519MultiCodec))
	multiCodec := append(prefix, key...)
	encoded, err := multibase.Encode(Base58BTCMultiBase, multiCodec)
	if err != nil {
		return nil, err
	}
	did := fmt.Sprintf("%s:%s", DIDPrefix, encoded)
	return &did, nil
}
