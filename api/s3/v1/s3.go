package v1

import "github.com/gogf/gf/v2/frame/g"

// 查询存储桶列表
type S3ListReq struct {
	g.Meta `path:"/s3/list" method:"get" tags:"s3" dc:"查询存储桶列表"`
}
type S3ListRes struct {
}
