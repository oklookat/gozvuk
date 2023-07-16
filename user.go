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
