package gows

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
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

func init() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
}

func StartWebSocket() {
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/acc", wspage)
	go clientManager.start()
	port := ":8081"
	log.Infof("start websocket service in port %s", port)
	_ = http.ListenAndServe(port, nil)
}

func metrics(w http.ResponseWriter, req *http.Request) {
	managerInfo := clientManager.GetManagerInfo("true")
	infoByte := map2Byte(managerInfo)
	w.Write(infoByte)
	log.Infof("monitor client manager, %s", string(infoByte))
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

func map2Byte(mapTmp map[string]interface{}) []byte {
	str, _ := json.Marshal(mapTmp)
	return str

}
