package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserModelExtra interface {
	UserModel // 继承UserModel接口
	FindByMobile(ctx context.Context, mobile string) (*User, error)
}

type customUserModel struct {
	defaultUserModel // 内嵌defaultUserModel
}

func NewUserModelExtra(conn sqlx.SqlConn, c cache.CacheConf) UserModelExtra {
	return &customUserModel{
		defaultUserModel: defaultUserModel{
			CachedConn: sqlc.NewConn(conn, c),
			table:      "`user`",
		},
	}
}

func (m *customUserModel) FindByMobile(ctx context.Context, mobile string) (*User, error) {
	var user User
	err := m.QueryRowNoCacheCtx(ctx, &user, fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table), mobile)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
