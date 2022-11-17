package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gopkg.in/yaml.v2"
	"os"
)

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

// NewService 创建连接
func NewService() (ossService OssService, err error) {
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
	return ossService, err
}

func (s OssService) Upload(filePath string, objectName string) error {
	//文件判断
	bucket, err := s.Client.Bucket(s.BucketName)
	if err != nil {
		fmt.Println("s.Client.Bucket err:", err)
		return err
	}
	fmt.Println("oss:", s.Path+objectName, filePath)
	err = bucket.PutObjectFromFile(s.Path+objectName, filePath)
	if err != nil {
		fmt.Println("bucket.PutObjectFromFile err:", err)
		return err
	}
	return err
}

func (s OssService) Download() {

}
