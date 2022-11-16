package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/yaml.v2"
	"os"
	"synchron-space/model/file"
)

var ossService = OssService{}

type (
	OssService struct {
		Client     *oss.Client
		BucketName string
		Path       string
	}

	Config struct {
		Oss struct {
			Endpoint        string `yaml:"endpoint"`
			AccessKeyId     string `yaml:"access-key-id"`
			AccessKeySecret string `yaml:"access-key-secret"`
			BucketName      string `yaml:"bucket-name"`
			Path            string `yaml:"path"`
		}
	}
)

// GetService 获取服务
func GetService() file.Service {
	if ossService.Client == nil {
		newService()
	}
	return ossService
}

// 创建连接
func newService() {
	var config Config
	yamlFile, err := os.ReadFile("/Users/Mr_J/Desktop/synchron-space/conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ossConfig := config.Oss
	client, err := oss.New(ossConfig.Endpoint, ossConfig.AccessKeyId, ossConfig.AccessKeySecret)
	if err != nil {
		fmt.Println("oss.New err:", err)
	}
	ossService.Client = client
	ossService.BucketName = ossConfig.BucketName
	ossService.Path = ossConfig.Path
}

func (s OssService) Upload(filePath string) bool {
	//文件判断

	return false
}

func (s OssService) Download() {

}
