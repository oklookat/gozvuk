package schema

import "testing"

func TestGetArtists(t *testing.T) {
	jsond, err := GetArtists([]ID{"123"}, true, true, true, true, 0, 10, 10, 0, 10)
	if err != nil {
		t.Fatal(t)
	}
	println(jsond)
}
