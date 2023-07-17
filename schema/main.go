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
	Palette *string `json:"palette"`
	// Пример: #ffffff,#2b2726,#887f70
	PaletteBottom *string `json:"palette_bottom"`
	// Пример: https://cdn51.zvuk.com/pic?type=release&id=3322849&ext=jpg&size={size}
	Src string `json:"src"`
	// Ширина.
	W *int `json:"w"`
}

// Получить ссылку на изображение.
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
			Message    string `json:"message"`
			Extensions struct {
				Code        string `json:"code"`
				ServiceName string `json:"serviceName"`
				FieldErrors []struct {
					Loc  []string `json:"loc"`
					Msg  string   `json:"msg"`
					Type string   `json:"type"`
				} `json:"field_errors"`
				Exception struct {
					Message   string `json:"message"`
					Locations []struct {
						Line   int `json:"line"`
						Column int `json:"column"`
					} `json:"locations"`
					Path []string `json:"path"`
				} `json:"exception"`
			} `json:"extensions"`
		} `json:"errors"`
		Message *string `json:"message"`

		// Результат запроса.

		Data T `json:"data"`
	}

	Background struct {
		Type BackgroundType `json:"type"`
		// Например ссылка на изображение.
		Image    string `json:"image"`
		Color    any    `json:"color"`
		Gradient any    `json:"gradient"`
	}

	Animation struct {
		ArtistID ID              `json:"artistId"`
		Effect   AnimationEffect `json:"effect"`
		// Например ссылка на изображение.
		Image      string     `json:"image"`
		Background Background `json:"background"`
	}

	// Лейбл / мейджор.
	Label struct {
		// ID лейбла.
		ID ID `json:"id"`

		// Название лейбла.
		Title string `json:"title"`

		SearchTitle *string `json:"searchTitle"`
		Description *string `json:"description"`
		Image       *Image  `json:"image"`
		HasImage    *bool   `json:"hasImage"`
	}
)

// Дата, время.
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

// Тип сущности.
type Typename string

const (
	TypenameArtist   Typename = "Artist"
	TypenameTrack    Typename = "Track"
	TypenameRelease  Typename = "Release"
	TypenamePlaylist Typename = "Playlist"
	TypenameEpisode  Typename = "Episode"
	TypenamePodcast  Typename = "Podcast"
	TypenameProfile  Typename = "Profile"
	TypenameBook     Typename = "Book"
	TypenameChapter  Typename = "Chapter"
)

// Сортировка по дате.
type OrderDirection string

const (
	// По возрастанию (от меньшего к большему).
	OrderDirectionAsc OrderDirection = "asc"

	// По убыванию (от большего к меньшему).
	OrderDirectionDesc OrderDirection = "desc"
)

// Сортировка по полю.
type OrderBy string

const (
	// По алфавиту. Например по названию трека.
	OrderByAlphabet OrderBy = "alphabet"

	// По имени артиста.
	OrderByArtist OrderBy = "artist"

	// По дате добавления.
	OrderByDateAdded OrderBy = "dateAdded"
)

// Тип фона.
type BackgroundType string

const (
	// Картинка.
	BackgroundTypeImage BackgroundType = "image"
)

// Эффект анимации.
type AnimationEffect string

const (
	// Вправо.
	AnimationEffectToRight AnimationEffect = "toRight"
)
