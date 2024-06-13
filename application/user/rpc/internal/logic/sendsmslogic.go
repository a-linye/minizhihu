package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"minizhihu/application/user/rpc/internal/svc"
	"minizhihu/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *service.SendSmsRequest) (*service.SendSmsResponse, error) {

	// 构造请求
	request := dysmsapi.CreateSendSmsRequest()       //创建请求
	request.Scheme = "https"                         //请求协议
	request.PhoneNumbers = in.Mobile                 //接收短信的手机号码
	request.SignName = "教程"                          //短信签名名称
	request.TemplateCode = "SMS_******236"           //短信模板ID
	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": in.Code,
	})
	request.TemplateParam = string(par) //将短信模板参数传入短信模板

	// 发送短信
	response, err := l.svcCtx.SMSClient.SendSms(request)
	if err != nil {
		logx.Errorf("Failed to send SMS: %v", err)
		return nil, err
	}
	if response.Code != "OK" {
		logx.Errorf("Failed to send SMS, response code: %s, message: %s", response.Code, response.Message)
		return nil, fmt.Errorf("failed to send SMS, response code: %s, message: %s", response.Code, response.Message)
	}

	return &service.SendSmsResponse{}, nil
}
