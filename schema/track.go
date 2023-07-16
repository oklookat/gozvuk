package schema

type (
	Track struct {
		Entity
		Title *string `json:"title"`

		// Имя артиста?
		Credits     *string       `json:"credits"`
		SearchTitle *string       `json:"searchTitle"`
		Release     *ShortRelease `json:"release"`
		Artists     []Entity      `json:"artists"`
		ArtistNames []string      `json:"artistNames"`
		Position    *int          `json:"position"`
		// В секундах.
		Duration     *int    `json:"duration"`
		Availability *int    `json:"availability"`
		Condition    *string `json:"condition"`
		Explicit     *bool   `json:"explicit"`
		Lyrics       any     `json:"lyrics"`
		HasFlac      *bool   `json:"hasFlac"`
		// Название лейбла?
		Zchan          *string `json:"zchan"`
		ArtistTemplate *string `json:"artistTemplate"`
	}
)
