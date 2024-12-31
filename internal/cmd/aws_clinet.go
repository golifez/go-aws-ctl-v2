package cmd

import (
	"fmt"
	"path/filepath"
	"runtime"

	clinet "github.com/golifez/go-aws-ctl-v2/cmd/aws_clinet"
)

// 使用更可靠的路径解析方式
var (
	// 获取当前文件的目录
	_, currentFile, _, _ = runtime.Caller(0)                                        // 获取当前文件的目录
	projectRoot          = filepath.Join(filepath.Dir(currentFile), "../..")        // 获取项目根目录
	configPath           = filepath.Join(projectRoot, "hack/awsconfig/config")      // 获取配置文件路径
	credentialsPath      = filepath.Join(projectRoot, "hack/awsconfig/credentials") // 获取凭证文件路径
)

func GetNewAwsClinet(regoin string) (aws *clinet.AwsClient) {
	ConfigFile := clinet.NewConfigFile(configPath, credentialsPath, "default")
	ct := clinet.Awsconfig{
		ConfigFile: ConfigFile,
	}
	cfg, err := ct.LoadConfigFromConfigFile()
	if regoin != "" {
		cfg.Region = regoin
	}
	if err != nil {
		fmt.Println("err:", err)
	}
	client := clinet.NewAwsClient(cfg)
	fmt.Println("获取aws客户端成功:", client)
	return client
}
