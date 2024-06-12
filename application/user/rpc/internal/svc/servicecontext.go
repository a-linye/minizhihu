package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"minizhihu/application/user/rpc/internal/config"
	"minizhihu/application/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModelExtra
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModelExtra(sqlx.NewMysql(c.DataSource), c.CacheRedis),
	}
}
