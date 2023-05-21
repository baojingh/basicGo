package gows

import "net/http"

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
}

func wspage(writer http.ResponseWriter, request *http.Request) {

}
