package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/conf"
	"user/core"
	to_favorite "user/core/to_favorite"
	to_publish "user/core/to_publish"
	to_relation "user/core/to_relation"
	"user/services"
	to_favorite_proto "user/services/to_favorite"
	to_publish_proto "user/services/to_publish"
	to_relation_proto "user/services/to_relation"
	"user/utils/redis"
)

func main() {
	conf.Init()
	// 初始化redis-DB0的连接，follow选择的DB0.
	redis.InitRedis()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd注册件
		micro.Metadata(map[string]string{"protocol": "http"}),
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	_ = to_relation_proto.RegisterToRelationServiceHandler(microService.Server(), new(to_relation.ToRelationService))
	_ = to_publish_proto.RegisterToPublishServiceHandler(microService.Server(), new(to_publish.ToPublishService))
	_ = to_favorite_proto.RegisterToFavoriteServiceHandler(microService.Server(), new(to_favorite.ToFavoriteService))
	// 启动微服务
	_ = microService.Run()
}
