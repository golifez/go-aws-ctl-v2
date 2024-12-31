package s3

import (
	"fmt"

	"github.com/golifez/go-aws-ctl-v2/internal/cmd"
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
	fmt.Println("s3 init")
	cmd.RootCmd.AddCommand(S3Cmd)
}
