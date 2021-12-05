package streams

import (
	"github.com/glcohen/ceramic-client-golang/pkg/api"
)

type TileMetadataArgs struct {
	Controllers            []string
	Family                 string
	Tags                   []string
	Schema                 string
	Deterministic          bool
	ForbidControllerChange bool
}

var (
	DefaultCreateOpts = api.CreateOpts{
		UpdateOpts: &api.UpdateOpts{
			PublishOpts: &api.PublishOpts{Publish: true},
			AnchorOpts:  &api.AnchorOpts{Anchor: true},
		},
		SyncOpts: &api.SyncOpts{
			Sync: api.PreferCache,
		},
	}

	DefaultLoadOpts = api.LoadOpts{
		SyncOpts: &api.SyncOpts{
			Sync: api.PreferCache,
		},
	}

	DefaultUpdateOpts = api.UpdateOpts{
		PublishOpts:  &api.PublishOpts{Publish: true},
		AnchorOpts:   &api.AnchorOpts{Anchor: true},
		InternalOpts: &api.InternalOpts{ThrowOnInvalidCommit: true},
	}
)
