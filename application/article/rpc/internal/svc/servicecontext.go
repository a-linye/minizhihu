package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
	"minizhihu/application/article/rpc/internal/config"
	"minizhihu/application/article/rpc/internal/model"
)

type ServiceContext struct {
	Config            config.Config
	ArticleModelExtra model.ArticleModelExtra
	BizRedis          *redis.Redis
	SingleFlightGroup singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:            c,
		ArticleModelExtra: model.NewArticleModelExtra(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		BizRedis:          rds,
	}
}
