package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

func (c Client) GetArtists(ctx context.Context, ids []schema.ID,
	withReleases, withPopTracks, withRelatedArtists, withDescription bool,
	releasesOffset, releasesLimit, tracksOffset, tracksLimit, relatedArtistsLimit int) (*schema.Response[schema.GetArtistsResponse], error) {
	body, err := schema.GetArtists(ids, withReleases, withPopTracks, withRelatedArtists,
		withDescription, releasesOffset, releasesLimit, tracksOffset, tracksLimit, relatedArtistsLimit)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetArtistsResponse](ctx, c, body)
}

// Получить кол-во подписчиков на профиль артиста.
func (c Client) ArtistFollowersCount(ctx context.Context, ids []schema.ID) (*schema.Response[schema.ArtistFollowersCountResponse], error) {
	body, err := schema.ArtistFollowersCount(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.ArtistFollowersCountResponse](ctx, c, body)
}
