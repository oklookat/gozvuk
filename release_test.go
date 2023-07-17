package gozvuk

import (
	"context"
	"testing"
)

func TestGetReleases(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.GetReleases(ctx, _albumIds[:4], 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data.GetReleases) == 0 {
		t.Fatal("len(resp.Data.GetReleases) == 0")
	}
	if len(resp.Data.GetReleases[0].ID) == 0 {
		t.Fatal("len(resp.Data.GetReleases[0].ID) == 0")
	}
	// release type enum check
	// for _, r := range resp.Data.GetReleases {
	// 	for _, r2 := range r.Related {
	// 		if r2.Type == nil {
	// 			continue
	// 		}
	// 		println(*r2.Type)
	// 	}
	// }
}
