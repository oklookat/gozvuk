package gozvuk

import (
	"context"
	"errors"
	"fmt"

	"github.com/oklookat/gozvuk/schema"
	"github.com/oklookat/vantuz"
)

func genApiPath(paths ...string) string {
	return genApiPathDirty(schema.ApiUrl, paths...)
}

func genApiPathTiny(paths ...string) string {
	return genApiPathDirty(schema.ApiTinyUrl, paths...)
}

func genApiPathDirty(source string, paths ...string) string {
	if len(paths) == 0 {
		return source
	}

	base := source + "/" + paths[0]
	for i := 1; i < len(paths); i++ {
		base += "/" + paths[i]
	}

	return base
}

func wrapError(err error) error {
	return fmt.Errorf("%s%w", errPrefix, err)
}

var (
	// Странная ошибка.
	ErrUnknown = errors.New(errPrefix + "unknown")
)

func checkResponse[T any](resp *vantuz.Response, data *schema.Response[T]) error {
	if resp == nil || data == nil {
		return ErrUnknown
	}
	if data.Error != nil {
		return fmt.Errorf(errPrefix+"%s", *data.Error)
	}
	if len(data.Errors) > 0 {
		return fmt.Errorf(errPrefix+"%s", data.Errors[0].Message)
	}
	if data.Message != nil {
		return fmt.Errorf(errPrefix+"%s", *data.Message)
	}
	if data.StatusCode != nil {
		return fmt.Errorf(errPrefix+"%d", *data.StatusCode)
	}
	// Успешность проверяется последней, потому что
	// graphql может выдавать код 200, даже когда в запросе есть ошибки.
	// Например при ошибке в самом query.
	if resp.IsSuccess() {
		return nil
	}
	return ErrUnknown
}

func sendRequestWithBody[T any](ctx context.Context, cl Client, body string) (*schema.Response[T], error) {
	result := &schema.Response[T]{}
	resp, err := cl.Http.R().SetJsonString(body).SetError(result).SetResult(result).Post(ctx, genApiPath())
	if err != nil {
		return result, err
	}
	return result, checkResponse(resp, result)
}
