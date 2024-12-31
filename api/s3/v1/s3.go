package v1

import "github.com/aws/aws-sdk-go-v2/service/s3"

// 查询存储桶列表
type S3ListReq struct {
	S3cl *s3.Client
}
type S3ListRes struct {
}
