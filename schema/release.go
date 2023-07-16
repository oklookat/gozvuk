package schema

type (
	ShortRelease struct {
		ID    ID     `json:"id"`
		Title string `json:"title"`
		Image *Image `json:"image"`
	}

	Release struct {
		ShortRelease
		SearchTitle    *string             `json:"searchTitle"`
		Type           *CollectionItemType `json:"type"`
		Date           *Time               `json:"date"`
		Availability   *int                `json:"availability"`
		ArtistTemplate *string             `json:"artistTemplate"`
		Artists        []Entity            `json:"artists"`
		Genres         []Genre             `json:"genres"`
	}

	Genre struct {
		Name string `json:"name"`
	}
)
