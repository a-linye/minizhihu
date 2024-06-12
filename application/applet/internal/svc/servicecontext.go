package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"minizhihu/application/applet/internal/config"
	"minizhihu/application/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	UserRPC  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	//TODO: 添加自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
		BizRedis: redis.MustNewRedis(c.BizRedis),
	}
}
