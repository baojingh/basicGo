package net

import (
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

/**
  @Author   : bob
  @Datetime : 6/5/2023 下午 5:26
  @File     : client.go
  @Desc     :
*/

func Client(addr string) error {

	conn, err := net.DialTimeout("tcp", addr, time.Second*3)
	if err != nil {
		log.Info("Client fail to connect server")
	}

	defer func() {
		_ = conn.Close()
	}()

}
