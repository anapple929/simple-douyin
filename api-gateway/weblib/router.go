package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service map[string]interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.InitMiddleware(service))
	//store := cookie.NewStore([]byte("something-very-secret"))
	//ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/douyin")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		//v1.GET("/user/", handlers.UserInfo)
		v1.POST("/user/register/", handlers.Register)
		v1.POST("/user/login/", handlers.Login)
		v1.GET("/feed/", handlers.Feed)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			v1.GET("user/", handlers.UserInfo)
			v1.POST("publish/action/", handlers.Publish)
			v1.GET("publish/list/", handlers.PublishList)
			/*authed.POST("task", handlers.CreateTask)
			authed.GET("task/:id", handlers.GetTaskDetail) // task_id
			authed.PUT("task/:id", handlers.UpdateTask)    // task_id
			authed.DELETE("task/:id", handlers.DeleteTask) // task_id*/
		}
	}
	return ginRouter
}
