package wshandle

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        5/12/2023 10:34 AM
  @Description:
*/

var host = ":8081"

// 消息体
type wsMessage struct {
	messageType int
	data        []byte
}

// 客户端结构
type clientStruct struct {
	conn     *websocket.Conn
	sendChan chan wsMessage
}

// 服务端结构
type serverStruct struct {
	clients        map[*clientStruct]bool
	registerChan   chan *clientStruct
	unregisterChan chan *clientStruct
}

var MsgChan chan string

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2014,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
}

func (c *clientStruct) read() {
	defer func() {
		c.conn.Close()
	}()
	for true {
		messageType, data, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("Server fail to receive data from client, ", err)
			break
		}
		log.Infof("Server receive data from client, %s", string(data))
		msg := wsMessage{
			messageType: messageType,
			data:        data,
		}
		c.sendChan <- msg
	}
}

func (c *clientStruct) write() {
	defer func() {
		c.conn.Close()
	}()
	for true {

		go func() {
			for i := 0; i < 5; i++ {
				msg := strconv.Itoa(i)
				MsgChan <- msg
				time.Sleep(time.Second * 1)
			}
		}()
		time.Sleep(time.Second * 5)

		// 阻塞，直到有数据进入msgChan
		msg := <-MsgChan
		log.Infof("get msg from channel, %s", msg)
		err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Info("Failed to write message to WebSocket:", err)
			break
		}
		log.Infof("Server send message: %s", msg)
	}
}

func (s *serverStruct) serverHandleRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("fail to create ws success, ", err)
		return
	}
	log.Infof("get ws success")

	cli := &clientStruct{
		conn:     conn,
		sendChan: make(chan wsMessage, 100),
	}
	log.Info("client has registered success")
	go cli.read()
	go cli.write()
}

func (s *serverStruct) unregister(cli *clientStruct) {
	if _, ok := s.clients[cli]; ok {
		delete(s.clients, cli)
		cli.conn.Close()
	}
}

func (s *serverStruct) register(cli *clientStruct) {
	if cli != nil && cli.conn != nil {
		s.clients[cli] = true
		s.registerChan <- cli
	}
}

func StartWs() {
	MsgChan = make(chan string, 100)
	ser := &serverStruct{
		clients:      make(map[*clientStruct]bool, 100),
		registerChan: make(chan *clientStruct, 100),
	}
	http.HandleFunc("/ws", ser.serverHandleRequest)
	log.Infof("Starting server on port %s", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Error("Failed to start server:", err)
	}
}
