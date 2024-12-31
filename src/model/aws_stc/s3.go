package aws_stc

import "github.com/aws/aws-sdk-go-v2/service/s3"

type S3opInput struct {
	S3Client   *s3.Client
	Region     string
	BucketName string
}
