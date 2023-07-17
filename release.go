package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить релизы по ID.
func (c Client) GetReleases(ctx context.Context, ids []schema.ID,
	withTracks bool, withArtists bool,
	withLabel bool, withRelated bool, relatedLimit int) (*schema.Response[schema.GetReleasesResponse], error) {
	body, err := schema.GetReleases(ids, withTracks, withArtists, withLabel, withRelated, relatedLimit)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetReleasesResponse](ctx, c, body)
}
