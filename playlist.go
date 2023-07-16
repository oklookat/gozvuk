package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить плейлисты по ID.
func (c Client) GetPlaylists(ctx context.Context, ids []schema.ID, shortTrackList, uniqueReleases bool,
	first int) (*schema.Response[schema.GetPlaylistsResponse], error) {
	body, err := schema.GetPlaylists(ids, shortTrackList, uniqueReleases, first)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetPlaylistsResponse](ctx, c, body)
}

// Получить треки плейлиста.
func (c Client) GetPlaylistTracks(ctx context.Context, id schema.ID,
	limit, offset int, withStream bool) (*schema.Response[schema.GetPlaylistTracksResponse], error) {
	body, err := schema.GetPlaylistTracks(id, limit, offset, withStream)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetPlaylistTracksResponse](ctx, c, body)
}

// Создать плейлист. По умолчанию создается публичный плейлист.
func (c Client) CreatePlaylist(ctx context.Context, items []schema.PlaylistItem, name string) (*schema.Response[schema.CreatePlaylistResponse], error) {
	body, err := schema.CreatePlaylist(items, name)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.CreatePlaylistResponse](ctx, c, body)
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, id schema.ID) (*schema.Response[schema.DeletePlaylistResponse], error) {
	body, err := schema.DeletePlaylist(id)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.DeletePlaylistResponse](ctx, c, body)
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, id schema.ID, newName string) (*schema.Response[schema.RenamePlaylistResponse], error) {
	body, err := schema.RenamePlaylist(id, newName)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.RenamePlaylistResponse](ctx, c, body)
}

// Изменить видимость плейлиста.
func (c Client) SetPlaylistToPublic(ctx context.Context, id schema.ID, isPublic bool) (*schema.Response[schema.SetPlaylistToPublicResponse], error) {
	body, err := schema.SetPlaylistToPublic(id, isPublic)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.SetPlaylistToPublicResponse](ctx, c, body)
}
