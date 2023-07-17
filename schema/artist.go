package schema

import _ "embed"

var (
	//go:embed gql/getArtists.gql
	_getArtists string
	//go:embed gql/artistFollowersCount.gql
	_artistFollowersCount string
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

func ArtistFollowersCount(ids []ID) (string, error) {
	return getGraphqlBody(_artistFollowersCount, "artistFollowersCount", map[string]any{
		"ids": ids,
	})
}

type ArtistFollowersCountResponse struct {
	GetArtists []struct {
		CollectionItemData struct {
			LikesCount int `json:"likesCount"`
		} `json:"collectionItemData"`
	} `json:"getArtists"`
}

type (
	GetArtistsResponse struct {
		GetArtists []Artist `json:"getArtists"`
	}

	Artist struct {
		ID ID `json:"id"`

		Title *string `json:"title"`

		SearchTitle *string  `json:"searchTitle"`
		Description *string  `json:"description"`
		HasPage     *bool    `json:"hasPage"`
		Profile     *Profile `json:"profile"`

		// Фото артиста.
		Image *Image `json:"image"`

		// Дополнительное фото артиста.
		SecondImage *Image `json:"secondImage"`

		Animation              *Animation `json:"animation"`
		CollectionLastModified *Time      `json:"collectionLastModified"`
		Releases               []Release  `json:"releases"`
		PopularTracks          []Track    `json:"popularTracks"`
		RelatedArtists         []Artist   `json:"relatedArtists"`
	}
)
