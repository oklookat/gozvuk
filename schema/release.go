package schema

import _ "embed"

var (
	//go:embed gql/getReleases.gql
	_getReleases string
)

func GetReleases(ids []ID, withTracks, withArtists, withLabel, withRelated bool, relatedLimit int) (string, error) {
	return getGraphqlBody(_getReleases, "getReleases", map[string]any{
		"withTracks":   withTracks,
		"withArtists":  withArtists,
		"withLabel":    withLabel,
		"withRelated":  withRelated,
		"relatedLimit": relatedLimit,
		"ids":          ids,
	})
}

type GetReleasesResponse struct {
	GetReleases []Release `json:"getReleases"`
}

type (
	// Например альбом.
	Release struct {
		ID          ID           `json:"id"`
		Title       string       `json:"title"`
		SearchTitle *string      `json:"searchTitle"`
		Type        *ReleaseType `json:"type"`
		Date        *Time        `json:"date"`
		Genres      []Genre      `json:"genres"`
		Image       *Image       `json:"image"`
		Artists     []Artist     `json:"artists"`
		Tracks      []Track      `json:"tracks"`
		Label       *Label       `json:"label"`
		// Схожие релизы (ремиксы, и так далее)?
		Related        []Release `json:"related"`
		Availability   *int      `json:"availability"`
		ArtistTemplate *string   `json:"artistTemplate"`
	}

	// Жанр.
	Genre struct {
		// ID жанра.
		ID *ID `json:"id"`

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
