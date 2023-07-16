package schema

type (
	ShortArtist struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	Artist struct {
		ShortArtist
		SearchTitle *string `json:"searchTitle"`
		Description *string `json:"description"`
		HasPage     *bool   `json:"hasPage"`
		// Фото артиста.
		Image *Image `json:"image"`
		// Дополнительное фото артиста.
		SecondImage            *Image        `json:"secondImage"`
		Animation              any           `json:"animation"`
		CollectionLastModified any           `json:"collectionLastModified"`
		Releases               []Release     `json:"releases"`
		PopularTracks          []Track       `json:"popularTracks"`
		RelatedArtists         []ShortArtist `json:"relatedArtists"`
	}
)
