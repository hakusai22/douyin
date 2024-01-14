package main

import (
	"douyin/v2/jwt/app/rpc/Jwt"
	"douyin/v2/jwt/app/rpc/internal/config"
	"douyin/v2/jwt/app/rpc/internal/server"
	"douyin/v2/jwt/app/rpc/internal/svc"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "/Users/yinpeng/GoWorkSpace/douyin/v2/jwt/app/rpc/etc/jwt.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		Jwt.RegisterJwtRpcServer(grpcServer, server.NewJwtRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
