package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

func (c Client) GetArtists(ctx context.Context, ids []schema.ID,
	withReleases bool, releasesLimit, releasesOffset int,
	withPopTracks bool, tracksLimit, tracksOffset int,
	withRelatedArtists bool, relatedArtistsLimit int,
	withDescription bool,
) (*schema.Response[schema.GetArtistsResponse], error) {
	body, err := schema.GetArtists(ids, withReleases, withPopTracks, withRelatedArtists,
		withDescription, releasesOffset, releasesLimit, tracksOffset, tracksLimit, relatedArtistsLimit)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetArtistsResponse](ctx, c, body)
}
