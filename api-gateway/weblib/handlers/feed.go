package handlers

import (
	"api-gateway/services/feed"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//视频流
func Feed(ginCtx *gin.Context) {
	//数据绑定
	var feedReq feed.DouyinFeedRequest
	lastTime, err := strconv.ParseInt(ginCtx.Query("latest_time"), 10, 64)
	PanicIfFeedError(err)
	feedReq.LatestTime = lastTime
	feedReq.Token = ginCtx.Query("token")

	//设置最长响应时长
	ctx, _ := context.WithTimeout(ginCtx, time.Minute*1)

	// 从gin.Key中取出服务实例
	feedService := ginCtx.Keys["feedService"].(feed.FeedService)
	// 调用
	feedResp, err := feedService.Feed(ctx, &feedReq)
	PanicIfFeedError(err)
	//返回
	ginCtx.JSON(http.StatusOK, feed.DouyinFeedResponse{
		StatusCode: feedResp.StatusCode,
		StatusMsg:  feedResp.StatusMsg,
		VideoList:  feedResp.VideoList,
		NextTime:   time.Now().Unix(),
	})
}
