// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	IS3Op interface {
		// 获取存储桶列表
		GetS3BucketList(ctx context.Context, s3clinet *s3.Client) ([]string, error)
		// 获取S3桶对象列表
		GetS3BucketObjectList(ctx context.Context, s3clinet *s3.Client) ([]string, error)
		// 创建S3桶
		CreateS3Bucket(ctx context.Context, s3clinet *s3.Client) error
	}
)

var (
	localS3Op IS3Op
)

func S3Op() IS3Op {
	if localS3Op == nil {
		panic("implement not found for interface IS3Op, forgot register?")
	}
	return localS3Op
}

func RegisterS3Op(i IS3Op) {
	localS3Op = i
}
