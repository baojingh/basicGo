package main

/**
  @Author   : bob
  @Datetime : 2023-08-03 下午 8:20
  @File     : main.go
  @Desc     :
*/

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

var monitorPath = "basic/file-detector/"

func main() {
	watcher, err := fsnotify.NewWatcher()
	log.Info("watcher is created success, ", watcher.Events)
	defer watcher.Close()

	doneChan := make(chan bool)

	if err != nil {
		log.Error("cannot create watcher ", err)
		return
	}

	go func() {
		for true {
			log.Info("in loop")
			select {
			case event, _ := <-watcher.Events:
				log.Infof("some, %s, %s", event, event.Name)
			case err, _ := <-watcher.Errors:
				log.Info("fail to monitor path", err)
			}
		}
	}()

	_ = watcher.Add(monitorPath)
	log.Info("path is added ", watcher.WatchList())

	<-doneChan
	log.Info("exit the process")
}
