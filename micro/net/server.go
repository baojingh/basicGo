package net

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

/**
  @Author   : bob
  @Datetime : 6/5/2023 上午 10:41
  @File     : server.go
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

func ServerListen(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Info("Server fail to listen connection")
		return err
	}

	for {
		// 一直接收请求
		conn, err := listener.Accept()
		if err != nil {
			log.Info("Server fail to accept connection")
			return err
		}
		log.Info("Server receive connection request success")
		go func() {
			handleConn(conn)
		}()
	}
}

func handleConn(conn net.Conn) {
	for {
		bs := make([]byte, 20)
		_, err := conn.Read(bs)
		log.Infof("Server receives msg from Client: %s", string(bs))
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF || err != nil {
			_ = conn.Close()
			log.Info("Server fail to read msg and return")
			return
		}
		if err != nil {
			continue
		}
		res := handleMsg(bs)
		_, err = conn.Write(res)
		log.Infof("Server sends msg to Client: %s", string(res))
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			log.Errorf("Server fail to write msg to client and return, %s", err)
			return
		}
		if err != nil {
			continue
		}
	}
}

func handleMsg(req []byte) []byte {
	log.Info("Server handle message: ", string(req))
	res := string(req) + ", this is msg from server"
	return []byte(res)
}
