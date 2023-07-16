package schema

type (
	Episode struct {
		ID              ID      `json:"id"`
		Title           string  `json:"title"`
		Availability    int     `json:"availability"`
		Explicit        bool    `json:"explicit"`
		Duration        int     `json:"duration"`
		PublicationDate Time    `json:"publicationDate"`
		Image           Image   `json:"image"`
		Podcast         Podcast `json:"podcast"`
	}

	Podcast struct {
		ID      ID       `json:"id"`
		Authors []Author `json:"authors"`
		Image   Image    `json:"image"`
	}

	Author struct {
		ID   ID     `json:"id"`
		Name string `json:"name"`
	}
)
