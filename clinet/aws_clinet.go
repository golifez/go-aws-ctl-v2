package clinet

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsClient struct {
	S3Client *s3.Client
}

type Awsconfig struct {
	// 指定ID和KEY创建cfg
	AccessKeyIdType *AccessKeyIdType
	// 通过配置文件创建cfg
	ConfigFile *ConfigFile
}

type AccessKeyIdType struct {
	// 从环境变量中获取
	IsEnv           bool //是否从环境变量中获取
	AccessKeyId     string
	SecretAccessKey string
	Region          string
}
type ConfigFile struct {
	ConfigPath      string // 配置config
	CredentialsPath string // 配置credentials
	ProfileName     string // 配置块
}

// 默认配置
const (
	DefaultRegion      = "us-east-1"
	DefaultConfig      = "~/.aws/config"
	DefaultCredentials = "~/.aws/credentials"
	DefuletUser        = "default" // 默认配置块
)

func NewAwsClient(cfg aws.Config) *AwsClient {
	return &AwsClient{
		S3Client: s3.NewFromConfig(cfg),
	}
}

// 直接传入id和key,获取 AWS 配置
func (c *Awsconfig) LoadConfigFromAccessKeyId() (aws.Config, error) {
	// 从环境变量中获取
	if c.AccessKeyIdType.IsEnv {
		cfg, err := GetconfigFromEnv()
		if err != nil {
			return aws.Config{}, err
		}
		return cfg, nil
	}
	// 如果通过access_id创建
	if c.AccessKeyIdType.AccessKeyId != "" {
		cfg, err := GetAwsConfigFromId(c.AccessKeyIdType.AccessKeyId, c.AccessKeyIdType.SecretAccessKey, c.AccessKeyIdType.Region)
		if err != nil {
			return aws.Config{}, err
		}
		return cfg, nil
	}
	return aws.Config{}, errors.New("access_id is empty")
}

func GetAwsConfigFromId(accessKeyId, secretAccessKey, Region string) (aws.Config, error) {
	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, ""))
	return aws.Config{
		Region:      Region,
		Credentials: creds,
	}, nil
}
func GetconfigFromEnv() (aws.Config, error) {
	// 从环境变量中获取
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = DefaultRegion
	}
	cfg, err := GetAwsConfigFromId(accessKeyId, secretAccessKey, region)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

// 通过配置文件创建cfg
func (c *Awsconfig) LoadConfigFromConfigFile() (aws.Config, error) {
	// 指定配置文件和凭据文件的位置
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigFiles([]string{c.ConfigFile.ConfigPath}),           // 配置文件路径
		config.WithSharedCredentialsFiles([]string{c.ConfigFile.CredentialsPath}), // 凭据文件路径
		config.WithSharedConfigProfile(c.ConfigFile.ProfileName),                  // 指定 Profile
	)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

// 传入配置文件路径，创建ConfigFile
func NewConfigFile(configPath, credentialsPath, profileName string) *ConfigFile {
	return &ConfigFile{
		ConfigPath:      configPath,
		CredentialsPath: credentialsPath,
		ProfileName:     profileName,
	}
}

// 使用默认配置文件创建cfg
func NewDefaultConfigFile() *ConfigFile {
	return &ConfigFile{
		ConfigPath:      DefaultConfig,
		CredentialsPath: DefaultCredentials,
		ProfileName:     DefuletUser,
	}
}
