package schema

type (
	// Например автор плейлиста, подкаста, etc.
	Author struct {
		ID    ID     `json:"id"`
		Name  string `json:"name"`
		Image *Image `json:"image"`
		// Совпадение по вкусам.
		Matches *struct {
			// Коэффициент.
			Score float64 `json:"score"`
		} `json:"matches"`
	}
)
