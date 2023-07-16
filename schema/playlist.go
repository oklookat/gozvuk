package schema

import (
	_ "embed"
	"time"
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

type (
	// Плейлист.
	Playlist struct {
		ID          ID         `json:"id"`
		Title       string     `json:"title"`
		SearchTitle *string    `json:"search_title"`
		Updated     *time.Time `json:"updated"`
		Description *string    `json:"description"`
		Image       *Image     `json:"image"`
		Tracks      []Track    `json:"tracks"`
		Ftracks     []any      `json:"ftracks"`
		Buttons     []any      `json:"buttons"`
		Branded     *bool      `json:"branded"`
		Cover       any        `json:"cover"`

		// Я не знаю в чем разница между IsPublic и Is_Public.
		Is_Public *bool `json:"is_public"`
		IsPublic  *bool `json:"isPublic"`

		Duration  *int    `json:"duration"`
		IsDeleted *bool   `json:"isDeleted"`
		UserID    *string `json:"userId"`
		Shared    *bool   `json:"shared"`
	}
)

func NewPlaylistItem(itype PlaylistItemType, item ID) PlaylistItem {
	return PlaylistItem{
		Type:   itype,
		ItemID: item,
	}
}

type PlaylistItemType string

type PlaylistItem struct {
	Type   PlaylistItemType `json:"type"`
	ItemID ID               `json:"item_id"`
}

const (
	PlaylistItemTypeTrack PlaylistItemType = "track"
)
