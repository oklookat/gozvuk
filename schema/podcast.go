package schema

import (
	_ "embed"
)

var (
	//go:embed gql/getPodcasts.gql
	_getPodcasts string
	//go:embed gql/getEpisodes.gql
	_getEpisodes string
	//go:embed gql/episodes.gql
	_episodes string
)

func GetPodcasts(ids []ID, withEpisodes bool) (string, error) {
	return getGraphqlBody(_getPodcasts, "getPodcasts", map[string]any{
		"withEpisodes": withEpisodes,
		"ids":          ids,
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

func Episodes(ids []ID) (string, error) {
	return getGraphqlBody(_episodes, "episodes", map[string]any{
		"ids": ids,
	})
}

type EpisodesResponse struct {
	GetEpisodes []Episode `json:"getEpisodes"`
}

type (
	// Краткая информация о подкасте.
	SimplePodcast struct {
		ID       ID                    `json:"id"`
		Title    string                `json:"title"`
		Explicit bool                  `json:"explicit"`
		Image    Image                 `json:"image"`
		Authors  []SimplePodcastAuthor `json:"authors"`
	}

	// Подкаст.
	Podcast struct {
		SimplePodcast

		// Подкаст лайкнут?
		CollectionItemData CollectionItem `json:"collectionItemData"`

		Description  *string   `json:"description"`
		UpdatedDate  *Time     `json:"updatedDate"`
		Availability *int      `json:"availability"`
		Type         *string   `json:"type"`
		Episodes     []Episode `json:"episodes"`
	}

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

		Description  *string  `json:"description"`
		Availability *int     `json:"availability"`
		Season       *Season  `json:"season"`
		Podcast      *Podcast `json:"podcast"`
		Link         any      `json:"link"`
		ListenState  any      `json:"listenState"`
		TrackId      any      `json:"trackId"`
	}

	// Сезон подкаста.
	Season struct {
		ID           ID     `json:"id"`
		Name         string `json:"name"`
		SeasonNumber int    `json:"seasonNumber"`
	}

	SimplePodcastAuthor struct {
		ID   ID     `json:"id"`
		Name string `json:"name"`
	}
)
