# gozvuk

Неофициальный клиент для API [zvuk.com](https://zvuk.com).

# Важное и не очень

## Несуществующие треки

1. Получаем айдишники лайкнутых треков: UserTracks (OrderByDateAdded, OrderDirectionAsc)
2. Получаем полные версии треков: GetFullTrack
3. Некоторые айдишники будут пустые. То есть в коллекции у нас якобы есть трек,
   но по факту трека не существует.

Из-за этого дальнейшие действия с треками могут привести к ошибке, потому что
вы можете передать пустой ID.

Мораль: проверяйте, не пустые ли ID у полученных треков (и не только треков).

# Реализовано

- [x] search

- [x] quickSearch

- [x] getArtists

- [x] getFullTrack

- [x] getTracks

- [x] getReleases

- [x] getPlaylists

- [x] getPlaylistTracks

- [x] getShortPlaylist

- [x] renamePlaylist

- [x] setPlaylistToPublic

- [x] userPlaylists

- [x] addTracksToPlaylist

- [x] updataPlaylist

- [x] userTracks

- [x] getAllHiddenCollection

- [x] getHiddenTracks

- [x] userCollection

- [x] userPaginatedPodcasts

- [x] synthesisPlaylist

- [x] synthesisPlaylistBuild

- [x] getEpisodes

- [x] getPodcasts

- [x] addItemToCollection

- [x] addItemToHidden

- [x] removeItemFromCollection

- [x] removeItemFromHidden

- [x] createPlaylist

- [x] deletePlaylist

- [x] followingCount

- [x] profileFollowersCount

- [ ] listenedEpisodes

- [ ] listeningHistory

- [ ] notificationsHasUnread
