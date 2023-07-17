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
)

func GetPlaylists(ids []ID, shortTrackList, uniqueReleases bool,
	first int) (string, error) {
	return getGraphqlBody(_getPlaylists, "getPlaylists", map[string]any{
		"shortTrackList": shortTrackList,
		"first":          first,
		"uniqueReleases": uniqueReleases,
		"ids":            ids,
	})
}

type GetPlaylistsResponse struct {
	GetPlaylists []Playlist `json:"getPlaylists"`
}

func GetPlaylistTracks(id ID, limit, offset int, withStream bool) (string, error) {
	return getGraphqlBody(_getPlaylistTracks, "getPlaylistTracks", map[string]any{
		"limit":      limit,
		"offset":     offset,
		"withStream": withStream,
		"id":         id,
	})
}

type GetPlaylistTracksResponse struct {
	GetPlaylistTracks []Track `json:"playlistTracks"`
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

func DeletePlaylist(id ID) (string, error) {
	return getGraphqlBody(_deletePlaylist, "deletePlaylist", map[string]any{
		"id": id,
	})
}

type DeletePlaylistResponse struct {
	Playlist struct {
		Delete *ID `json:"delete"`
	} `json:"playlist"`
}

func RenamePlaylist(id ID, name string) (string, error) {
	return getGraphqlBody(_renamePlaylist, "renamePlaylist", map[string]any{
		"id":   id,
		"name": name,
	})
}

type RenamePlaylistResponse struct {
	Playlist struct {
		Rename struct {
			ID   *ID     `json:"id"`
			Name *string `json:"name"`
		} `json:"rename"`
	} `json:"playlist"`
}

func SetPlaylistToPublic(id ID, isPublic bool) (string, error) {
	return getGraphqlBody(_setPlaylistToPublic, "setPlaylistToPublic", map[string]any{
		"id":       id,
		"isPublic": isPublic,
	})
}

type SetPlaylistToPublicResponse struct {
	Playlist struct {
		SetPublic struct {
			ID       *ID   `json:"id"`
			IsPublic *bool `json:"isPublic"`
		} `json:"setPublic"`
	} `json:"playlist"`
}

func GetShortPlaylist(ids []ID, first int, uniqueReleases bool) (string, error) {
	return getGraphqlBody(_getShortPlaylist, "getShortPlaylist", map[string]any{
		"first":          first,
		"uniqueReleases": uniqueReleases,
		"ids":            ids,
	})
}

type GetShortPlaylistResponse struct {
	GetPlaylists []Playlist `json:"getPlaylists"`
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

type (
	// Плейлист.
	Playlist struct {
		ID ID `json:"id"`

		// Название.
		Title string `json:"title"`

		// Описание.
		Description *string `json:"description"`

		// Я не знаю в чем разница между IsPublic и Is_Public. Легаси?
		Is_Public *bool `json:"is_public"`

		// Плейлист публичный?
		IsPublic *bool `json:"isPublic"`

		CollectionLastModified *Time `json:"collectionLastModified"`
		Updated                *Time `json:"updated"`

		// ID автора.
		UserID *string `json:"userId"`

		Image *Image `json:"image"`

		// Треки.
		Tracks []Track `json:"tracks"`

		// ???.
		Ftracks []Track `json:"ftracks"`

		// В секундах.
		Duration *int `json:"duration"`

		Branded *bool `json:"branded"`

		Cover   any    `json:"cover"`
		CoverV1 *Image `json:"coverV1"`

		SearchTitle *string `json:"searchTitle"`
		Buttons     []any   `json:"buttons"`
		IsDeleted   *bool   `json:"isDeleted"`
		Shared      *bool   `json:"shared"`
	}

	SynthesisPlaylist struct {
		ID      ID       `json:"id"`
		Tracks  []Track  `json:"tracks"`
		Authors []Author `json:"authors"`
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
