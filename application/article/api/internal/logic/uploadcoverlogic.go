package logic

import (
	"context"

	"minizhihu/application/article/api/internal/svc"
	"minizhihu/application/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover() (resp *types.UploadCoverResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
