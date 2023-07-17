package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить плейлисты по ID.
func (c Client) GetPlaylists(ctx context.Context, ids []schema.ID) (*schema.Response[schema.GetPlaylistsResponse], error) {
	body, err := schema.GetPlaylists(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetPlaylistsResponse](ctx, c, body)
}

// Получить треки плейлиста.
func (c Client) GetPlaylistTracks(ctx context.Context, id schema.ID,
	limit, offset int) (*schema.Response[schema.GetPlaylistTracksResponse], error) {
	body, err := schema.GetPlaylistTracks(id, limit, offset)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetPlaylistTracksResponse](ctx, c, body)
}

// Создать плейлист. По умолчанию создается публичный плейлист.
func (c Client) CreatePlaylist(ctx context.Context, items []schema.PlaylistItem, name string) (*schema.Response[schema.CreatePlaylistResponse], error) {
	body, err := schema.CreatePlaylist(items, name)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.CreatePlaylistResponse](ctx, c, body)
}

// Удалить плейлист.
func (c Client) DeletePlaylist(ctx context.Context, id schema.ID) (*schema.Response[any], error) {
	body, err := schema.DeletePlaylist(id)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[any](ctx, c, body)
}

// Переименовать плейлист.
func (c Client) RenamePlaylist(ctx context.Context, id schema.ID, newName string) (*schema.Response[any], error) {
	body, err := schema.RenamePlaylist(id, newName)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[any](ctx, c, body)
}

// Добавить треки в плейлист.
func (c Client) AddTracksToPlaylist(ctx context.Context, id schema.ID, items []schema.PlaylistItem) (*schema.Response[any], error) {
	body, err := schema.AddTracksToPlaylist(id, items)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[any](ctx, c, body)
}

// Обновить плейлист. Например обновить треки.
//
// Внимание: плейлист обновится целиком. То есть если нужно
// удалить трек(и), то в items должны быть все треки из плейлиста,
// кроме тех которые нужно удалить.
func (c Client) UpdataPlaylist(ctx context.Context, id schema.ID, items []schema.PlaylistItem, isPublic bool, name string) (*schema.Response[any], error) {
	body, err := schema.UpdataPlaylist(id, items, isPublic, name)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[any](ctx, c, body)
}

// Изменить видимость плейлиста.
func (c Client) SetPlaylistToPublic(ctx context.Context, id schema.ID, isPublic bool) (*schema.Response[any], error) {
	body, err := schema.SetPlaylistToPublic(id, isPublic)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[any](ctx, c, body)
}

// Получить плейлисты в коротком варианте.
func (c Client) GetShortPlaylist(ctx context.Context, ids []schema.ID) (*schema.Response[schema.GetShortPlaylistResponse], error) {
	body, err := schema.GetShortPlaylist(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.GetShortPlaylistResponse](ctx, c, body)
}

// Создать (получить?) плейлист на основе вкусов двух аккаунтов.
//
// Плейлист будет доступен по ссылке типа: https://zvuk.com/playlist/synthesis/ID_созданного_плейлиста
//
// Вы не будете автором этого плейлиста.
//
// Если установить одинаковые ID, то будет ошибка валидации.
func (c Client) SynthesisPlaylistBuild(ctx context.Context, firstAuthor, secondAuthor schema.ID) (*schema.Response[schema.SynthesisPlaylistBuildResponse], error) {
	body, err := schema.SynthesisPlaylistBuild(firstAuthor, secondAuthor)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.SynthesisPlaylistBuildResponse](ctx, c, body)
}

// Получить синтезированные плейлисты по их ID.
//
// Перед вызовом нужно(?) получить ID плейлистов через SynthesisPlaylistBuild. Иначе ответ будет пустой.
func (c Client) SynthesisPlaylist(ctx context.Context, ids []schema.ID) (*schema.Response[schema.SynthesisPlaylistResponse], error) {
	body, err := schema.SynthesisPlaylistQ(ids)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.SynthesisPlaylistResponse](ctx, c, body)
}
