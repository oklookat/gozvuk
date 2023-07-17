package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

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

// Получить кол-во подписок пользователя.
func (c Client) FollowingCount(ctx context.Context, id schema.ID) (*schema.Response[schema.FollowingCountResponse], error) {
	body, err := schema.FollowingCount(id)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.FollowingCountResponse](ctx, c, body)
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
