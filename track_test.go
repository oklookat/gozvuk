package gozvuk

import (
	"context"
	"testing"
)

func TestGetTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetTracks(ctx, _trackIds[:])
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetTracks) == 0 {
		t.Fatal("len(resp.Data.GetTracks) == 0")
	}
	if len(resp.Data.GetTracks[0].ID) == 0 {
		t.Fatal("len(resp.Data.GetTracks[0].ID) == 0")
	}
}

func TestFullTrack(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetFullTrack(ctx, _trackIds[:4], true, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetTracks) == 0 {
		t.Fatal("len(resp.Data.GetTracks) == 0")
	}
	if len(resp.Data.GetTracks[0].ID) == 0 {
		t.Fatal("len(resp.Data.GetTracks[0].ID) == 0")
	}
}
