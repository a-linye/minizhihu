package main

import (
	"flag"
	"fmt"
	"minizhihu/package/interceptors"

	"minizhihu/application/user/rpc/internal/config"
	"minizhihu/application/user/rpc/internal/server"
	"minizhihu/application/user/rpc/internal/svc"
	"minizhihu/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	cs "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterUserServer(grpcServer, svr)

		if c.Mode == cs.DevMode || c.Mode == cs.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())
	defer s.Stop()
	// 自定义拦截器

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
