package s3

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/golifez/go-aws-ctl-v2/api/s3/v1"
)

func (c *ControllerV1) S3List(ctx context.Context, req *v1.S3ListReq) (res *v1.S3ListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
