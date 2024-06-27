package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ArticleModelExtra interface {
	ArticleModel
	ArticlesByUserId(ctx context.Context, userId int64, status int, likeNum int64, pubTime, sortField string, limit int) ([]*Article, error)
	UpdateArticleStatus(ctx context.Context, id int64, status int) error
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

func (m *customArticleModel) ArticlesByUserId(ctx context.Context, userId int64, status int, likeNum int64, pubTime, sortField string, limit int) ([]*Article, error) {
	var (
		err      error
		sql      string
		anyField any
		articles []*Article
	)
	if sortField == "like_num" {
		anyField = likeNum
		sql = fmt.Sprintf("select "+articleRows+" from "+m.table+" where author_id=? and status=? and like_num < ? order by %s desc limit ?", sortField)
	} else {
		anyField = pubTime
		sql = fmt.Sprintf("select "+articleRows+" from "+m.table+" where author_id=? and status=? and publish_time < ? order by %s desc limit ?", sortField)
	}
	err = m.QueryRowsNoCacheCtx(ctx, &articles, sql, userId, status, anyField, limit)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *customArticleModel) UpdateArticleStatus(ctx context.Context, id int64, status int) error {
	beyondArticleArticleIdKey := fmt.Sprintf("%s%v", cacheBeyondArticleArticleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set status = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, status, id)
	}, beyondArticleArticleIdKey)

	return err
}
