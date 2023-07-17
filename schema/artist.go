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

type GetArtistsResponse struct {
	GetArtists []Artist `json:"getArtists"`
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
	SimpleArtist struct {
		// ID артиста.
		ID ID `json:"id"`

		// Имя артиста.
		Title string `json:"title"`

		// Фото артиста.
		Image *Image `json:"image"`
	}

	Artist struct {
		SimpleArtist

		// Дополнительное фото артиста.
		SecondImage *Image `json:"secondImage"`

		// Артист лайкнут?
		CollectionItemData CollectionItem `json:"collectionItemData"`

		// Имя артиста по которому его можно найти(?).
		SearchTitle string `json:"searchTitle"`

		// Об артисте.
		Description *string `json:"description"`

		Releases       []SimpleRelease `json:"releases"`
		PopularTracks  []SimpleTrack   `json:"popularTracks"`
		RelatedArtists []SimpleArtist  `json:"relatedArtists"`
	}

	// Лайкнутый артист.
	CollectionArtist struct {
		SimpleArtist

		// Когда был добавлен(?).
		CollectionLastModified *Time `json:"collectionLastModified"`
	}
)
