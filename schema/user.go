package schema

import _ "embed"

var (
	//go:embed gql/userCollection.gql
	_userCollection string
	//go:embed gql/userPlaylists.gql
	_userPlaylists string
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

type (
	Collection struct {
		Artists            []UserCollectionItem `json:"artists"`
		Episodes           []UserCollectionItem `json:"episodes"`
		Podcasts           []UserCollectionItem `json:"podcasts"`
		Playlists          []UserCollectionItem `json:"playlists"`
		SynthesisPlaylists []UserCollectionItem `json:"synthesis_playlists"`
		Profiles           []UserCollectionItem `json:"profiles"`
		Releases           []UserCollectionItem `json:"releases"`
		Tracks             []UserCollectionItem `json:"tracks"`
	}

	UserCollectionItem struct {
		ID           ID    `json:"id"`
		LastModified *Time `json:"last_modified"`
	}

	UserCollectionPlaylist struct {
		ID           ID    `json:"id"`
		UserID       ID    `json:"user_id"`
		LastModified *Time `json:"last_modified"`
	}
)
