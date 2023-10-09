package wshandle

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        5/12/2023 1:06 PM
  @Description:
*/

func TestStartWs(t *testing.T) {
	StartWs()
}

func TestMockProducer(t *testing.T) {
	go func() {
		for i := 0; i < 100; i++ {
			msg := strconv.Itoa(i)
			MsgChan <- msg
			log.Infof("send data to channel, %s", msg)
			time.Sleep(time.Second * 1)
		}
	}()
	time.Sleep(time.Second * 39999)

}
