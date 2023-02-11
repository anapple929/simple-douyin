package rpc_server

import (
	"fmt"
	services "publish/services/favorite_to_video_proto"
	"testing"
)

func TestGetFavoritesStatus(t *testing.T) {
	var isFavorites []*services.FavoriteStatus
	isFavorites = append(isFavorites, &services.FavoriteStatus{UserId: 51, VideoId: 28, IsFavorite: false})
	isFavorites = append(isFavorites, &services.FavoriteStatus{UserId: 54, VideoId: 39, IsFavorite: false})
	isFavorites = append(isFavorites, &services.FavoriteStatus{UserId: 58, VideoId: 6, IsFavorite: false})
	//result, _ := GetUsersInfo(userId, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NTEsImV4cCI6MTY3NjExODIxNSwiaXNzIjoiMTEyMjIzMyJ9.WD4uT4aQciCyxq-3kEpIi8XGtUpEDhrx32H6MqkUn2Q")
	result, _ := GetFavoritesStatus(isFavorites)

	fmt.Println(result)
}
