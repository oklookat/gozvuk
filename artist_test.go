package gozvuk

import (
	"context"
	"testing"
)

var _artistIds = [4]string{
	"292150",
	"681777",
	"1221500",
	"76007136",
}

func TestGetArtists(t *testing.T) {
	cl := getClient(t)

	ctx := context.Background()
	resp, err := cl.GetArtists(ctx, _artistIds[:1], true, true, true, true, 0, 2, 0, 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetArtists) == 0 {
		t.Fatal("empty getArtists")
	}
}
