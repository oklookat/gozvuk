package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить релизы по ID. relatedLimit ниже 1 будет выдавать ошибку валидации.
func (c Client) GetReleases(ctx context.Context, ids []schema.ID, relatedLimit int) (*schema.Response[schema.GetReleasesResponse], error) {
	body, err := schema.GetReleases(ids, relatedLimit)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetReleasesResponse](ctx, c, body)
}
