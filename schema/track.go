package schema

type (
	Track struct {
		ID          string        `json:"id"`
		Title       string        `json:"title"`
		SearchTitle *string       `json:"searchTitle"`
		Release     *Release      `json:"release"`
		Artists     []ShortArtist `json:"artists"`
		Position    *int          `json:"position"`
		// В секундах.
		Duration       *int        `json:"duration"`
		Availability   *int        `json:"availability"`
		Condition      *string     `json:"condition"`
		Explicit       *bool       `json:"explicit"`
		Lyrics         interface{} `json:"lyrics"`
		HasFlac        *bool       `json:"hasFlac"`
		Zchan          *string     `json:"zchan"`
		ArtistTemplate *string     `json:"artistTemplate"`
	}
)
