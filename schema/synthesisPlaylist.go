package schema

import _ "embed"

var (
	//go:embed gql/synthesisPlaylistBuild.gql
	_synthesisPlaylistBuild string
	//go:embed gql/synthesisPlaylist.gql
	_synthesisPlaylist string
)

func SynthesisPlaylistBuild(firstAuthor, secondAuthor ID) (string, error) {
	return getGraphqlBody(_synthesisPlaylistBuild, "synthesisPlaylistBuild", map[string]any{
		"firstAuthorId":  firstAuthor,
		"secondAuthorId": secondAuthor,
	})
}

type SynthesisPlaylistBuildResponse struct {
	SynthesisPlaylistBuild SynthesisPlaylist `json:"synthesisPlaylistBuild"`
}

func SynthesisPlaylistQ(ids []ID) (string, error) {
	return getGraphqlBody(_synthesisPlaylist, "synthesisPlaylist", map[string]any{
		"ids": ids,
	})
}

type SynthesisPlaylistResponse struct {
	SynthesisPlaylist []SynthesisPlaylist `json:"synthesisPlaylist"`
}

type (
	SynthesisPlaylist struct {
		ID      ID       `json:"id"`
		Tracks  []Track  `json:"tracks"`
		Authors []Author `json:"authors"`
	}
)
