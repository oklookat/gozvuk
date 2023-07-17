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
	//println(resp)
}

func TestHiddenTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.GetHiddenTracks(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserPaginatedPodcasts(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Add podcast.
	_, err := cl.AddItemToCollection(ctx, _podcastIds[0], schema.CollectionItemTypePodcast)
	if err != nil {
		t.Fatal(err)
	}

	// Get & check.
	resp, err := cl.UserPaginatedPodcasts(ctx, "", 5)
	if err != nil {
		t.Fatal(err)
	}
	podcasts := resp.Data.PaginatedCollection.Podcasts.Items
	if len(podcasts) == 0 {
		t.Fatal("len(resp.Data.PaginatedCollection.Podcasts) == 0 ")
	}
	if len(podcasts[0].ID) == 0 {
		t.Fatal("len(podcasts [0].ID) == 0")
	}

	// Remove podcast.
	_, err = cl.RemoveItemFromCollection(ctx, _podcastIds[0], schema.CollectionItemTypePodcast)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProfileFollowersCount(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.ProfileFollowersCount(ctx, _authorIds[0:2])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.Profiles) == 0 {
		t.Fatal("len(resp.Data.Profiles) == 0")
	}
	if resp.Data.Profiles[0].CollectionItemData.LikesCount < 1 {
		t.Fatal("resp.Data.Profiles[0].CollectionItemData.LikesCount < 1")
	}
}

func TestUserCollection(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.UserCollection(ctx)
	if err != nil {
		t.Fatal(err)
	}
	//println(resp)
}

func TestUserPlaylsits(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.UserPlaylists(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.UserTracks(ctx, schema.OrderByAlphabet, schema.OrderDirectionAsc)
	if err != nil {
		t.Fatal(err)
	}
	//println(resp)
}
