package main

import (
	"api-gateway/services/comment"
	"api-gateway/services/fav"
	"api-gateway/services/feed"
	message "api-gateway/services/message"
	"api-gateway/services/publish"
	"api-gateway/services/relation"
	"api-gateway/services/user"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	"time"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	userService := user.NewUserService("rpcUserService", userMicroService.Client())

	// publish
	publishMicroService := micro.NewService(
		micro.Name("publishService.client"),
		micro.WrapClient(wrappers.NewPublishWrapper),
	)
	// publish服务调用实例
	publishService := publish.NewPublishService("rpcPublishService", publishMicroService.Client())
	//点赞
	favoriteMicroService := micro.NewService(
		micro.Name("favoriteMicroService.client"),
		micro.WrapClient(wrappers.NewFavoriteWrapper),
	)
	//点赞服务实例
	favoriteService := fav.NewFavoriteService("rpcFavoriteService", favoriteMicroService.Client())
	// feed视频流
	feedMicroService := micro.NewService(
		micro.Name("feedService.client"),
		micro.WrapClient(wrappers.NewFeedWrapper),
	)
	// 视频流服务调用实例
	feedService := feed.NewFeedService("rpcFeedService", feedMicroService.Client())

	// comment
	commentMicroService := micro.NewService(
		micro.Name("commentService.client"),
		micro.WrapClient(wrappers.NewCommentWrapper),
	)
	// comment
	commentService := comment.NewCommentService("rpcCommentService", commentMicroService.Client())

	// relation
	relationMicroService := micro.NewService(
		micro.Name("relationService.client"),
		micro.WrapClient(wrappers.NewRelationWrapper),
	)
	// relation
	relationService := relation.NewRelationService("rpcRelationService", relationMicroService.Client())

	//message
	messageMicroService := micro.NewService(
		micro.Name("messageService.client"),
		micro.WrapClient(wrappers.NewMessageWrapper),
	)
	//message
	messageService := message.NewMessageService("rpcMessageService", messageMicroService.Client())

	serviceMap := make(map[string]interface{})
	serviceMap["userService"] = userService
	serviceMap["publishService"] = publishService
	serviceMap["feedService"] = feedService
	serviceMap["favoriteService"] = favoriteService
	serviceMap["commentService"] = commentService
	serviceMap["relationService"] = relationService
	serviceMap["messageService"] = messageService

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		//将服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(serviceMap)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*300),
		web.RegisterInterval(time.Second*150),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
