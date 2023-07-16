package gozvuk

import (
	"context"
	"testing"

	"github.com/oklookat/gozvuk/schema"
)

func TestAddItemToCollection(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	if err := cl.AddItemToCollection(ctx, _artistIds[0], schema.CollectionItemTypeArtist); err != nil {
		t.Fatal(err)
	}
	if err := cl.AddItemToCollection(ctx, _trackIds[0], schema.CollectionItemTypeTrack); err != nil {
		t.Fatal(err)
	}
	if err := cl.AddItemToCollection(ctx, _albumIds[0], schema.CollectionItemTypeRelease); err != nil {
		t.Fatal(err)
	}
	if err := cl.AddItemToCollection(ctx, _podcastIds[0], schema.CollectionItemTypePodcast); err != nil {
		t.Fatal(err)
	}
	if err := cl.AddItemToCollection(ctx, _episodeIds[0], schema.CollectionItemTypeEpisode); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveItemFromCollection(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	if err := cl.RemoveItemFromCollection(ctx, _artistIds[0], schema.CollectionItemTypeArtist); err != nil {
		t.Fatal(err)
	}
	if err := cl.RemoveItemFromCollection(ctx, _trackIds[0], schema.CollectionItemTypeTrack); err != nil {
		t.Fatal(err)
	}
	if err := cl.RemoveItemFromCollection(ctx, _albumIds[0], schema.CollectionItemTypeRelease); err != nil {
		t.Fatal(err)
	}
	if err := cl.RemoveItemFromCollection(ctx, _podcastIds[0], schema.CollectionItemTypePodcast); err != nil {
		t.Fatal(err)
	}
	if err := cl.RemoveItemFromCollection(ctx, _episodeIds[0], schema.CollectionItemTypeEpisode); err != nil {
		t.Fatal(err)
	}
}
