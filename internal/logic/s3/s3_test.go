package s3

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golifez/go-aws-ctl-v2/internal/service"
)

var (
	ctx = context.Background()
)

func TestS3op(t *testing.T) {
	//service.S3Op().GetS3BucketList(ctx)
}

// 测试创建S3桶
func TestCreateS3Bucket(t *testing.T) {
	service.S3Op().CreateS3Bucket(ctx, s3.NewFromConfig(aws.Config{
		Region: "us-east-1",
	}))
}
