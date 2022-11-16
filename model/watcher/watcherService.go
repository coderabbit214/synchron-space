package watcher

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"strings"
	"synchron-space/model/file/oss"
)

// BeginWatcher 开启监听
func BeginWatcher(watcher *fsnotify.Watcher) {
	service := oss.GetService()
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			split := strings.Split(event.Name, "/")
			if split[len(split)-1] != ".DS_Store" {
				//创建新文件
				if event.Has(fsnotify.Write) {
					service.Upload(event.Name)
					log.Println("文件/文件夹创建", event.Op, event.Name)
				} else if event.Has(fsnotify.Create) {
					log.Println("文件/文件夹创建", event.Op, event.Name)
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
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return
}
