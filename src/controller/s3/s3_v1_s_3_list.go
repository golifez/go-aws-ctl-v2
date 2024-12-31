package s3

import (
	"context"
	"fmt"

	v1 "github.com/golifez/go-aws-ctl-v2/api/s3/v1"
	awsStc "github.com/golifez/go-aws-ctl-v2/src/model/aws_stc"
	"github.com/golifez/go-aws-ctl-v2/src/service"
)

// 查询存储桶列表
func (c *ControllerV1) S3List(ctx context.Context, req *v1.S3ListReq) (res *v1.S3ListRes, err error) {
	if req.S3cl != nil {
		fmt.Println("客户端正常")
	}
	result, err := service.S3Op().GetS3BucketList(ctx, &awsStc.S3opInput{
		S3Client: req.S3cl,
	})
	if err != nil {
		return nil, err
	}
	for _, bucketName := range result {
		fmt.Println("bucket:", bucketName)
	}
	return nil, nil
}
