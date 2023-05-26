package dataSource

/**
  @Author:      He Bao Jing
  @Date:        5/12/2023 1:15 PM
  @Description:
*/

import (
	"basicGo/ws-demo/wshandle"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func Produce() {
	for i := 0; i < 100; i++ {
		msg := strconv.Itoa(i)
		wshandle.MsgChan <- msg
		log.Infof("send data to channel, %s", msg)
		time.Sleep(time.Second * 1)
	}
	log.Info("produce data success")
}
