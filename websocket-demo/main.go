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

var msgChan = make(chan []byte, 128)

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

	// ws://127.0.0.1:8081/logs?id=111
	// Get the container ID from the query parameter
	containerId := r.URL.Query().Get("id")
	log.Infof("get containerId from request, %s", containerId)

	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				log.Info("fail to read data from client, ", err)
				break
			}
			msgChan <- data
		}
	}()

	for data := range msgChan {
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Error("fal to send data to client, ", err)
			break
		}
	}

}

func main() {
	http.HandleFunc("/logs", handleWebSocket)

	log.Info("Starting server on port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Info("Failed to start server:", err)
	}
}
