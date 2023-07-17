package schema

import _ "embed"

var (
	//go:embed gql/followingCount.gql
	_followingCount string
)

func FollowingCount(id ID) (string, error) {
	return getGraphqlBody(_followingCount, "followingCount", map[string]any{
		"id": id,
	})
}

type FollowingCountResponse struct {
	Follows struct {
		Followings struct {
			Count int `json:"count"`
		} `json:"followings"`
	} `json:"follows"`
}

type (
	// Публичный профиль пользователя.
	SimpleProfile struct {
		ID          ID     `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       *Image `json:"image"`
	}

	Profile struct {
		Result struct {
			// Не nil если пользователь не авторизован.
			IsAnonymous     *bool            `json:"is_anonymous"`
			AllowExplicit   *bool            `json:"allow_explicit"`
			Birthday        *int64           `json:"birthday"`
			Created         *int64           `json:"created"`
			Email           *string          `json:"email"`
			ExternalProfile *ExternalProfile `json:"external_profile"`
			Gender          *string          `json:"gender"`
			HideBirthday    *bool            `json:"hide_birthday"`
			HideGender      *bool            `json:"hide_gender"`
			ID              *int             `json:"id"`
			Image           *Image           `json:"image"`
			IsActive        *bool            `json:"is_active"`
			IsAgreement     *bool            `json:"is_agreement"`
			IsEditor        *bool            `json:"is_editor"`
			IsRegistered    *bool            `json:"is_registered"`
			Name            *string          `json:"name"`
			Phone           *string          `json:"phone"`
			// Unix ms?
			Registered *int64 `json:"registered"`
			// Токен для доступа к API.
			Token    string  `json:"token"`
			Username *string `json:"username"`
		} `json:"result"`
	}

	// Профиль вне сервиса?
	ExternalProfile struct {
		Birthday   int64  `json:"birthday"`
		Email      string `json:"email"`
		ExternalID string `json:"external_id"`
		FirstName  string `json:"first_name"`
		Gender     string `json:"gender"`
		LastName   string `json:"last_name"`
		MiddleName string `json:"middle_name"`
		Phone      string `json:"phone"`
		Type       string `json:"type"`
	}
)
