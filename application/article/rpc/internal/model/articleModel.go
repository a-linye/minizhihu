package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleFieldNames          = builder.RawFieldNames(&Article{})
	articleRows                = strings.Join(articleFieldNames, ",")
	articleRowsExpectAutoSet   = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	articleRowsWithPlaceHolder = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBeyondArticleArticleIdPrefix = "cache:beyondArticle:article:id:"
)

type (
	ArticleModel interface {
		Insert(ctx context.Context, data *Article) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Article, error)
		Update(ctx context.Context, data *Article) error
		Delete(ctx context.Context, id int64) error
	}

	defaultArticleModel struct {
		sqlc.CachedConn
		table string
	}

	Article struct {
		Id          int64     `db:"id"`           // 主键ID
		Title       string    `db:"title"`        // 标题
		Content     string    `db:"content"`      // 内容
		Cover       string    `db:"cover"`        // 封面
		Description string    `db:"description"`  // 描述
		AuthorId    int64     `db:"author_id"`    // 作者ID
		Status      int64     `db:"status"`       // 状态 0:待审核 1:审核不通过 2:可见 3:用户删除
		CommentNum  int64     `db:"comment_num"`  // 评论数
		LikeNum     int64     `db:"like_num"`     // 点赞数
		CollectNum  int64     `db:"collect_num"`  // 收藏数
		ViewNum     int64     `db:"view_num"`     // 浏览数
		ShareNum    int64     `db:"share_num"`    // 分享数
		TagIds      string    `db:"tag_ids"`      // 标签ID
		PublishTime time.Time `db:"publish_time"` // 发布时间
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdateTime  time.Time `db:"update_time"`  // 最后修改时间
	}
)

func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleModel {
	return &defaultArticleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`article`",
	}
}

func (m *defaultArticleModel) Insert(ctx context.Context, data *Article) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, articleRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.Title, data.Content, data.Cover, data.Description, data.AuthorId, data.Status, data.CommentNum, data.LikeNum, data.CollectNum, data.ViewNum, data.ShareNum, data.TagIds, data.PublishTime)

	return ret, err
}

func (m *defaultArticleModel) FindOne(ctx context.Context, id int64) (*Article, error) {
	beyondArticleArticleIdKey := fmt.Sprintf("%s%v", cacheBeyondArticleArticleIdPrefix, id)
	var resp Article
	err := m.QueryRowCtx(ctx, &resp, beyondArticleArticleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleModel) Update(ctx context.Context, data *Article) error {
	beyondArticleArticleIdKey := fmt.Sprintf("%s%v", cacheBeyondArticleArticleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Content, data.Cover, data.Description, data.AuthorId, data.Status, data.CommentNum, data.LikeNum, data.CollectNum, data.ViewNum, data.ShareNum, data.TagIds, data.PublishTime, data.Id)
	}, beyondArticleArticleIdKey)
	return err
}

func (m *defaultArticleModel) Delete(ctx context.Context, id int64) error {
	beyondArticleArticleIdKey := fmt.Sprintf("%s%v", cacheBeyondArticleArticleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, beyondArticleArticleIdKey)
	return err
}

func (m *defaultArticleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBeyondArticleArticleIdPrefix, primary)
}

func (m *defaultArticleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
