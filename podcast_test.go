package gozvuk

import (
	"context"
	"testing"
)

func TestGetPodcasts(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetPodcasts(ctx, _podcastIds[:4])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetPodcasts) == 0 {
		t.Fatal("empty getPlaylists")
	}
	if len(resp.Data.GetPodcasts[0].ID) == 0 {
		t.Fatal("len(resp.Data.GetPodcasts[0].ID) == 0")
	}
}

func TestGetEpisodes(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetEpisodes(ctx, _episodeIds[:4])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetEpisodes) == 0 {
		t.Fatal("len(resp.Data.GetEpisodes) == 0")
	}
	if len(resp.Data.GetEpisodes[0].ID) == 0 {
		t.Fatal("len(resp.Data.GetEpisodes[0].ID) == 0")
	}
}
