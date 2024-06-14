package svc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"minizhihu/application/user/rpc/internal/config"
	"minizhihu/application/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModelExtra
	SMSClient *dysmsapi.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := dysmsapi.NewClientWithAccessKey(c.SMSClient.ConnAddress, c.SMSClient.AccessKeyID, c.SMSClient.AccessKeySecret)
	if err != nil {
		logx.Errorf("Failed to create SMS client: %v", err)
		return nil
	}
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModelExtra(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		SMSClient: client,
	}
}
