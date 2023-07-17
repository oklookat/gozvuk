package gozvuk

import (
	"context"
	"testing"
)

func TestGetArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetArtists(ctx, _artistIds[:3], true, 5, 0, true, 5, 0, true, 5, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetArtists) == 0 {
		t.Fatal("empty getArtists")
	}
}
