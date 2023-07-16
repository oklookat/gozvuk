package schema

import _ "embed"

var (
	//go:embed gql/getArtists.gql
	_getArtists string
)

func GetArtists(ids []ID,
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

type (
	GetArtistsResponse struct {
		GetArtists []Artist `json:"getArtists"`
	}

	Artist struct {
		Entity
		Title *string `json:"title"`

		SearchTitle *string  `json:"searchTitle"`
		Description *string  `json:"description"`
		HasPage     *bool    `json:"hasPage"`
		Profile     *Profile `json:"profile"`
		// Фото артиста.
		Image *Image `json:"image"`
		// Дополнительное фото артиста.
		SecondImage            *Image     `json:"secondImage"`
		Animation              *Animation `json:"animation"`
		CollectionLastModified any        `json:"collectionLastModified"`
		Releases               []Release  `json:"releases"`
		PopularTracks          []Track    `json:"popularTracks"`
		RelatedArtists         []Entity   `json:"relatedArtists"`
	}
)
