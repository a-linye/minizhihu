package interceptors

import (
	"context"

	"minizhihu/package/xcode"

	"google.golang.org/grpc"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 处理RPC服务的业务
		resp, err = handler(ctx, req)
		// 处理调用RPC服务后的错误
		return resp, xcode.FromError(err).Err()
	}
}
