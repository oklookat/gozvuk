package schema

import (
	_ "embed"
)

var (
	//go:embed gql/getPlaylists.gql
	_getPlaylists string
	//go:embed gql/getPlaylistTracks.gql
	_getPlaylistTracks string
	//go:embed gql/createPlaylist.gql
	_createPlaylist string
	//go:embed gql/deletePlaylist.gql
	_deletePlaylist string
	//go:embed gql/renamePlaylist.gql
	_renamePlaylist string
	//go:embed gql/setPlaylistToPublic.gql
	_setPlaylistToPublic string
	//go:embed gql/getShortPlaylist.gql
	_getShortPlaylist string
	//go:embed gql/synthesisPlaylistBuild.gql
	_synthesisPlaylistBuild string
	//go:embed gql/synthesisPlaylist.gql
	_synthesisPlaylist string
	//go:embed gql/addTracksToPlaylist.gql
	_addTracksToPlaylist string
	//go:embed gql/updataPlaylist.gql
	_updataPlaylist string
)

func GetPlaylists(ids []ID) (string, error) {
	return getGraphqlBody(_getPlaylists, "getPlaylists", map[string]any{
		"ids": ids,
	})
}

type GetPlaylistsResponse struct {
	GetPlaylists []Playlist `json:"getPlaylists"`
}

func GetShortPlaylist(ids []ID) (string, error) {
	return getGraphqlBody(_getShortPlaylist, "getShortPlaylist", map[string]any{
		"ids": ids,
	})
}

type GetShortPlaylistResponse struct {
	GetPlaylists []SimplePlaylist `json:"getPlaylists"`
}

func GetPlaylistTracks(id ID, limit, offset int) (string, error) {
	return getGraphqlBody(_getPlaylistTracks, "getPlaylistTracks", map[string]any{
		"limit":  limit,
		"offset": offset,
		"id":     id,
	})
}

type GetPlaylistTracksResponse struct {
	GetPlaylistTracks []SimpleTrack `json:"playlistTracks"`
}

func CreatePlaylist(items []PlaylistItem, name string) (string, error) {
	return getGraphqlBody(_createPlaylist, "createPlayList", map[string]any{
		"name":  name,
		"items": items,
	})
}

type CreatePlaylistResponse struct {
	Playlist struct {
		Create ID `json:"create"`
	} `json:"playlist"`
}

func SynthesisPlaylistBuild(firstAuthor, secondAuthor ID) (string, error) {
	return getGraphqlBody(_synthesisPlaylistBuild, "synthesisPlaylistBuild", map[string]any{
		"firstAuthorId":  firstAuthor,
		"secondAuthorId": secondAuthor,
	})
}

type SynthesisPlaylistBuildResponse struct {
	SynthesisPlaylistBuild SynthesisPlaylist `json:"synthesisPlaylistBuild"`
}

func SynthesisPlaylistQ(ids []ID) (string, error) {
	return getGraphqlBody(_synthesisPlaylist, "synthesisPlaylist", map[string]any{
		"ids": ids,
	})
}

type SynthesisPlaylistResponse struct {
	SynthesisPlaylist []SynthesisPlaylist `json:"synthesisPlaylist"`
}

func DeletePlaylist(id ID) (string, error) {
	return getGraphqlBody(_deletePlaylist, "deletePlaylist", map[string]any{
		"id": id,
	})
}

func RenamePlaylist(id ID, name string) (string, error) {
	return getGraphqlBody(_renamePlaylist, "renamePlaylist", map[string]any{
		"id":   id,
		"name": name,
	})
}

func SetPlaylistToPublic(id ID, isPublic bool) (string, error) {
	return getGraphqlBody(_setPlaylistToPublic, "setPlaylistToPublic", map[string]any{
		"id":       id,
		"isPublic": isPublic,
	})
}

func AddTracksToPlaylist(id ID, items []PlaylistItem) (string, error) {
	return getGraphqlBody(_addTracksToPlaylist, "addTracksToPlaylist", map[string]any{
		"id":    id,
		"items": items,
	})
}

func UpdataPlaylist(id ID, items []PlaylistItem, isPublic bool, name string) (string, error) {
	return getGraphqlBody(_updataPlaylist, "updataPlaylist", map[string]any{
		"id":       id,
		"items":    items,
		"isPublic": isPublic,
		"name":     name,
	})
}

type (
	// Краткая информация о плейлисте.
	SimplePlaylist struct {
		ID ID `json:"id"`

		// Название.
		Title string `json:"title"`

		// Плейлист публичный?
		IsPublic bool `json:"isPublic"`

		// Описание.
		Description string `json:"description"`

		// Общая длина треков в секундах.
		Duration int `json:"duration"`

		// Обложка.
		Image *Image `json:"image"`
	}

	// Плейлист.
	Playlist struct {
		SimplePlaylist

		// Треки.
		Tracks []struct {
			ID ID `json:"id"`
		} `json:"tracks"`

		// ID автора.
		UserID    string `json:"userId"`
		IsDeleted bool   `json:"isDeleted"`
		Shared    bool   `json:"shared"`
		Branded   bool   `json:"branded"`

		// Дата обновления.
		Updated     Time   `json:"updated"`
		SearchTitle string `json:"searchTitle"`
	}

	SynthesisPlaylist struct {
		ID      ID               `json:"id"`
		Tracks  []SimpleTrack    `json:"tracks"`
		Authors []PlaylistAuthor `json:"authors"`
	}

	PlaylistAuthor struct {
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

func NewPlaylistItem(itype PlaylistItemType, item ID) PlaylistItem {
	return PlaylistItem{
		Type:   itype,
		ItemID: item,
	}
}

// Тип сущности в плейлисте.
type PlaylistItemType string

// Сущность в плейлисте.
type PlaylistItem struct {
	// Тип сущности.
	Type PlaylistItemType `json:"type"`

	// ID сущности.
	ItemID ID `json:"item_id"`
}

const (
	// Трек.
	PlaylistItemTypeTrack PlaylistItemType = "track"
)
