package gows

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

/**
  @Author   : bob
  @Datetime : 2023-05-21 下午 10:32
  @File     : init_acc.go
  @Desc     :
*/

var (
	clientManager = NewClientManager()
)

func StartWebSocket() {
	http.HandleFunc("/acc", wspage)
	go clientManager.start()
	_ = http.ListenAndServe(":8081", nil)

}

func wspage(w http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		// websocket: request origin not allowed by Upgrader.CheckOrigin
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Info("Failed to upgrade to WebSocket:", err)
		return
	}
	remoteIp := conn.RemoteAddr().String()
	log.Infof("create websocket success, client IP is %s", remoteIp)

	currentTime := uint64(time.Now().Unix())
	client := NewClient(remoteIp, conn, currentTime)
	go client.read()
	go client.write()

	clientManager.Register <- client
	log.Info("client is registered success...")
}
