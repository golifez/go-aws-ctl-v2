package s3

import (
	"context"
	"testing"

	v1 "github.com/golifez/go-aws-ctl-v2/api/s3/v1"
	_ "github.com/golifez/go-aws-ctl-v2/src/logic"
)

func TestListS3(t *testing.T) {
	NewV1().S3List(context.Background(), &v1.S3ListReq{})
}
