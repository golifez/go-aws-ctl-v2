// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	awsStc "github.com/golifez/go-aws-ctl-v2/internal/model/aws_stc"
)

type (
	IS3Op interface {
		// 获取存储桶列表
		GetS3BucketList(ctx context.Context, input *awsStc.S3opInput) ([]string, error)
		// 获取S3桶对象列表
		GetS3BucketObjectList(ctx context.Context, input *awsStc.S3opInput) ([]string, error)
		// 创建S3桶
		CreateS3Bucket(ctx context.Context, input *awsStc.S3opInput) error
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
