package clinet

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// 使用更可靠的路径解析方式
var (
	// 获取当前文件的目录
	_, currentFile, _, _ = runtime.Caller(0)                                        // 获取当前文件的目录
	projectRoot          = filepath.Join(filepath.Dir(currentFile), "../")          // 获取项目根目录
	configPath           = filepath.Join(projectRoot, "hack/awsconfig/config")      // 获取配置文件路径
	credentialsPath      = filepath.Join(projectRoot, "hack/awsconfig/credentials") // 获取凭证文件路径
)

// 通过配置文件创建aws客户端
func GetNewAwsClinetWithRegion(regoin string) (aws *AwsClient) {
	ConfigFile := NewConfigFile(configPath, credentialsPath, "default")
	ct := Awsconfig{
		ConfigFile: ConfigFile,
	}
	cfg, err := ct.LoadConfigFromConfigFile()
	if regoin != "" {
		cfg.Region = regoin
	}
	fmt.Println("区域是", cfg.Region)
	ctv, err := cfg.Credentials.Retrieve(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("秘钥是", ctv.AccessKeyID, ctv.SecretAccessKey)

	if err != nil {
		fmt.Println("err:", err)
	}
	client := NewAwsClient(cfg)
	fmt.Println("获取aws客户端成功:", client)
	return client
}

// 通过传入key和ID创建
func GetNewAwsClinetWithKeyAndID(keyID, secretKey string) (aws *AwsClient) {
	AccessKeyIdType := AccessKeyIdType{
		IsEnv:           false,
		AccessKeyId:     keyID,
		SecretAccessKey: secretKey,
		Region:          "us-east-1",
	}
	cfgop := Awsconfig{
		AccessKeyIdType: &AccessKeyIdType,
	}
	cfg, err := cfgop.LoadConfigFromAccessKeyId()
	if err != nil {
		fmt.Println("创建aws客户端失败:", err)
	}
	client := NewAwsClient(cfg)
	return client
}

// 通过环境变量创建
func GetNewAwsClinetWithEnv() (aws *AwsClient) {
	AccessKeyIdType := AccessKeyIdType{
		IsEnv: true,
	}
	cfgop := Awsconfig{
		AccessKeyIdType: &AccessKeyIdType,
	}
	cfg, err := cfgop.LoadConfigFromAccessKeyId()
	if err != nil {
		fmt.Println("创建aws客户端失败:", err)
	}
	client := NewAwsClient(cfg)
	return client
}

func GetS3clinet(region string) *s3.Client {
	awsclinet := GetNewAwsClinetWithRegion(region)
	if awsclinet == nil {
		return nil
	}
	return awsclinet.S3Client
}
