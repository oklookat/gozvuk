package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Лайкнуть артиста, трек, и так далее.
func (c Client) AddItemToCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType) error {
	return addRemoveFromCollection(ctx, id, itype, c, true)
}

// Снять лайк с артиста, трека, и так далее.
func (c Client) RemoveItemFromCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType) error {
	return addRemoveFromCollection(ctx, id, itype, c, false)
}

func addRemoveFromCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType, cl Client, add bool) error {
	var body string
	var err error

	if add {
		if body, err = schema.AddItemToCollection(id, itype); err != nil {
			return err
		}
	} else {
		if body, err = schema.RemoveItemFromCollection(id, itype); err != nil {
			return err
		}
	}

	result := &schema.Response[any]{}
	resp, err := cl.Http.R().SetJsonString(body).SetResult(result).Post(ctx, genApiPath())
	if err != nil {
		return err
	}

	return checkResponse(resp, result)
}
