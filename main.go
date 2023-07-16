package gozvuk

import (
	"context"

	"github.com/oklookat/gozvuk/schema"
	"github.com/oklookat/vantuz"
)

const (
	errPrefix = "gozvuk: "
)

// Получить Client для запросов к API.
//
// accessToken можно получить тут: https://zvuk.com/api/tiny/profile
func New(accessToken string) (*Client, error) {
	httpCl := vantuz.C().SetGlobalHeader("X-Auth-Token", accessToken)
	cl := &Client{
		accessToken: accessToken,
		Http:        httpCl,
	}
	cl.SetUserAgent("gozvuk")

	// Ping.
	profile, err := cl.Profile()
	if err != nil {
		return nil, err
	}
	if profile.Result.IsAnonymous != nil && *profile.Result.IsAnonymous {
		return nil, schema.ErrNotAuthorized
	}

	return cl, err
}

// Клиент для запросов к API.
type Client struct {
	accessToken string

	// ID текущего пользователя.
	UserId schema.ID

	// Отправляет запросы.
	Http *vantuz.Client
}

// Установить user agent для запросов.
func (c *Client) SetUserAgent(val string) {
	c.Http.SetUserAgent(val)
}

func (c *Client) Profile() (*schema.Profile, error) {
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
