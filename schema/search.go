package schema

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed gql/search.gql
	_search string
	//go:embed gql/quickSearch.gql
	_quickSearch string
)

// Аргументы для поиска.
type SearchArguments struct {
	// Запрос.
	Query string

	// Максимальное кол-во результатов.
	Limit int

	// Включать в результаты.

	Tracks,
	Artists,
	Releases,
	Playlists,
	Profiles,
	Books,
	Episodes,
	Podcasts,
	Categories bool

	// Для навигации по страницам.

	TrackCursor,
	ArtistsCursor,
	ReleasesCursor,
	PlaylistsCursor,
	EpisodesCursor,
	ProfilesCursor,
	PodcastsCursor *Cursor
}

func SearchQ(args SearchArguments) (string, error) {
	return getGraphqlBody(_search, "search", map[string]any{
		"limit":           args.Limit,
		"tracks":          args.Tracks,
		"trackCursor":     args.TrackCursor,
		"artistsCursor":   args.ArtistsCursor,
		"releasesCursor":  args.ReleasesCursor,
		"playlistsCursor": args.PlaylistsCursor,
		"episodesCursor":  args.EpisodesCursor,
		"profilesCursor":  args.ProfilesCursor,
		"podcastsCursor":  args.PodcastsCursor,
		"artists":         args.Artists,
		"releases":        args.Releases,
		"playlists":       args.Playlists,
		"profiles":        args.Profiles,
		"books":           args.Books,
		"episodes":        args.Episodes,
		"podcasts":        args.Podcasts,
		"categories":      args.Categories,
		"query":           args.Query,
	})
}

type SearchResponse struct {
	Search Search `json:"search"`
}

type (
	// Курсор для навигации по страницам.
	Cursor string

	// Результат поиска.
	Search struct {
		SearchID  string                        `json:"searchId"`
		Tracks    *SearchResult[SimpleTrack]    `json:"tracks"`
		Artists   *SearchResult[SimpleArtist]   `json:"artists"`
		Releases  *SearchResult[SimpleRelease]  `json:"releases"`
		Playlists *SearchResult[SimplePlaylist] `json:"playlists"`
		Profiles  *SearchResult[SimpleProfile]  `json:"profiles"`
		Books     *SearchResult[SimpleBook]     `json:"books"`
		Episodes  *SearchResult[SimpleEpisode]  `json:"episodes"`
		Podcasts  *SearchResult[SimplePodcast]  `json:"podcasts"`
	}

	// Результат поиска.
	SearchResult[T any] struct {
		// Страница.
		Page *Page `json:"page"`

		// Коэффициент совпадения от 0 до 1.0?
		Score float64 `json:"score"`

		// Результаты.
		Items []T `json:"items"`
	}

	// Страница.
	Page struct {
		// Всего.
		Total *int `json:"total"`

		// Предыдущая страница.
		Prev *int `json:"prev"`

		// Следующая страница.
		Next *int `json:"next"`

		// Курсор для навигации по страницам.
		Cursor string `json:"cursor"`
	}
)

func QuickSearchQ(query string, limit int, searchSessionId string) (string, error) {
	return getGraphqlBody(_quickSearch, "quickSearch", map[string]any{
		"query":           query,
		"limit":           limit,
		"searchSessionId": searchSessionId,
	})
}

type QuickSearchResponse struct {
	QuickSearch QuickSearch `json:"quickSearch"`
}

type QuickSearch struct {
	SearchSessionID string `json:"searchSessionId"`
	Content         []any  `json:"content"`

	// Поля ниже не входят в ответ.
	// Автоматический демаршал Content. Для удобства.

	Tracks    []SimpleTrack    `json:"-"`
	Artists   []SimpleArtist   `json:"-"`
	Releases  []SimpleRelease  `json:"-"`
	Playlists []SimplePlaylist `json:"-"`
	Episodes  []SimpleEpisode  `json:"-"`
	Podcasts  []SimplePodcast  `json:"-"`
	Profiles  []SimpleProfile  `json:"-"`
	Books     []SimpleBook     `json:"-"`
}

func (q *QuickSearch) UnmarshalJSON(data []byte) error {
	type fake QuickSearch
	var theSearch fake
	if err := json.Unmarshal(data, &theSearch); err != nil {
		return err
	}
	*q = QuickSearch(theSearch)
	for _, ent := range theSearch.Content {
		theEnt, ok := ent.(map[string]any)
		if !ok {
			continue
		}

		typeName, ok := theEnt["__typename"]
		if !ok {
			continue
		}
		typeNameStr, ok := typeName.(string)
		if !ok {
			continue
		}

		switch Typename(typeNameStr) {
		case TypenameArtist:
			var conv SimpleArtist
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Artists = append(q.Artists, conv)
			continue

		case TypenameTrack:
			var conv SimpleTrack
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Tracks = append(q.Tracks, conv)
			continue
		case TypenameRelease:
			var conv SimpleRelease
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Releases = append(q.Releases, conv)
			continue
		case TypenamePlaylist:
			var conv SimplePlaylist
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Playlists = append(q.Playlists, conv)
			continue
		case TypenameEpisode:
			var conv SimpleEpisode
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Episodes = append(q.Episodes, conv)
			continue
		case TypenamePodcast:
			var conv SimplePodcast
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Podcasts = append(q.Podcasts, conv)
			continue
		case TypenameProfile:
			var conv SimpleProfile
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Profiles = append(q.Profiles, conv)
			continue
		case TypenameBook:
			var conv SimpleBook
			if err := q.remarshal(theEnt, &conv); err != nil {
				continue
			}
			q.Books = append(q.Books, conv)
			continue
		}
	}

	return nil
}

func (q QuickSearch) remarshal(data map[string]any, to any) error {
	byted, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(byted, to)
}
