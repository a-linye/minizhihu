package logic

import (
	"context"

	"minizhihu/application/user/rpc/internal/svc"
	"minizhihu/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &service.RegisterResponse{}, nil
}
