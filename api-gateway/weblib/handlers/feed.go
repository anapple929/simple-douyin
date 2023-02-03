package handlers

import (
	"api-gateway/services/feed"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(ginCtx *gin.Context) {
	var feedReq feed.DouyinFeedRequest
	lastTime, _ := strconv.ParseInt(ginCtx.Query("latest_time"), 10, 64)
	feedReq.LatestTime = lastTime
	feedReq.Token = ginCtx.Query("token")

	//设置最长响应时长
	ctx, _ := context.WithTimeout(ginCtx, time.Minute*1)

	// 从gin.Key中取出服务实例
	feedService := ginCtx.Keys["feedService"].(feed.FeedService)
	feedResp, err := feedService.Feed(ctx, &feedReq)
	if err != nil {
		fmt.Println(err)
	}

	ginCtx.JSON(http.StatusOK, feed.DouyinFeedResponse{
		StatusCode: feedResp.StatusCode,
		StatusMsg:  feedResp.StatusMsg,
		VideoList:  feedResp.VideoList,
		NextTime:   time.Now().Unix(),
	})
}
