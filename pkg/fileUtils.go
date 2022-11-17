package pkg

import (
	"os"
	"path/filepath"
)

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// GetAllDirs 获取所有文件夹
func GetAllDirs(dirPth string) (dirs []string, err error) {
	fis, err := os.ReadDir(filepath.Clean(filepath.ToSlash(dirPth)))
	if err != nil {
		return nil, err
	}
	for _, f := range fis {
		_path := filepath.Join(dirPth, f.Name()) + "/"
		if f.IsDir() {
			dirs = append(dirs, _path)
			fs, _ := GetAllDirs(_path)
			dirs = append(dirs, fs...)
			continue
		}
	}
	return dirs, nil
}
