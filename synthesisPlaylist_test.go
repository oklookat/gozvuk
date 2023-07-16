package gozvuk

import (
	"context"
	"testing"

	"github.com/oklookat/gozvuk/schema"
)

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
