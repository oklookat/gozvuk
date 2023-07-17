package gozvuk

import (
	"context"
	"testing"

	"github.com/oklookat/gozvuk/schema"
)

func TestSearch(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	args := schema.SearchArguments{
		Query:      "crystal",
		Limit:      5,
		Tracks:     true,
		Artists:    true,
		Releases:   true,
		Playlists:  true,
		Profiles:   true,
		Books:      true,
		Episodes:   true,
		Podcasts:   true,
		Categories: true,
	}

	resp, err := cl.Search(ctx, args)
	if err != nil {
		t.Fatal(err)
	}

	artists := resp.Data.Search.Artists
	if artists == nil {
		t.Fatal("artists cannot be nil")
	}
	if len(artists.Items) == 0 {
		t.Fatal("empty artists")
	}
}

func TestQuickSearch(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.QuickSearch(ctx, "after hours", 5, "")
	if err != nil {
		t.Fatal(err)
	}
	// for _, sr := range resp.Data.QuickSearch.Releases {
	// 	urld := sr.Image.SrcURL(100, 100)
	// 	if urld != nil {
	// 		println(urld.String())
	// 	}
	// }
	// println(resp)
}
