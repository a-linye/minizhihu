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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBeyondUserUserIdPrefix     = "cache:beyondUser:user:id:"
	cacheBeyondUserUserMobilePrefix = "cache:beyondUser:user:mobile:"
)

type (
	UserModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64     `db:"id"`          // 主键ID
		Username   string    `db:"username"`    // 用户名
		Avatar     string    `db:"avatar"`      // 头像
		Mobile     string    `db:"mobile"`      // 手机号
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 最后修改时间
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	beyondUserUserIdKey := fmt.Sprintf("%s%v", cacheBeyondUserUserIdPrefix, data.Id)
	beyondUserUserMobileKey := fmt.Sprintf("%s%v", cacheBeyondUserUserMobilePrefix, data.Mobile)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Avatar, data.Mobile)
	}, beyondUserUserIdKey, beyondUserUserMobileKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	beyondUserUserIdKey := fmt.Sprintf("%s%v", cacheBeyondUserUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, beyondUserUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	beyondUserUserMobileKey := fmt.Sprintf("%s%v", cacheBeyondUserUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, beyondUserUserMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows:
		return nil, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound

	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	beyondUserUserIdKey := fmt.Sprintf("%s%v", cacheBeyondUserUserIdPrefix, data.Id)
	beyondUserUserMobileKey := fmt.Sprintf("%s%v", cacheBeyondUserUserMobilePrefix, data.Mobile)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Username, data.Avatar, data.Mobile, data.Id)
	}, beyondUserUserIdKey, beyondUserUserMobileKey)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	beyondUserUserIdKey := fmt.Sprintf("%s%v", cacheBeyondUserUserIdPrefix, id)
	beyondUserUserMobileKey := fmt.Sprintf("%s%v", cacheBeyondUserUserMobilePrefix, data.Mobile)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, beyondUserUserIdKey, beyondUserUserMobileKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBeyondUserUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
