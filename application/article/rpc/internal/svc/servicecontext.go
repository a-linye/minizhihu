package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"minizhihu/application/article/rpc/internal/config"
	"minizhihu/application/article/rpc/internal/model"
)

type ServiceContext struct {
	Config            config.Config
	ArticleModelExtra model.ArticleModelExtra
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		ArticleModelExtra: model.NewArticleModelExtra(sqlx.NewMysql(c.DataSource), c.CacheRedis),
	}
}
