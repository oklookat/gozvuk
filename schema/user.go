package schema

import _ "embed"

var (
	//go:embed gql/userCollection.gql
	_userCollection string
	//go:embed gql/userPlaylists.gql
	_userPlaylists string
	//go:embed gql/userTracks.gql
	_userTracks string
	//go:embed gql/followingCount.gql
	_followingCount string
	//go:embed gql/profileFollowersCount.gql
	_profileFollowersCount string
	//go:embed gql/userPaginatedPodcasts.gql
	_userPaginatedPodcasts string
)

func UserColelctionQ() (string, error) {
	return getGraphqlBody(_userCollection, "userCollection", nil)
}

type UserCollectionResponse struct {
	Collection Collection `json:"collection"`
}

func UserPlaylists() (string, error) {
	return getGraphqlBody(_userPlaylists, "userPlaylists", nil)
}

type UserPlaylistsResponse struct {
	Collection struct {
		Playlists []UserCollectionPlaylist `json:"playlists"`
	} `json:"collection"`
}

func UserTracks(direction OrderDirection, by OrderBy) (string, error) {
	return getGraphqlBody(_userTracks, "userTracks", map[string]any{
		"orderDirection": direction,
		"orderBy":        by,
	})
}

type UserTracksResponse struct {
	Collection struct {
		Tracks []UserCollectionTrack `json:"tracks"`
	} `json:"collection"`
}

func FollowingCount(id ID) (string, error) {
	return getGraphqlBody(_followingCount, "followingCount", map[string]any{
		"id": id,
	})
}

type FollowingCountResponse struct {
	Follows struct {
		Followings Followings `json:"followings"`
	} `json:"follows"`
}

func ProfileFollowersCount(ids []ID) (string, error) {
	return getGraphqlBody(_profileFollowersCount, "profileFollowersCount", map[string]any{
		"ids": ids,
	})
}

type ProfileFollowersCountResponse struct {
	Profiles []struct {
		CollectionItemData struct {
			LikesCount int `json:"likesCount"`
		} `json:"collectionItemData"`
	} `json:"profiles"`
}

func UserPaginatedPodcasts(cursor string, count int) (string, error) {
	return getGraphqlBody(_userPaginatedPodcasts, "userPaginatedPodcasts", map[string]any{
		"cursor": cursor,
		"count":  count,
	})
}

type UserPaginatedPodcastsResponse struct {
	PaginatedCollection struct {
		Podcasts struct {
			Items []Podcast `json:"items"`
			Page  struct {
				EndCursor   string `json:"endCursor"`
				HasNextPage bool   `json:"hasNextPage"`
			} `json:"page"`
		} `json:"podcasts"`
	} `json:"paginatedCollection"`
}

type (
	Followings struct {
		Count int `json:"count"`
	}

	// Публичный профиль пользователя.
	PublicProfile struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
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
