// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package s3

import (
	"context"

	"github.com/golifez/go-aws-ctl-v2/api/s3/v1"
)

type IS3V1 interface {
	S3List(ctx context.Context, req *v1.S3ListReq) (res *v1.S3ListRes, err error)
}
