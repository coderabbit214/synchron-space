package watcher

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"strings"
	"synchron-space/model/file"
	"synchron-space/pkg"
)

var PathList []string

// BeginWatcher 开启监听
func BeginWatcher(watcher *fsnotify.Watcher) {
	service, err := file.GetService()
	if err != nil {
		fmt.Println("file.GetService err:", err)
		return
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			var objectName string
			for _, path := range PathList {
				if strings.Contains(event.Name, path) {
					objectName = strings.Replace(event.Name, path, "", 1)
				}
			}
			split := strings.Split(event.Name, "/")
			if split[len(split)-1] != ".DS_Store" {
				//创建
				if event.Has(fsnotify.Write) {

				} else if event.Has(fsnotify.Create) {
					//文件处理
					if pkg.IsFile(event.Name) && objectName != "" {
						fmt.Println("en", event.Name)
						fmt.Println("object:", objectName)
						err := service.Upload(event.Name, objectName)
						if err != nil {
							fmt.Println("service.Upload err:", err)
							return
						}
					}
					//文件夹处理
					if pkg.IsDir(event.Name) {

					}
				} else if event.Has(fsnotify.Rename) {
					if event.Has(fsnotify.Remove) {
						log.Println("文件夹删除/重命名前", event.Op, event.Name)
					} else {
						log.Println("文件删除/重命名前", event.Op, event.Name)
					}
				} else if event.Has(fsnotify.Remove) {
					log.Println("文件删除", event.Op, event.Name)
				} else {
					//log.Println("查看操作", event.Op, event.Name)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

// AddWatcher 增加监听目录
func AddWatcher(watcher *fsnotify.Watcher, path string) (err error) {
	//TODO:嵌套目录判断

	PathList = append(PathList, path)
	err = watcher.Add(path)
	dirs, err := pkg.GetAllDirs(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, dir := range dirs {
		err = watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	if err != nil {
		log.Fatal(err)
		return err
	}
	return
}
