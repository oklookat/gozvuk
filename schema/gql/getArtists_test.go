package gql

import "testing"

func TestGetArtists(t *testing.T) {
	jsond, err := GetArtists([]string{"123"}, true, true, true, true, 0, 10, 10, 0, 10)
	if err != nil {
		t.Fatal(t)
	}
	println(jsond)
}
