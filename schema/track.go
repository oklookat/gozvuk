package schema

import (
	_ "embed"
)

var (
	//go:embed gql/getTracks.gql
	_getTracks string
	//go:embed gql/getFullTrack.gql
	_getFullTrack string
)

func GetTracks(ids []ID) (string, error) {
	return getGraphqlBody(_getTracks, "getTracks", map[string]any{
		"ids": ids,
	})
}

type GetFullTrackResponse struct {
	GetTracks []Track `json:"getTracks"`
}

func GetFullTrack(ids []ID, withReleases, withArtists bool) (string, error) {
	return getGraphqlBody(_getFullTrack, "getFullTrack", map[string]any{
		"withReleases": withReleases,
		"withArtists":  withArtists,
		"ids":          ids,
	})
}

type GetTracksResponse struct {
	GetTracks []Track `json:"getTracks"`
}

type (
	// Трек.
	Track struct {
		ID ID `json:"id"`

		// Название.
		Title *string `json:"title"`

		// Альтернативное название по которому трек можно найти?
		SearchTitle *string `json:"searchTitle"`

		Position *int `json:"position"`

		// Длительность в секундах.
		Duration *int `json:"duration"`

		Availability   *int    `json:"availability"`
		ArtistTemplate *string `json:"artistTemplate"`
		Condition      *string `json:"condition"`

		// Трек с ненормативной лексикой?
		Explicit *bool `json:"explicit"`

		// Текст.
		Lyrics any `json:"lyrics"`

		// Название лейбла?
		Zchan *string `json:"zchan"`

		CollectionItemData *struct {
			ItemStatus any `json:"itemStatus"`
		} `json:"collectionItemData"`

		// Артисты.
		Artists []Artist `json:"artists"`

		// Имена артистов.
		ArtistNames []string `json:"artistNames"`

		// Релиз.
		Release Release `json:"release"`

		// Фиты?
		Credits *string `json:"credits"`

		// Есть FLAC-версия?
		HasFlac *bool `json:"hasFlac"`
	}
)
