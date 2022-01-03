package streams

import (
	"github.com/decentralgabe/ceramic-client-golang/pkg/models"
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
	DefaultCreateOpts = models.CreateOpts{
		UpdateOpts: &models.UpdateOpts{
			PublishOpts: &models.PublishOpts{Publish: true},
			AnchorOpts:  &models.AnchorOpts{Anchor: true},
		},
		SyncOpts: &models.SyncOpts{
			Sync: models.PreferCache,
		},
	}

	DefaultLoadOpts = models.LoadOpts{
		SyncOpts: &models.SyncOpts{
			Sync: models.PreferCache,
		},
	}

	DefaultUpdateOpts = models.UpdateOpts{
		PublishOpts:  &models.PublishOpts{Publish: true},
		AnchorOpts:   &models.AnchorOpts{Anchor: true},
		InternalOpts: &models.InternalOpts{ThrowOnInvalidCommit: true},
	}
)
