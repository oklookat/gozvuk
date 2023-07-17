package schema

import (
	_ "embed"
)

var (
	//go:embed gql/getPodcasts.gql
	_getPodcasts string
	//go:embed gql/getEpisodes.gql
	_getEpisodes string
)

func GetPodcasts(ids []ID) (string, error) {
	return getGraphqlBody(_getPodcasts, "getPodcasts", map[string]any{
		"ids": ids,
	})
}

type GetPodcastsResponse struct {
	GetPodcasts []Podcast `json:"getPodcasts"`
}

func GetEpisodes(ids []ID) (string, error) {
	return getGraphqlBody(_getEpisodes, "getEpisodes", map[string]any{
		"ids": ids,
	})
}

type GetEpisodesResponse struct {
	GetEpisodes []Episode `json:"getEpisodes"`
}

type (
	// Краткая информация о подкасте.
	SimplePodcast struct {
		ID       ID              `json:"id"`
		Title    string          `json:"title"`
		Explicit bool            `json:"explicit"`
		Image    Image           `json:"image"`
		Authors  []PodcastAuthor `json:"authors"`
	}

	// Подкаст.
	Podcast struct {
		SimplePodcast

		// Подкаст лайкнут?
		CollectionItemData CollectionItem `json:"collectionItemData"`

		Description  string `json:"description"`
		UpdatedDate  Time   `json:"updatedDate"`
		Availability int    `json:"availability"`
		Type         string `json:"type"`
		Episodes     []struct {
			ID ID `json:"id"`
		} `json:"episodes"`
	}

	// Краткая информация об авторе подкаста.
	PodcastAuthor struct {
		ID   ID     `json:"id"`
		Name string `json:"name"`
	}

	// Краткая информация об эпизоде подкаста.
	SimpleEpisode struct {
		ID              ID     `json:"id"`
		Title           string `json:"title"`
		Explicit        bool   `json:"explicit"`
		Duration        int    `json:"duration"`
		PublicationDate Time   `json:"publicationDate"`
		Image           Image  `json:"image"`
	}

	// Эпизод подкаста.
	Episode struct {
		SimpleEpisode

		// Эпизод лайкнут?
		CollectionItemData CollectionItem `json:"collectionItemData"`
		Description        string         `json:"description"`
		Availability       int            `json:"availability"`
		Podcast            SimplePodcast  `json:"podcast"`
	}
)
