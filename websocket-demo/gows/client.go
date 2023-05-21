package gows

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

/**
  @Author   : bob
  @Datetime : 2023-05-21 下午 5:55
  @File     : client.go
  @Desc     :
*/

const (
	// user connection timeout
	heartbeatExpirationTime = 1 * 60
)

var (
	clientManager = NewClientManager()
)

type Client struct {
	// // 客户端地址
	Addr string

	Socket *websocket.Conn

	// 待发送的数据
	Send chan []byte

	// 登录的平台Id app/web/ios
	AppId uint32

	// 用户Id，用户登录以后才有
	UserId string

	// 首次连接事件
	FirstTime uint64

	// 用户上次心跳时间
	HeartbeatTime uint64

	// 登录时间 登录以后才有
	LoginTime uint64
}

func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Addr:          addr,
		Socket:        socket,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		Send:          make(chan []byte, 100),
	}
	return client
}

// read read data from client
func (c *Client) read() {
	defer func() {
		log.Info("read the data from client, close send channel", c)
		close(c.Send)
	}()

	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			log.Error("fail to read data from client, ", err)
			return
		}
		log.Infof("read data from client, %s", string(msg))
		ProcessData(c, msg)
	}
}

// write send data to client
func (c *Client) write() {
	defer func() {
		clientManager.Unregister <- c
		c.Socket.Close()
		log.Info("Client send data defer, ", c)
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				log.Errorf("close conn, %s, %s", c.Addr, ok)
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func (c *Client) SendMsg(msg []byte) {
	if c != nil {
		return
	}
	c.Send <- msg
}

func (c *Client) close() {
	close(c.Send)
}
