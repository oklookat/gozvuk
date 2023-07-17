package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
)

// Получить текущий аккаунт.
func (c Client) Profile() (*schema.Profile, error) {
	var respErr schema.Error
	var result schema.Profile
	_, err := c.Http.R().SetError(&respErr).SetResult(&result).Get(context.Background(), genApiPathTiny("profile"))
	if err != nil {
		return nil, wrapError(err)
	}
	if len(respErr.Detail) == 0 {
		return &result, nil
	}
	return &result, respErr
}

// Валидный токен / пользователь авторизован?
func (c Client) IsAuthorized() (bool, error) {
	profile, err := c.Profile()
	if err != nil {
		return false, err
	}
	if profile.Result.IsAnonymous != nil && *profile.Result.IsAnonymous {
		return false, nil
	}
	return true, nil
}

// Получить кол-во подписок пользователя.
func (c Client) FollowingCount(ctx context.Context, id schema.ID) (*schema.Response[schema.FollowingCountResponse], error) {
	body, err := schema.FollowingCount(id)
	if err != nil {
		return nil, err
	}
	return sendRequestWithBody[schema.FollowingCountResponse](ctx, c, body)
}
