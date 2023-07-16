package gql

import (
	_ "embed"

	"github.com/oklookat/gozvuk/schema"
)

var (
	//go:embed internal/getArtists.gql
	_getArtists string
)

type (
	GetArtistsResponse struct {
		GetArtists []schema.Artist `json:"getArtists"`
	}
)

func GetArtists(ids []string,
	withReleases, withPopTracks, withRelatedArtists, withDescription bool,
	releasesOffset, releasesLimit, tracksOffset, tracksLimit, relatedArtistsLimit int) (string, error) {
	return getGraphqlBody(_getArtists, "getArtists", map[string]any{
		"withReleases":         withReleases,
		"withPopTracks":        withPopTracks,
		"withRelatedArtists":   withRelatedArtists,
		"withDescription":      withDescription,
		"releasesOffset":       releasesOffset,
		"releasesLimit":        releasesLimit,
		"tracksOffset":         tracksOffset,
		"tracksLimit":          tracksLimit,
		"releatedArtistsLimit": relatedArtistsLimit,
		"ids":                  ids,
	})
}
