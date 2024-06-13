package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"minizhihu/application/applet/internal/config"
	"minizhihu/application/user/rpc/user"
	"minizhihu/package/interceptors"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	UserRPC  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 添加自定义拦截器
	// 将RPC服务返回的错误码，解析为XCode类型
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
		BizRedis: redis.MustNewRedis(c.BizRedis),
	}
}
