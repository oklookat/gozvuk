package gozvuk

import (
	"context"
	"testing"

	"github.com/oklookat/gozvuk/schema"
)

func TestAddRemoveItemCollection(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Add.

	if _, err := cl.AddItemToCollection(ctx, _artistIds[0], schema.CollectionItemTypeArtist); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToCollection(ctx, _trackIds[0], schema.CollectionItemTypeTrack); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToCollection(ctx, _albumIds[0], schema.CollectionItemTypeRelease); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToCollection(ctx, _podcastIds[0], schema.CollectionItemTypePodcast); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToCollection(ctx, _episodeIds[0], schema.CollectionItemTypeEpisode); err != nil {
		t.Fatal(err)
	}

	// Remove.

	if _, err := cl.RemoveItemFromCollection(ctx, _artistIds[0], schema.CollectionItemTypeArtist); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromCollection(ctx, _trackIds[0], schema.CollectionItemTypeTrack); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromCollection(ctx, _albumIds[0], schema.CollectionItemTypeRelease); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromCollection(ctx, _podcastIds[0], schema.CollectionItemTypePodcast); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromCollection(ctx, _episodeIds[0], schema.CollectionItemTypeEpisode); err != nil {
		t.Fatal(err)
	}
}

func TestAddRemoveItemHidden(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Add.

	if _, err := cl.AddItemToHidden(ctx, _artistIds[0], schema.CollectionItemTypeArtist); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToHidden(ctx, _trackIds[0], schema.CollectionItemTypeTrack); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToHidden(ctx, _albumIds[0], schema.CollectionItemTypeRelease); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToHidden(ctx, _podcastIds[0], schema.CollectionItemTypePodcast); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.AddItemToHidden(ctx, _episodeIds[0], schema.CollectionItemTypeEpisode); err != nil {
		t.Fatal(err)
	}

	// Remove.

	if _, err := cl.RemoveItemFromHidden(ctx, _artistIds[0], schema.CollectionItemTypeArtist); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromHidden(ctx, _trackIds[0], schema.CollectionItemTypeTrack); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromHidden(ctx, _albumIds[0], schema.CollectionItemTypeRelease); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromHidden(ctx, _podcastIds[0], schema.CollectionItemTypePodcast); err != nil {
		t.Fatal(err)
	}
	if _, err := cl.RemoveItemFromHidden(ctx, _episodeIds[0], schema.CollectionItemTypeEpisode); err != nil {
		t.Fatal(err)
	}
}

func TestGetAllHiddenCollection(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.GetAllHiddenCollection(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHiddenTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.GetHiddenTracks(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
