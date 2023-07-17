package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить подкасты по ID.
func (c Client) GetPodcasts(ctx context.Context, ids []schema.ID) (*schema.Response[schema.GetPodcastsResponse], error) {
	body, err := schema.GetPodcasts(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetPodcastsResponse](ctx, c, body)
}

// Получить эпизоды подкастов по их ID.
func (c Client) GetEpisodes(ctx context.Context, ids []schema.ID) (*schema.Response[schema.GetEpisodesResponse], error) {
	body, err := schema.GetEpisodes(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetEpisodesResponse](ctx, c, body)
}
