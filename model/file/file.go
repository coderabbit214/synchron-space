package file

import (
	"fmt"
	oss2 "synchron-space/model/file/oss"
	"synchron-space/pkg"
)

// Service 文件相关接口
type Service interface {
	// Upload 文件上传
	Upload(filePath string, objectName string) error
	// Download 文件下载
	Download()
}

type File struct {
	LocalPath string
	// 类型 file 文件,folder 文件夹
	Type string
	// 状态 已同步(finish) 未同步(ready) 正在同步(synchronizing)
	State string
	Files []File
}

// GetService 获取服务
func GetService() (service Service, err error) {

	switch pkg.GlobalConfig.Server {
	case "oss":
		service, err = oss2.NewService()
		if err != nil {
			fmt.Println("oss2.NewService err:", err)
			return
		}
		return service, err
	}
	return
}
