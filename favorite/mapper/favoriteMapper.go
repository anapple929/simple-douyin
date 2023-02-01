package mapper

import "sync"

type FavoriteMapper struct {
}

var favoriteMapper *FavoriteMapper
var favoriteOnce sync.Once //单例模式，favoriteOnce，提高性能

func FavoriteMapperInstance() *FavoriteMapper {
	favoriteOnce.Do(
		func() {
			favoriteMapper = &FavoriteMapper{}
		})
	return favoriteMapper
}
