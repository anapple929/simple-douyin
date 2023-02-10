package to_comment

import (
	"context"
	"errors"
	"publish/model"
	proto "publish/services/to_comment"
)

type ToCommentService struct {
}

/**
给Comment微服务调用，更新视频表的评论数。
req携带的参数：videoId 视频id   count 增加或者减少的数字   type 1增加2减少
*/
func (ToCommentService) UpdateCommentCount(ctx context.Context, req *proto.UpdateCommentCountRequest, resp *proto.UpdateCommentCountResponse) error {
	if req.VideoId <= 0 || (req.Type != 1 && req.Type != 2) {
		resp.StatusCode = -1
		return errors.New("传入的videoId或者type有误")
	}
	//查一下，这个videoId能否查到，查不到报错，查到了返回count
	if _, err := model.NewVideoDaoInstance().FindVideoById(req.VideoId); err != nil {
		return errors.New("传入的VideoId查不到")
	}
	//调用数据库的修改功能
	if req.Type == 1 {
		//增加
		model.NewVideoDaoInstance().AddCommentCount(req.VideoId, req.Count)
	} else if req.Type == 2 {
		//减少
		model.NewVideoDaoInstance().ReduceCommentCount(req.VideoId, req.Count)
	}

	resp.StatusCode = 0
	return nil
}
