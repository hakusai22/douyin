package main

import (
	"douyin/v2/user/app/rpc/internal/config"
	"douyin/v2/user/app/rpc/internal/server"
	"douyin/v2/user/app/rpc/internal/svc"
	"douyin/v2/user/app/rpc/user"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "/Users/yinpeng/GoWorkSpace/douyin/v2/user/app/rpc/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	err := ctx.Redis.InitRedis(&c.CacheConfig, ctx.Db) // 初始化 Redis 缓存
	if err != nil {
		return
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserRpcServer(grpcServer, server.NewUserRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
