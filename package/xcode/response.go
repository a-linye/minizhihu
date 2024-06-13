package xcode

import (
	"net/http"

	"minizhihu/package/xcode/types"
)

func ErrHandler(err error) (int, any) {
	// 核心关键方法
	// 判断传入的错误能否转换成自定义的错误类型
	code := CodeFromError(err)

	// 写入header部分的code永远是200
	// 返回给前端的错误码永远是从common.xcode中找
	// 当匹配不到时默认返回ServerErr
	// ServerErr          = add(500, "INTERNAL_ERROR")
	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
