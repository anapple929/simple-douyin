package mapper

import (
	"favorite/model"
	"fmt"
	"sync"
)

type FavoriteMapper struct {
}

func (m FavoriteMapper) GetVideoIds(id int64) []int64 {
	db := model.DB
	//db, _ := gorm.Open("mysql", "root:jhr292023@tcp(43.138.51.56:3306)/simple-douyin?charset=utf8&parseTime=True&loc=Local")

	var favs []*model.Favorite
	var videoIds []int64
	db.Select("video_id").Where("user_id=?", id).Find(&favs)

	for i := 0; i < len(favs); i++ {
		videoIds = append(videoIds, favs[i].VideoId)
	}
	return videoIds

}

func (m FavoriteMapper) GetFavoriteStatus(vid int64, uid int64) (bool, error) {
	db := model.DB
	var count int
	err := db.Model(&model.Favorite{}).Where("user_id=? and video_id=?", uid, vid).Count(&count).Error
	if err != nil {
		fmt.Println("数据库ERR：", err)
		return false, err
	}

	return count == 1, nil
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
