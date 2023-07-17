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

func TestGetShortPlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetShortPlaylist(ctx, _playlistIds[:4], 3, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetPlaylists) == 0 {
		t.Fatal("empty getPlaylists")
	}
}

func TestSynthesisPlaylistBuild(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	var ids []schema.ID
	add := func(id1, id2 schema.ID) {
		res, err := cl.SynthesisPlaylistBuild(ctx, id1, id2)
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, res.Data.SynthesisPlaylistBuild.ID)
	}
	add(_authorIds[0], _authorIds[1])
	add(_authorIds[1], _authorIds[2])

	testSynthesisPlaylist(ctx, t, cl, ids)
}

func testSynthesisPlaylist(ctx context.Context, t *testing.T, cl *Client, ids []schema.ID) {
	resp, err := cl.SynthesisPlaylist(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.SynthesisPlaylist) == 0 {
		t.Fatal("empty resp.Data.SynthesisPlaylist")
	}
	if len(resp.Data.SynthesisPlaylist[0].ID) == 0 {
		t.Fatal("empty resp.Data.SynthesisPlaylist[0].ID")
	}
}
