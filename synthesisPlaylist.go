package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

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
