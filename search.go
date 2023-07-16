package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Поиск.
func (c Client) Search(ctx context.Context, args schema.SearchArguments) (*schema.Response[schema.SearchResponse], error) {
	body, err := schema.SearchQ(args)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.SearchResponse](ctx, c, body)
}

// Быстрый поиск. Я не знаю зачем нужен searchSessionId, но если его оставить пустым
// то кажется ничего не будет.
func (c Client) QuickSearch(ctx context.Context, query string, limit int, searchSessionId string) (*schema.Response[schema.QuickSearchResponse], error) {
	body, err := schema.QuickSearchQ(query, limit, searchSessionId)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.QuickSearchResponse](ctx, c, body)
}
