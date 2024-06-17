package code

import "minizhihu/package/xcode"

var (
	GetBucketErr = xcode.New(30001, "获取bucket实例失败")
	PutBucketErr = xcode.New(30002, "上传bucket失败")
)
