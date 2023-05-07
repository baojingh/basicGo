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

func Server(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Info("Server fail to listen connection")
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Info("Server fail to accept connection")
			return err
		}
		go func() {
			handleConn(conn)
		}()
	}

}

func handleConn(conn net.Conn) {
	for {
		bs := make([]byte, 8)
		_, err := conn.Read(bs)
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			log.Info("Server fail to read msg and return")
			return
		}
		if err != nil {
			continue
		}
		res := handleMsg(bs)
		_, err = conn.Write(res)
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			log.Info("Server fail to write msg to client and return")
			return
		}
		if err != nil {
			continue
		}

	}
}

func handleMsg(req []byte) []byte {
	log.Info("Server handle message")
	res := make([]byte, 2*len(req))
	copy(res[:len(req)], req)
	copy(res[len(req):], req)
	return res
}
