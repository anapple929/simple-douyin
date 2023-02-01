package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services/publish"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func Publish(ginCtx *gin.Context) {
	var publishReq publish.DouyinPublishActionRequest

	//publishReq.Data = []byte(ginCtx.PostForm("data"))
	publishReq.Title = ginCtx.PostForm("title")
	publishReq.Token = ginCtx.PostForm("token")
	fileHeader, _ := ginCtx.FormFile("data")

	file, err := fileHeader.Open()
	if err != nil {
		//文件读取失败
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		//SendResponse(c, pack.BuildPublishResp(err))
		//转换异常
		return
	}

	publishReq.Data = buf.Bytes()
	//fmt.Println("绑定的数据")
	//fmt.Println(publishReq.Data)
	//fmt.Println(publishReq.Title)
	//fmt.Println(publishReq.Token)

	//token中的userId提取出来
	claim, _ := utils.ParseToken(publishReq.Token)
	currentPublishId := strconv.FormatInt(claim.Id, 10)
	publishReq.Token = currentPublishId

	// 从gin.Key中取出服务实例
	publishService := ginCtx.Keys["publishService"].(publish.PublishService)
	publishResp, _ := publishService.Publish(context.Background(), &publishReq)

	ginCtx.JSON(http.StatusOK, publish.DouyinPublishActionResponse{
		StatusCode: publishResp.StatusCode,
		StatusMsg:  publishResp.StatusMsg,
	})
}

func PublishList(ginCtx *gin.Context) {
	var publishReq publish.DouyinPublishListRequest
	// 从gin.Key中取出服务实例
	publishService := ginCtx.Keys["publishService"].(publish.PublishService)
	publishResp, _ := publishService.PublishList(context.Background(), &publishReq)

	ginCtx.JSON(http.StatusOK, publish.DouyinPublishListResponse{
		StatusCode: publishResp.StatusCode,
		StatusMsg:  publishResp.StatusMsg,
	})
}
