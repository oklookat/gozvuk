package schema

type (
	SimpleBook struct {
		ID          ID       `json:"id"`
		Title       string   `json:"title"`
		AuthorNames []string `json:"authorNames"`
	}

	Book struct {
		SimpleBook

		Availability int `json:"availability"`
	}
)
