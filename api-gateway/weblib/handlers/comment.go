package handlers

import (
	"api-gateway/services/comment"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentAction(ginCtx *gin.Context) {
	var commentReq comment.DouyinCommentActionRequest
	//获取request的信息
	commentReq.CommentId, _ = strconv.ParseInt(ginCtx.Query("comment_id"), 10, 64)
	commentReq.CommentText = ginCtx.Query("comment_text")
	actionType, _ := strconv.Atoi(ginCtx.Query("action_type"))
	commentReq.ActionType = int32(actionType)
	commentReq.Token = ginCtx.Query("token")
	commentReq.VideoId, _ = strconv.ParseInt(ginCtx.Query("video_id"), 10, 64)
	//if actionType != 1 && actionType != 2 {
	//	ginCtx.JSON(http.StatusOK, comment.DouyinCommentActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "action_type有问题",
	//	})
	//}
	//if commentReq.CommentText == "" {
	//	ginCtx.JSON(http.StatusOK, comment.DouyinCommentActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "comment_text有问题",
	//	})
	//}
	//if commentReq.VideoId <= 0 {
	//	ginCtx.JSON(http.StatusOK, comment.DouyinCommentActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "video_id有问题",
	//	})
	//}
	//if commentReq.ActionType == 2 && commentReq.CommentId <= 0 {
	//	ginCtx.JSON(http.StatusOK, comment.DouyinCommentActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "删除评论必须携带正确的id",
	//	})
	//}

	// 从gin.Key中取出服务实例
	commentService := ginCtx.Keys["commentService"].(comment.CommentService)
	//调用comment微服务，将context的上下文传入
	commentResp, _ := commentService.CommentAction(context.Background(), &commentReq)

	//返回
	ginCtx.JSON(http.StatusOK, comment.DouyinCommentActionResponse{
		StatusCode: commentResp.StatusCode,
		StatusMsg:  commentResp.StatusMsg,
		Comment:    commentResp.Comment,
	})

}
func CommentList(ginCtx *gin.Context) {
	var commentReq comment.DouyinCommentListRequest

	commentReq.Token = ginCtx.Query("token")
	commentReq.VideoId, _ = strconv.ParseInt(ginCtx.Query("video_id"), 10, 64)

	//if commentReq.VideoId <= 0 {
	//	ginCtx.JSON(http.StatusOK, comment.DouyinCommentListResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "video_id有问题",
	//	})
	//}

	commentService := ginCtx.Keys["commentService"].(comment.CommentService)
	commentResp, _ := commentService.CommentList(context.Background(), &commentReq)

	ginCtx.JSON(http.StatusOK, comment.DouyinCommentListResponse{
		StatusCode:  commentResp.StatusCode,
		StatusMsg:   commentResp.StatusMsg,
		CommentList: commentResp.CommentList,
	})

}
