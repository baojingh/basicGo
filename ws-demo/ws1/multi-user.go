package main1

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	// websocket: request origin not allowed by Upgrader.CheckOrigin
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义 WebSocket 连接最大读取字节数
const maxMessageSize = 1024 * 1024

// 定义 WebSocket 消息类型
type message struct {
	messageType int
	data        []byte
}

// 定义 WebSocket 客户端连接
type client struct {
	conn     *websocket.Conn
	sendChan chan message
}

// 定义 WebSocket 服务器结构体
type server struct {
	clients        map[*client]bool
	broadcastChan  chan message
	registerChan   chan *client
	unregisterChan chan *client
}

// 处理 WebSocket 连接请求
func (s *server) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, maxMessageSize, maxMessageSize)
	if err != nil {
		log.Println(err)
		return
	}

	c := &client{
		conn:     ws,
		sendChan: make(chan message),
	}

	s.registerChan <- c
	defer func() { s.unregisterChan <- c }()

	go c.writePump()
	c.readPump()
}

// 读取客户端发送的消息
func (c *client) readPump() {
	defer func() {
		c.conn.Close()
		close(c.sendChan)
	}()

	for {
		messageType, data, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error:", err)
			}
			break
		}

		m := message{messageType: messageType, data: data}
		c.sendChan <- m
	}
}

// 向客户端发送消息
func (c *client) writePump() {
	defer func() { c.conn.Close() }()

	for {
		select {
		case m, ok := <-c.sendChan:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(m.messageType)
			if err != nil {
				return
			}

			w.Write(m.data)

			n := len(c.sendChan)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				wm := <-c.sendChan
				w.Write(wm.data)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// 启动 WebSocket 服务器
func startWebSocketServer() {
	s := &server{
		clients:        make(map[*client]bool),
		broadcastChan:  make(chan message),
		registerChan:   make(chan *client),
		unregisterChan: make(chan *client),
	}

	// 处理新连接请求
	http.HandleFunc("/", s.serveWebSocket)

	// 处理广播任务
	go func() {
		for {
			m := <-s.broadcastChan

			for c := range s.clients {
				select {
				case c.sendChan <- m:
				default:
					close(c.sendChan)
					delete(s.clients, c)
				}
			}
		}
	}()

	// 处理注册/注销请求
	for {
		select {
		case c := <-s.registerChan:
			s.clients[c] = true
		case c := <-s.unregisterChan:
			if _, ok := s.clients[c]; ok {
				close(c.sendChan)
				delete(s.clients, c)
			}
		}
	}
}

func main() {
	startWebSocketServer()
}
