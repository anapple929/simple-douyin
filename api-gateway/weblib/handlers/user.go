package handlers

import (
	"api-gateway/pkg/utils"
	user "api-gateway/services/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 用户注册
func Register(ginCtx *gin.Context) {
	var userReq user.DouyinUserRegisterRequest
	//PanicIfUserError(ginCtx.Bind(&userReq))
	userReq.Username = ginCtx.Query("username")
	userReq.Password = ginCtx.Query("password")

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, _ := userService.Register(context.Background(), &userReq)
	//PanicIfUserError(err)
	var token string
	if userResp.UserId > 0 {
		token, _ = utils.GenerateToken(userResp.UserId)
	}

	ginCtx.JSON(http.StatusOK, user.DouyinUserRegisterResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		UserId:     userResp.UserId,
		Token:      token,
	})
}

// 用户登录
func Login(ginCtx *gin.Context) {
	var userReq user.DouyinUserLoginRequest
	//PanicIfUserError(ginCtx.Bind(&userReq))
	userReq.Username = ginCtx.Query("username")
	userReq.Password = ginCtx.Query("password")

	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, _ := userService.Login(context.Background(), &userReq)
	//PanicIfUserError(err)
	var token string
	if userResp.UserId > 0 {
		token, _ = utils.GenerateToken(userResp.UserId)
	}

	fmt.Println("登录的token是:" + token)
	ginCtx.JSON(http.StatusOK, user.DouyinUserLoginResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		UserId:     userResp.UserId,
		Token:      token,
	})
}

func UserInfo(ginCtx *gin.Context) {
	var userReq user.DouyinUserRequest

	user_id, err := strconv.ParseInt(ginCtx.Query("user_id"), 10, 64)
	PanicIfUserError(err)

	userReq.UserId = user_id

	userReq.Token = ginCtx.Query("token")
	//claim, err := utils.ParseToken(userReq.Token)
	//if err != nil {
	//	ginCtx.JSON(http.StatusOK, publish.DouyinPublishListResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "token失效，请重新登录",
	//	})
	//}
	//currentUserId := strconv.FormatInt(claim.Id, 10)
	//userReq.Token = currentUserId

	userService := ginCtx.Keys["userService"].(user.UserService)
	userResp, err := userService.UserInfo(context.Background(), &userReq)
	//PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, user.DouyinUserResponse{
		StatusCode: userResp.StatusCode,
		StatusMsg:  userResp.StatusMsg,
		User:       userResp.User,
	})
}
