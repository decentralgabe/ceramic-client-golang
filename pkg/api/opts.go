package api

import "github.com/ockam-network/did"

type SyncOptions int

const (
	NeverSync SyncOptions = iota
	PreferCache
	SyncAlways
)

type PublishOpts struct {
	Publish bool `json:"publish,omitempty"`
}

type AnchorOpts struct {
	Anchor bool `json:"anchor,omitempty"`
}

type InternalOpts struct {
	ThrowOnInvalidCommit bool `json:"throwOnInvalidCommit"`
}

type PinningOpts struct {
	Pin bool `json:"pin,omitempty"`
}

type CreateOpts struct {
	*UpdateOpts
	*PinningOpts
	*SyncOpts
}

type SyncOpts struct {
	Sync               SyncOptions `json:"sync,omitempty"`
	SyncTimeoutSeconds uint64      `json:"syncTimeoutSeconds,omitempty"`
}

type UpdateOpts struct {
	*PublishOpts
	*AnchorOpts
	*InternalOpts
	*PinningOpts
	AsDID did.DID `json:"asDID,omitempty"`
}

type LoadOpts struct {
	*SyncOpts
	*PinningOpts
	AtTime uint64 `json:"atTime,omitempty"`
}
