package gozvuk

import (
	"context"
	"testing"

	"github.com/oklookat/gozvuk/schema"
)

func TestGetPlaylists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetPlaylists(ctx, _playlistIds[:], true, true, 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetPlaylists) == 0 {
		t.Fatal("empty getPlaylists")
	}
}

func TestGetPlaylistTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetPlaylistTracks(ctx, _playlistIds[0], 10, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetPlaylistTracks) == 0 {
		t.Fatal("empty getPlaylistTracks")
	}
}

func TestPlaylistCRUD(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.CreatePlaylist(ctx, []schema.PlaylistItem{
		schema.NewPlaylistItem(schema.PlaylistItemTypeTrack, _trackIds[0]),
	}, "testing")
	if err != nil {
		t.Fatal(err)
	}
	id := resp.Data.Playlist.Create
	if len(id) == 0 {
		t.Fatal("empty Playlist.Create ID")
	}
	testRenamePlaylist(ctx, cl, t, id, "testingNew")
	testSetPlaylistToPublic(ctx, cl, t, id, false)
	testSetPlaylistToPublic(ctx, cl, t, id, true)
	testDeletePlaylist(ctx, cl, t, id)
}

func testDeletePlaylist(ctx context.Context,
	cl *Client, t *testing.T, id schema.ID) {
	_, err := cl.DeletePlaylist(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
}

func testRenamePlaylist(ctx context.Context,
	cl *Client, t *testing.T, id schema.ID, newName string) {
	_, err := cl.RenamePlaylist(ctx, id, newName)
	if err != nil {
		t.Fatal(err)
	}
}

func testSetPlaylistToPublic(ctx context.Context,
	cl *Client, t *testing.T, id schema.ID, isPublic bool) {
	_, err := cl.SetPlaylistToPublic(ctx, id, isPublic)
	if err != nil {
		t.Fatal(err)
	}
}
