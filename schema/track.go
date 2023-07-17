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

type GetTracksResponse struct {
	GetTracks []SimpleTrack `json:"getTracks"`
}

func GetFullTrack(ids []ID) (string, error) {
	return getGraphqlBody(_getFullTrack, "getFullTrack", map[string]any{
		"ids": ids,
	})
}

type GetFullTrackResponse struct {
	GetTracks []Track `json:"getTracks"`
}

type (
	// Краткая информация о треке.
	SimpleTrack struct {
		ID ID `json:"id"`

		// Название.
		Title string `json:"title"`

		// Длительность в секундах.
		Duration int `json:"duration"`

		// Трек с ненормативной лексикой?
		Explicit bool `json:"explicit"`

		// Артисты.
		Artists []SimpleArtist `json:"artists"`

		// Релиз.
		Release *SimpleRelease `json:"release"`
	}

	// Трек.
	Track struct {
		SimpleTrack

		// Альтернативное название по которому трек можно найти(?).
		SearchTitle string `json:"searchTitle"`

		Position     int `json:"position"`
		Availability int `json:"availability"`

		// Пример: {0}
		ArtistTemplate string `json:"artistTemplate"`

		Condition string `json:"condition"`

		// Текст.
		Lyrics any `json:"lyrics"`

		// Название лейбла?
		Zchan string `json:"zchan"`

		// Трек лайкнут?
		CollectionItemData CollectionItem `json:"collectionItemData"`

		// Есть FLAC-версия?
		HasFlac bool `json:"hasFlac"`

		// Имена артистов.
		ArtistNames []string `json:"artistNames"`

		// Фиты?
		Credits string `json:"credits"`
	}
)
