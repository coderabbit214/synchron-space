package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	watcher2 "synchron-space/model/watcher"
	"synchron-space/pkg"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	// 开启监听
	go watcher2.BeginWatcher(watcher)
	// 添加监听目录
	watcher2.AddWatcher(watcher, pkg.GlobalConfig.WatcherPath)
	err = watcher.Add("/Users/mr_j/Downloads")
	if err != nil {
		log.Fatal(err)
	}

	//阻塞主线程
	<-make(chan struct{})
}
