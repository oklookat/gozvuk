package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Лайкнуть артиста, трек, и так далее.
func (c Client) AddItemToCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveFromCollection(ctx, id, itype, c, true)
}

// Снять лайк с артиста, трека, и так далее.
func (c Client) RemoveItemFromCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveFromCollection(ctx, id, itype, c, false)
}

// Добавить трек, артиста, etc, из скрытых.
func (c Client) AddItemToHidden(ctx context.Context,
	id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveHidden(ctx, id, itype, c, true)
}

// Убрать трек, артиста, etc, из скрытых.
func (c Client) RemoveItemFromHidden(ctx context.Context,
	id schema.ID, itype schema.CollectionItemType) (*schema.Response[any], error) {
	return addRemoveHidden(ctx, id, itype, c, false)
}

// Получить треки, артистов, etc, скрытых текущим пользователем.
func (c Client) GetAllHiddenCollection(ctx context.Context) (*schema.Response[schema.GetAllHiddenCollectionResponse], error) {
	body, err := schema.GetAllHiddenCollection()
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetAllHiddenCollectionResponse](ctx, c, body)
}

// Получить треки скрытые текущим пользователем.
func (c Client) GetHiddenTracks(ctx context.Context) (*schema.Response[schema.GetHiddenTracksResponse], error) {
	body, err := schema.GetHiddenTracks()
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetHiddenTracksResponse](ctx, c, body)
}

func addRemoveHidden(ctx context.Context, id schema.ID, itype schema.CollectionItemType, cl Client, add bool) (*schema.Response[any], error) {
	var body string
	var err error

	if add {
		if body, err = schema.AddItemToHidden(id, itype); err != nil {
			return nil, err
		}
	} else {
		if body, err = schema.RemoveItemFromHidden(id, itype); err != nil {
			return nil, err
		}
	}
	return sendRequestWithBody[any](ctx, cl, body)
}

func addRemoveFromCollection(ctx context.Context, id schema.ID, itype schema.CollectionItemType, cl Client, add bool) (*schema.Response[any], error) {
	var body string
	var err error

	if add {
		if body, err = schema.AddItemToCollection(id, itype); err != nil {
			return nil, err
		}
	} else {
		if body, err = schema.RemoveItemFromCollection(id, itype); err != nil {
			return nil, err
		}
	}
	return sendRequestWithBody[any](ctx, cl, body)
}
