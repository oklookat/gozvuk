package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить треки по ID.
func (c Client) GetTracks(ctx context.Context, ids []schema.ID) (*schema.Response[schema.GetTracksResponse], error) {
	body, err := schema.GetTracks(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetTracksResponse](ctx, c, body)
}

// Получить полные треки по ID.
func (c Client) GetFullTrack(ctx context.Context, ids []schema.ID,
	withReleases, withArtists bool) (*schema.Response[schema.GetFullTrackResponse], error) {
	body, err := schema.GetFullTrack(ids, withReleases, withArtists)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetFullTrackResponse](ctx, c, body)
}
