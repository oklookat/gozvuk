package schema

type (
	Book struct {
		ID           ID       `json:"id"`
		Availability int      `json:"availability"`
		Title        string   `json:"title"`
		AuthorNames  []string `json:"authorNames"`
	}
)
