package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
	"github.com/oklookat/gozvuk/schema/gql"
)

func (c Client) GetArtists(ctx context.Context, ids []string,
	withReleases, withPopTracks, withRelatedArtists, withDescription bool,
	releasesOffset, releasesLimit, tracksOffset, tracksLimit, relatedArtistsLimit int) (*schema.Response[gql.GetArtistsResponse], error) {

	body, err := gql.GetArtists(ids, withReleases, withPopTracks, withRelatedArtists,
		withDescription, releasesOffset, releasesLimit, tracksOffset, tracksLimit, relatedArtistsLimit)
	if err != nil {
		return nil, err
	}

	result := &schema.Response[gql.GetArtistsResponse]{}
	resp, err := c.Http.R().SetJsonString(body).SetResult(result).Post(ctx, genApiPath())
	if err != nil {
		return result, err
	}

	return result, checkResponse(resp, result)
}
