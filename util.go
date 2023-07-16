package gozvuk

import (
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
	ErrNilResponse = errors.New(errPrefix + "nil http or schema response (???)")

	// Ответ с ошибкой, но поля Error в ответе нет.
	ErrNilResponseError = errors.New(errPrefix + "nil Response.Error (API changed?)")
)

func checkResponse[T any](resp *vantuz.Response, data *schema.Response[T]) error {
	if resp == nil || data == nil {
		return ErrNilResponse
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
	return ErrNilResponseError
}
