package file

// Service 文件相关接口
type Service interface {
	// Upload 文件上传
	Upload(filePath string) bool
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
