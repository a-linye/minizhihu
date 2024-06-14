package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ArticleModelExtra interface {
	ArticleModel
}
type customArticleModel struct {
	defaultArticleModel
}

func NewArticleModelExtra(conn sqlx.SqlConn, c cache.CacheConf) ArticleModelExtra {
	return &customArticleModel{
		defaultArticleModel: defaultArticleModel{
			CachedConn: sqlc.NewConn(conn, c),
			table:      "`article`",
		},
	}
}
