package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Лайкнуть артиста, трек, и так далее.
func (c Client) AddItemToCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveFromCollection(ctx, id, itype, c, true)
}

// Снять лайк с артиста, трека, и так далее.
func (c Client) RemoveItemFromCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveFromCollection(ctx, id, itype, c, false)
}

// Добавить трек, артиста, etc, из скрытых.
func (c Client) AddItemToHidden(ctx context.Context,
	id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveHidden(ctx, id, itype, c, true)
}

// Убрать трек, артиста, etc, из скрытых.
func (c Client) RemoveItemFromHidden(ctx context.Context,
	id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveHidden(ctx, id, itype, c, false)
}

// Получить треки, артистов, etc, скрытых текущим пользователем.
func (c Client) GetAllHiddenCollection(ctx context.Context) (*schema.Response[schema.GetAllHiddenCollectionResponse], error) {
	body, err := schema.GetAllHiddenCollection()
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetAllHiddenCollectionResponse](ctx, c, body)
}

// Получить треки скрытые текущим пользователем.
func (c Client) GetHiddenTracks(ctx context.Context) (*schema.Response[schema.GetHiddenTracksResponse], error) {
	body, err := schema.GetHiddenTracks()
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetHiddenTracksResponse](ctx, c, body)
}

// Получить лайкнутые треки, альбомы, и так далее.
func (c Client) UserCollection(ctx context.Context) (*schema.Response[schema.UserCollectionResponse], error) {
	body, err := schema.UserColelctionQ()
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.UserCollectionResponse](ctx, c, body)
}

// Получить плейлисты из коллекции текущего пользователя.
// То есть и лайкнутые плейлисты, и свои.
func (c Client) UserPlaylists(ctx context.Context) (*schema.Response[schema.UserPlaylistsResponse], error) {
	body, err := schema.UserPlaylists()
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.UserPlaylistsResponse](ctx, c, body)
}

// Получить лайкнутые треки из коллекции текущего пользователя.
func (c Client) UserTracks(ctx context.Context,
	by schema.OrderBy, direction schema.OrderDirection) (*schema.Response[schema.UserTracksResponse], error) {
	body, err := schema.UserTracks(direction, by)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.UserTracksResponse](ctx, c, body)
}

// Получить кол-во подписчиков пользователей.
func (c Client) ProfileFollowersCount(ctx context.Context, ids []schema.ID) (*schema.Response[schema.ProfileFollowersCountResponse], error) {
	body, err := schema.ProfileFollowersCount(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.ProfileFollowersCountResponse](ctx, c, body)
}

// Получить подкасты из коллекции пользователя с разбивкой на страницы.
func (c Client) UserPaginatedPodcasts(ctx context.Context, cursor string, count int) (*schema.Response[schema.UserPaginatedPodcastsResponse], error) {
	body, err := schema.UserPaginatedPodcasts(cursor, count)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.UserPaginatedPodcastsResponse](ctx, c, body)
}

func addRemoveHidden(ctx context.Context, id schema.ID, itype schema.CollectionItemType, cl Client, add bool) (*schema.Response[any], error) {
	var body string
	var err error

	if add {
		if body, err = schema.AddItemToHidden(id, itype); err != nil {
			return nil, err
		}
	} else {
		if body, err = schema.RemoveItemFromHidden(id, itype); err != nil {
			return nil, err
		}
	}
	return sendRequestWithBody[any](ctx, cl, body)
}

func addRemoveFromCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType, cl Client, add bool) (*schema.Response[any], error) {
	var body string
	var err error

	if add {
		if body, err = schema.AddItemToCollection(id, itype); err != nil {
			return nil, err
		}
	} else {
		if body, err = schema.RemoveItemFromCollection(id, itype); err != nil {
			return nil, err
		}
	}
	return sendRequestWithBody[any](ctx, cl, body)
}
