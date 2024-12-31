package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golifez/go-aws-ctl-v2/internal/cmd"
	"github.com/golifez/go-aws-ctl-v2/internal/service"
	"github.com/spf13/cobra"
)

// lgCmd represents the lg command
var S3Cmd = &cobra.Command{
	Use:   "s3ctl",
	Short: "A brief description of your command",
	Long:  `AWS S3 服务相关命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("s3ctl called")
	},
}

func init() {
	// S3Cmd.Flags().BoolP("createBucket", "c", false, "创建S3桶")
	cmd.RootCmd.AddCommand(S3Cmd)
}

//// 创建 S3客户端
//func GetS3clinet(region ...string) *s3.Client {
//	if len(region) == 0 {
//		region = []string{"us-east-1"}
//	}
//}

// 创建S3桶
func CreateS3Bucket(ctx context.Context, s3 *s3.Client) error {
	return service.S3Op().CreateS3Bucket(ctx, s3)
}
