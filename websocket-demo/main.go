package main

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

/**
@Author   : bob
@Datetime : 3/5/2023 下午 2:08
@File     : net_test.go
@Desc     :
*/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	// websocket: request origin not allowed by Upgrader.CheckOrigin
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

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info("Failed to upgrade to WebSocket:", err)
		return
	}
	log.Info("create websocket success")
	defer conn.Close()

	for {
		messageLen, message, err := conn.ReadMessage()
		if err != nil {
			log.Info("Failed to read message from WebSocket:", err)
			break
		}
		log.Infof("Server received message: %s", message)

		err = conn.WriteMessage(messageLen, message)
		if err != nil {
			log.Info("Failed to write message to WebSocket:", err)
			break
		}
		log.Infof("Server send message: %s", message)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	log.Info("Starting server on port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Info("Failed to start server:", err)
	}
}
