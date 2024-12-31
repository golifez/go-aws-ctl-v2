// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IS3Op interface {
		// 获取带region的客户端
		GetS3clinetWithRegion(region string) *sS3op
		// 获取存储桶列表
		GetS3BucketList(ctx context.Context) ([]string, error)
		// 获取S3桶对象列表
		GetS3BucketObjectList(ctx context.Context, bucketName string, region string) ([]string, error)
		// 创建S3桶
		CreateS3Bucket(ctx context.Context, bucketName string) error
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
