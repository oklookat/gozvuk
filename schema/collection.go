package schema

import _ "embed"

var (
	//go:embed gql/addItemToCollection.gql
	_addItemToCollection string
	//go:embed gql/removeItemFromCollection.gql
	_removeItemFromCollection string
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
