package net

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

/**
  @author:      baojing.he
  @Date:        ${DATE} ${TIME}
  @Description:
*/

/**
 @Author   : bob
 @Datetime : 7/5/2023 下午 4:42
 @File     : net_test.go
 @Desc     : 单元测试发现server最后再次给client发送消息，
			  也能正常发出去，但是server的返回信息会err，导致执行if中的return

*/

/**

INFO[2023-05-07 23:28:52] Client sends connection request success
INFO[2023-05-07 23:28:52] Server receive connection request success
INFO[2023-05-07 23:28:52] Client sends msg to server success
INFO[2023-05-07 23:28:52] Server receives msg from Client: hello serve
INFO[2023-05-07 23:28:52] Server handle message: hello serve
INFO[2023-05-07 23:28:52] Server sends msg to Client: hello serve         , this is msg from server
INFO[2023-05-07 23:28:52] Client receive msg from server: hello serve         , this is msg from server
INFO[2023-05-07 23:28:52] Server receives msg from Client:
--- PASS: TestTcp (0.52s)
PASS
INFO[2023-05-07 23:28:52] Server fail to read msg and return

*/

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
}

func TestTcp(t *testing.T) {
	addr := ":8082"
	go func() {
		err := ServerListen(addr)
		if err != nil {
			log.Info("fail to start server")
		}
	}()

	err := ClientConnect(addr)
	if err != nil {
		log.Info("fail to start client")
	}

}
