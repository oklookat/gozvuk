package schema

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

const (
	errPrefix = "gozvuk/schema: "

	ApiUrl     = "https://zvuk.com/api/v1/graphql"
	ApiTinyUrl = "https://zvuk.com/api/tiny"
	ZvukUrl    = "https://zvuk.com"
)

var (
	ErrNotAuthorized = errors.New(errPrefix + "bad or expired token")
)

type ID string

func (i ID) String() string {
	return string(i)
}

type Error struct {
	Detail string `json:"detail"`
}

func (e Error) Error() string {
	return errPrefix + e.Detail
}

// Картинка.
type Image struct {
	// Высота.
	H *int `json:"h"`
	// Пример: #ffffff,#fcfcfb,#282422
	Palette string `json:"palette"`
	// Пример: #ffffff,#2b2726,#887f70
	PaletteBottom string `json:"palette_bottom"`
	// Пример: https://cdn51.zvuk.com/pic?type=release&id=3322849&ext=jpg&size={size}
	Src string `json:"src"`
	// Ширина.
	W *int `json:"w"`
}

func (i Image) URL(width, height int) *url.URL {
	parsed, _ := url.Parse(i.Src)
	if parsed == nil {
		return nil
	}
	w := strconv.Itoa(width)
	h := strconv.Itoa(height)
	query := parsed.Query()
	query.Set("size", w+"x"+h)
	parsed.RawQuery = query.Encode()
	return parsed
}

type (
	// Ответ.
	Response[T any] struct {
		// Ошибки.

		StatusCode *int    `json:"statusCode"`
		Error      *string `json:"error"`
		Errors     []struct {
			Message string `json:"message"`
		} `json:"errors"`
		Message *string `json:"message"`

		// Результат запроса.

		Data T `json:"data"`
	}

	Release struct {
		ID             string  `json:"id"`
		Title          string  `json:"title"`
		SearchTitle    *string `json:"searchTitle"`
		Type           *string `json:"type"`
		Date           *Time   `json:"date"`
		Image          *Image  `json:"image"`
		Availability   *int    `json:"availability"`
		ArtistTemplate *string `json:"artistTemplate"`
		Artists        []ShortArtist
	}
)

type Time struct {
	time.Time
}

// Парс времени.
func (t *Time) UnmarshalJSON(data []byte) error {
	var timeStr string
	err := json.Unmarshal(data, &timeStr)
	if err != nil {
		return err
	}
	if len(timeStr) == 0 {
		return nil
	}

	tt, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		*t = Time{tt}
		return nil
	}

	// Например у релизов такой формат даты.
	timeLayout := "2006-01-02T15:04:05"
	tt, err = time.Parse(timeLayout, timeStr)
	if err == nil {
		*t = Time{tt}
		return nil
	}

	return err
}
