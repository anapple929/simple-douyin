package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 用户注册
func Register(ginCtx *gin.Context) {
	var userReq services.DouyinUserRegisterRequest
	//PanicIfUserError(ginCtx.Bind(&userReq))
	userReq.Username = ginCtx.Query("username")
	userReq.Password = ginCtx.Query("password")
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.Register(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(userResp.UserId)

	ginCtx.JSON(http.StatusOK, services.DouyinUserRegisterResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		UserId:     userResp.UserId,
		Token:      token,
	})
}

// 用户登录
func Login(ginCtx *gin.Context) {
	var userReq services.DouyinUserLoginRequest
	//PanicIfUserError(ginCtx.Bind(&userReq))
	userReq.Username = ginCtx.Query("username")
	userReq.Password = ginCtx.Query("password")
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.Login(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(userResp.UserId)

	ginCtx.JSON(http.StatusOK, services.DouyinUserLoginResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		UserId:     userResp.UserId,
		Token:      token,
	})
}

func UserInfo(ginCtx *gin.Context) {
	var userReq services.DouyinUserRequest
	//PanicIfUserError(ginCtx.Bind(&userReq))
	user_id, err := strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)
	PanicIfUserError(err)
	userReq.UserId = user_id

	userReq.Token = ginCtx.Query("token")
	// 从gin.Key中取出服务实例
	//claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	claim, err := utils.ParseToken(userReq.Token)
	PanicIfUserError(err)

	currentUserId := strconv.FormatInt(claim.Id, 10)
	userReq.Token = currentUserId
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserInfo(context.Background(), &userReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, services.DouyinUserResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		User:       userResp.User,
	})
}
