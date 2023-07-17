package schema

import _ "embed"

var (
	//go:embed gql/addItemToCollection.gql
	_addItemToCollection string
	//go:embed gql/removeItemFromCollection.gql
	_removeItemFromCollection string
	//go:embed gql/getAllHiddenCollection.gql
	_getAllHiddenCollection string
	//go:embed gql/getHiddenTracks.gql
	_getHiddenTracks string
	//go:embed gql/addItemToHidden.gql
	_addItemToHidden string
	//go:embed gql/removeItemFromHidden.gql
	_removeItemFromHidden string
	//go:embed gql/userCollection.gql
	_userCollection string
	//go:embed gql/userPlaylists.gql
	_userPlaylists string
	//go:embed gql/userTracks.gql
	_userTracks string
	//go:embed gql/profileFollowersCount.gql
	_profileFollowersCount string
	//go:embed gql/userPaginatedPodcasts.gql
	_userPaginatedPodcasts string
)

func GetAllHiddenCollection() (string, error) {
	return getGraphqlBody(_getAllHiddenCollection, "getAllHiddenCollection", nil)
}

type GetAllHiddenCollectionResponse struct {
	HiddenCollection HiddenCollection `json:"hidden_collection"`
}

func GetHiddenTracks() (string, error) {
	return getGraphqlBody(_getHiddenTracks, "getHiddenTracks", nil)
}

type GetHiddenTracksResponse struct {
	HiddenCollection struct {
		Tracks []CollectionItem `json:"tracks"`
	} `json:"hidden_collection"`
}

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
		Playlists []CollectionItem `json:"playlists"`
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
		Tracks []struct {
			ID                 ID             `json:"id"`
			CollectionItemData CollectionItem `json:"collectionItemData"`
		} `json:"tracks"`
	} `json:"collection"`
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
			Items []SimplePodcast `json:"items"`
			Page  struct {
				EndCursor   string `json:"endCursor"`
				HasNextPage bool   `json:"hasNextPage"`
			} `json:"page"`
		} `json:"podcasts"`
	} `json:"paginatedCollection"`
}

func AddItemToCollection(id ID, itype CollectionItemType) (string, error) {
	return getGraphqlBody(_addItemToCollection, "addItemToCollection", map[string]any{
		"id":   id,
		"type": itype,
	})
}

func RemoveItemFromCollection(id ID, itype CollectionItemType) (string, error) {
	return getGraphqlBody(_removeItemFromCollection, "removeItemFromCollection", map[string]any{
		"id":   id,
		"type": itype,
	})
}

func AddItemToHidden(id ID, itype CollectionItemType) (string, error) {
	return getGraphqlBody(_addItemToHidden, "addItemToHidden", map[string]any{
		"id":   id,
		"type": itype,
	})
}

func RemoveItemFromHidden(id ID, itype CollectionItemType) (string, error) {
	return getGraphqlBody(_removeItemFromHidden, "removeItemFromHidden", map[string]any{
		"id":   id,
		"type": itype,
	})
}

type (
	HiddenCollection struct {
		Tracks  []CollectionItem `json:"tracks"`
		Artists []CollectionItem `json:"artists"`
	}

	Collection struct {
		Artists            []CollectionItem `json:"artists"`
		Episodes           []CollectionItem `json:"episodes"`
		Podcasts           []CollectionItem `json:"podcasts"`
		Playlists          []CollectionItem `json:"playlists"`
		SynthesisPlaylists []CollectionItem `json:"synthesis_playlists"`
		Profiles           []CollectionItem `json:"profiles"`
		Releases           []CollectionItem `json:"releases"`
		Tracks             []CollectionItem `json:"tracks"`
	}

	CollectionItem struct {
		// При вызове методов коллекции скорее всего будет не nil.
		ID *ID `json:"id"`

		UserID *ID `json:"userId"`

		// Статус. Может быть nil, как и LastModified.
		ItemStatus *CollectionItemStatus `json:"itemStatus"`

		// Дата добавления в коллекцию.
		//
		// При вызове методов коллекции скорее всего будет не nil.
		LastModified *Time `json:"lastModified"`

		// Дата последнего изменения коллекции (?).
		//
		// При вызове методов коллекции скорее всего будет не nil.
		CollectionLastModified *Time `json:"collectionLastModified"`
	}
)

// Тип в коллекции.
type CollectionItemType string

const (
	// Артист.
	CollectionItemTypeArtist CollectionItemType = "artist"

	// Альбом, EP, etc.
	CollectionItemTypeRelease CollectionItemType = "release"

	// Трек.
	CollectionItemTypeTrack CollectionItemType = "track"

	// Подкаст.
	CollectionItemTypePodcast CollectionItemType = "podcast"

	// Эпизод подкаста.
	CollectionItemTypeEpisode CollectionItemType = "episode"
)

// Статус сущности в коллекции.
type CollectionItemStatus string

const (
	// Лайкнута.
	CollectionItemStatusLiked CollectionItemStatus = "liked"
)
