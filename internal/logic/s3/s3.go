package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awsStc "github.com/golifez/go-aws-ctl-v2/internal/model/aws_stc"
	"github.com/golifez/go-aws-ctl-v2/internal/service"
)

type sS3op struct {
}

func init() {
	service.RegisterS3Op(NewsS3op())
}
func NewsS3op() *sS3op {
	return &sS3op{}
}

// 获取存储桶列表
func (s *sS3op) GetS3BucketList(ctx context.Context, input *awsStc.S3opInput) ([]string, error) {
	if input.S3Client == nil {
		fmt.Println("客户端为空")
		return nil, nil
	}
	fmt.Println("正在获取列表..")
	result, err := input.S3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	bucketNames := make([]string, len(result.Buckets))
	for i, bucket := range result.Buckets {
		bucketNames[i] = aws.ToString(bucket.Name)
	}
	return bucketNames, nil
}

// 获取S3桶对象列表
func (s *sS3op) GetS3BucketObjectList(ctx context.Context, input *awsStc.S3opInput) ([]string, error) {
	result, err := input.S3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(input.BucketName),
	})
	if err != nil {
		return nil, err
	}

	objectNames := make([]string, len(result.Contents))
	for i, obj := range result.Contents {
		objectNames[i] = *obj.Key
	}
	return objectNames, nil
}

// 创建S3桶
func (s *sS3op) CreateS3Bucket(ctx context.Context, input *awsStc.S3opInput) error {
	_, err := input.S3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(input.BucketName),
	})
	if err != nil {
		return err
	}
	fmt.Println("已成功创建S3存储桶:", input.BucketName)
	return nil
}
