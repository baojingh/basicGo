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

var monitorPath = ""

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error("cannot create watcher ", err)
		return
	}

	go func() {
		for true {
			select {
			case event, _ := <-watcher.Events:
				log.Info("some", event.Name)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Info("fail to monitor path", err)

			}

		}
	}()

	watcher.Add(monitorPath)

}
