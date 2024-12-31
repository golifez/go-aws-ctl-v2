package s3

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	S3QueryCmd = &cobra.Command{
		Use:   "s3qy",
		Short: "A brief description of your command",
		Long:  `AWS S3 查询相关命令1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("s3ctl called")
		},
	}
	//操作
	S3OpCmd = &cobra.Command{
		Use:   "s3op",
		Short: "A brief description of your command",
		Long:  `AWS S3 操作相关的命令`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("s3ctl called")
		},
	}
)
