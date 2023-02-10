package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(service map[string]interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.InitMiddleware(service))

	v1 := ginRouter.Group("/douyin")
	{
		//user
		user := v1.Group("/user")
		{
			user.POST("/register/", handlers.Register)
			user.POST("/login/", handlers.Login)
			user.GET("/", handlers.UserInfo)
		}

		//publish
		publish := v1.Group("/publish")
		{
			publish.POST("/action/", handlers.Publish)
			publish.GET("/list/", handlers.PublishList)
		}

		//feed
		feed := v1.Group("/feed")
		{
			feed.GET("/", handlers.Feed)
		}

		//favorite
		fav := v1.Group("/favorite")
		{
			fav.POST("/action/", handlers.FavoriteAction)
			fav.GET("/list/", handlers.FavoriteList)
		}

		//comment
		comment := v1.Group("/comment")
		{
			comment.POST("/action/", handlers.CommentAction)
			comment.GET("/list/", handlers.CommentList)
		}
	}
	return ginRouter
}
