package gozvuk

import (
	"context"
	"testing"
)

func TestUserCollection(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.UserCollection(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserPlaylsits(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.UserPlaylists(ctx)
	if err != nil {
		t.Fatal(err)
	}
	println(resp)
}
