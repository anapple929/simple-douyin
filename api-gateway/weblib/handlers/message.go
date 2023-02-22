package handlers

import (
	message "api-gateway/services/message"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
聊天记录
*/
func MessageList(ginCtx *gin.Context) {
	var messageReq message.DouyinMessageChatRequest

	messageReq.Token = ginCtx.Query("token")
	messageReq.ToUserId, _ = strconv.ParseInt(ginCtx.Query("to_user_id"), 10, 64)
	messageReq.PreMsgTime, _ = strconv.ParseInt(ginCtx.Query("pre_msg_time"), 10, 64)
	//if messageReq.ToUserId <= 0 {
	//	ginCtx.JSON(http.StatusOK, message.DouyinMessageChatResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "to_user_id有问题",
	//	})
	//}
	messageService := ginCtx.Keys["messageService"].(message.MessageService)
	resp, _ := messageService.MessageList(context.Background(), &messageReq)
	//for _, message := range resp.MessageList {
	//	fmt.Println(message)
	//}
	ginCtx.JSON(http.StatusOK, message.DouyinMessageChatResponse{
		StatusCode:  resp.StatusCode,
		StatusMsg:   resp.StatusMsg,
		MessageList: resp.MessageList,
	})
}

/**
发送消息
*/
func MessageAction(ginCtx *gin.Context) {
	var messageReq message.DouyinMessageActionRequest

	messageReq.Token = ginCtx.Query("token")
	messageReq.ToUserId, _ = strconv.ParseInt(ginCtx.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.Atoi(ginCtx.Query("action_type"))
	messageReq.ActionType = int32(actionType)
	messageReq.Content = ginCtx.Query("content")

	messageService := ginCtx.Keys["messageService"].(message.MessageService)
	resp, _ := messageService.MessageAction(context.Background(), &messageReq)

	ginCtx.JSON(http.StatusOK, message.DouyinMessageActionResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	})
}
