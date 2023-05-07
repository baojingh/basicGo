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

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
}

func ClientConnect(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, time.Second*3)
	if err != nil {
		log.Info("Client fail to connect server")
	}
	log.Info("Client sends connection request success")
	defer func() {
		_ = conn.Close()
	}()
	_, err = conn.Write([]byte("hello serve"))
	if err != nil {
		log.Info("Client fail to write data")
		return err
	}
	log.Info("Client sends msg to server success")
	bs := make([]byte, 100)
	_, err = conn.Read(bs)
	if err != nil {
		log.Info("Client fail to read data from server")
		return err
	}
	log.Infof("Client receive msg from server: %s", string(bs))
	return nil

}
