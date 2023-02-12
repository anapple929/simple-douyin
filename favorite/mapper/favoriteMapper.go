package mapper

import (
	"errors"
	etcdInit "favorite/etcd"
	"favorite/model"
	proto "favorite/service"
	redis "favorite/utils/redis"
	"fmt"
	"log"
	"strconv"
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

func (m FavoriteMapper) FavoriteAction(uid int64, vid int64, actionType int32) error {
	db := model.DB
	fav := &model.Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	if actionType == 1 {
		var count int
		db.Where("user_id=? and video_id=?", uid, vid).Count(&count)
		if count >= 1 {
			db.Where("user_id=? and video_id=?", uid, vid).Delete(fav)
		}

		err := db.Create(fav).Error
		if err != nil {
			fmt.Println("点赞失败")
			return err
		}

		//点赞成功了，点赞状态改变，删除缓存
		key := strconv.FormatInt(uid, 10) + "+" + strconv.FormatInt(vid, 10)
		countRedis, err := redis.RdbUserVideo.Exists(redis.Ctx, key).Result()
		if err != nil {
			log.Println(err)
		}
		if countRedis > 0 {
			redis.RdbUserVideo.Del(redis.Ctx, key)
		}

	} else if actionType == 2 {
		err := db.Where("user_id=? and video_id=?", uid, vid).Delete(fav).Error
		if err != nil {
			fmt.Println("点赞删除失败")
			return err
		}

		//取消点赞成功了，点赞状态改变，删除缓存
		key := strconv.FormatInt(uid, 10) + "+" + strconv.FormatInt(vid, 10)
		countRedis, err := redis.RdbUserVideo.Exists(redis.Ctx, key).Result()
		if err != nil {
			log.Println(err)
		}
		if countRedis > 0 {
			redis.RdbUserVideo.Del(redis.Ctx, key)
		}

	} else {
		return errors.New("参数错误")
	}
	action := etcdInit.CountAction(vid, 1, actionType)
	if !action {
		return errors.New("Mapper层Count维护失败")
	}
	return nil
}

func (m FavoriteMapper) GetFavoritesStatus(isFavorites []*proto.FavoriteStatus) ([]*proto.FavoriteStatus, error) {
	db := model.DB
	var result []*proto.FavoriteStatus
	var count int

	for _, isFavorite := range isFavorites {
		_ = db.Model(&model.Favorite{}).Where("user_id=? and video_id=?", isFavorite.UserId, isFavorite.VideoId).Count(&count).Error
		result = append(result, &proto.FavoriteStatus{IsFavorite: count > 0, UserId: isFavorite.UserId, VideoId: isFavorite.VideoId})
	}
	fmt.Println(result)
	return result, nil
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
