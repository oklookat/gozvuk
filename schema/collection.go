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
)

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
		Tracks []HiddenCollectionItem `json:"tracks"`
	} `json:"hidden_collection"`
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
		Tracks  []HiddenCollectionItem `json:"tracks"`
		Artists []HiddenCollectionItem `json:"artists"`
	}

	HiddenCollectionItem struct {
		ID                     ID    `json:"id"`
		CollectionLastModified *Time `json:"collectionLastModified"`
	}

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

	UserCollectionTrack struct {
		ID                 ID             `json:"id"`
		CollectionItemData CollectionItem `json:"collectionItemData"`
	}

	CollectionItem struct {
		// Статус. Может быть nil, как и LastModified.
		ItemStatus *CollectionItemStatus `json:"itemStatus"`

		// Дата добавления в коллекцию.
		LastModified *Time `json:"lastModified"`
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
