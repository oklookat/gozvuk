package gozvuk

import (
	"github.com/oklookat/gozvuk/schema"
	"github.com/oklookat/vantuz"
)

const (
	errPrefix = "gozvuk: "
)

// Получить Client для запросов к API.
//
// accessToken можно получить тут:
// https://zvuk.com/api/tiny/profile
func New(accessToken string) *Client {
	httpCl := vantuz.C().SetGlobalHeader("X-Auth-Token", accessToken)
	cl := &Client{
		accessToken: accessToken,
		Http:        httpCl,
	}
	cl.SetUserAgent("gozvuk")
	return cl
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
