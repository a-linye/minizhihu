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

/*
go zero当指定使用缓存方式生成查询代码时。底层会先查缓存，如果缓存不存在则查数据库。并将查到的数据放到缓存中，设置过期时间为7天
如果查询不到数据，则为查询的key设置一个空缓存，防止缓存击穿
*/

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
