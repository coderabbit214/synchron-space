package main

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
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
						log.Println("文件/文件夹创建", event.Op, event.Name)
					} else if event.Has(fsnotify.Create) {
						log.Println("文件/文件夹创建", event.Op, event.Name)
					} else if event.Has(fsnotify.Rename) {
						log.Println("文件/文件夹重命名前", event.Op, event.Name)
					} else if event.Has(fsnotify.Remove) {
						log.Println("文件/文件夹删除", event.Op, event.Name)
					} else {
						log.Println("查看操作", event.Op, event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	// Add a path.
	err = watcher.Add("/Users/Mr_J/Downloads")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
