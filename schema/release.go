package schema

import _ "embed"

var (
	//go:embed gql/getReleases.gql
	_getReleases string
)

func GetReleases(ids []ID, relatedLimit int) (string, error) {
	return getGraphqlBody(_getReleases, "getReleases", map[string]any{
		"relatedLimit": relatedLimit,
		"ids":          ids,
	})
}

type GetReleasesResponse struct {
	GetReleases []Release `json:"getReleases"`
}

type (
	// Краткая информация о релизе.
	SimpleRelease struct {
		ID ID `json:"id"`

		// Название.
		Title string `json:"title"`

		// Дата релиза.
		Date Time `json:"date"`

		// Тип релиза.
		Type ReleaseType `json:"type"`

		// Обложка.
		Image Image `json:"image"`

		// Релиз с ненормативной лексикой?
		Explicit bool `json:"explicit"`

		// Исполнители.
		Artists []SimpleArtist `json:"artists"`
	}

	// Например альбом.
	Release struct {
		SimpleRelease

		// Релиз лайкнут?
		CollectionItemData CollectionItem `json:"collectionItemData"`
		SearchTitle        string         `json:"searchTitle"`
		Genres             []Genre        `json:"genres"`

		// Треки (без поля Release).
		Tracks  []SimpleTrack  `json:"tracks"`
		Artists []SimpleArtist `json:"artists"`

		// Схожие релизы. Ремиксы, делюкс-версии, и так далее.
		Related []SimpleRelease `json:"related"`

		// Лейбл.
		Label          Label  `json:"label"`
		Availability   int    `json:"availability"`
		ArtistTemplate string `json:"artistTemplate"`
	}

	// Жанр.
	Genre struct {
		// ID жанра.
		ID ID `json:"id"`

		// Название жанра.
		Name string `json:"name"`

		// Краткое название жанра.
		ShortName *string `json:"shortName"`
	}
)

// Тип релиза.
type ReleaseType string

const (
	// Альбом.
	ReleaseTypeAlbum ReleaseType = "album"

	// Сингл.
	ReleaseTypeSingle ReleaseType = "single"
)
