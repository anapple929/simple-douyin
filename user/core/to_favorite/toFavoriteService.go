package to_favorite

import (
	"context"
	"errors"
	"user/model"
	proto "user/services/to_favorite"
)

type ToFavoriteService struct {
}

/**
给Favorite微服务调用，更新用户表的total_favorited。
req携带的参数：userId 用户id   count 增加或者减少的数字   type 1增加2减少
*/
func (ToFavoriteService) UpdateTotalFavorited(ctx context.Context, req *proto.UpdateTotalFavoritedRequest, resp *proto.UpdateTotalFavoritedResponse) error {
	if req.UserId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的userId或者type有误")
	}
	//查一下，这个userId能否查到，查不到报错，查到了返回count
	if _, err := model.NewUserDaoInstance().FindUserById(req.UserId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewUserDaoInstance().AddTotalFavorited(req.UserId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewUserDaoInstance().ReduceTotalFavorited(req.UserId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}

/**
给Favorite微服务调用，更新用户表的favorite_count。
req携带的参数：userId 用户id   count 增加或者减少的数字   type 1增加2减少
*/
func (ToFavoriteService) UpdateFavoriteCount(ctx context.Context, req *proto.UpdateFavoriteCountRequest, resp *proto.UpdateFavoriteCountResponse) error {
	if req.UserId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的userId或者type有误")
	}
	//查一下，这个userId能否查到，查不到报错，查到了返回count
	if _, err := model.NewUserDaoInstance().FindUserById(req.UserId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewUserDaoInstance().AddFavoriteCount(req.UserId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewUserDaoInstance().ReduceFavoriteCount(req.UserId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}
