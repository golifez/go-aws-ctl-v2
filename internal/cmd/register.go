package cmd

import (
	"github.com/golifez/go-aws-ctl-v2/internal/cmd/s3"
)

func init() {
	Inits3cmd()
}

// 初始化S3相关命令
func Inits3cmd() {
	RootCmd.AddCommand(s3.S3QueryCmd)
	RootCmd.AddCommand(s3.S3OpCmd)
}
