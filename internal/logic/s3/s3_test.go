package s3

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golifez/go-aws-ctl-v2/internal/cmd"
	awsStc "github.com/golifez/go-aws-ctl-v2/internal/model/aws_stc"
	"github.com/golifez/go-aws-ctl-v2/internal/service"
)

var (
	ctx = context.Background()
)

func TestS3op(t *testing.T) {
	cl := GetS3clinet()
	r, err := service.S3Op().GetS3BucketList(ctx, &awsStc.S3opInput{
		S3Client: cl,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println("结果是", r)
}

func GetS3clinet(region ...string) *s3.Client {
	if len(region) == 0 {
		region = []string{"us-east-1"}
	}
	client := cmd.GetNewAwsClinetWithRegion(region[0])
	return client.S3Client
}

// 测试创建S3桶
//func TestCreateS3Bucket(t *testing.T) {
//	service.S3Op().CreateS3Bucket(ctx, s3.NewFromConfig(aws.Config{
//		Region: "us-east-1",
//	}))
//}
