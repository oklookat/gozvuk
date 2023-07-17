package gozvuk

import (
	"context"
	"testing"
)

func TestFollowingCount(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.FollowingCount(ctx, _authorIds[0])
	if err != nil {
		t.Fatal(err)
	}
}
