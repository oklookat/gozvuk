package gozvuk

import (
	"context"
	"testing"
)

func TestGetArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetArtists(ctx, _artistIds[:], true, true, true, true, 0, 10, 0, 10, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetArtists) == 0 {
		t.Fatal("empty getArtists")
	}
}
