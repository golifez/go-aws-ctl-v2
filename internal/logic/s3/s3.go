package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golifez/go-aws-ctl-v2/internal/cmd"
	"github.com/golifez/go-aws-ctl-v2/internal/service"
)

type sS3op struct {
	s3cl *s3.Client
}

func init() {
	service.RegisterS3Op(NewsS3op())
}
func NewsS3op() *sS3op {
	clinet := cmd.GetNewAwsClinet("us-east-1")
	return &sS3op{
		s3cl: clinet.S3Client,
	}
}

// 获取带region的客户端
func (s *sS3op) GetS3clinetWithRegion(region string) *sS3op {
	clinet := cmd.GetNewAwsClinet(region)
	return &sS3op{
		s3cl: clinet.S3Client,
	}
}

// 获取存储桶列表
func (s *sS3op) GetS3BucketList(ctx context.Context) ([]string, error) {
	result, err := s.s3cl.ListBuckets(ctx, &s3.ListBucketsInput{})
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
func (s *sS3op) GetS3BucketObjectList(ctx context.Context, bucketName string, region string) ([]string, error) {
	// 通过区域获取S3客户端
	sl := s.GetS3clinetWithRegion(region)
	result, err := sl.s3cl.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
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
func (s *sS3op) CreateS3Bucket(ctx context.Context, bucketName string) error {
	_, err := s.s3cl.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}
	fmt.Println("已成功创建S3存储桶:", bucketName)
	return nil
}
