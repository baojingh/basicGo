package task

import (
	"basicGo/websocket-demo/gows"
	"time"
)

/**
  @Author   : bob
  @Datetime : 2023-05-23 下午 10:42
  @File     : clean_task.go
  @Desc     :
*/

func InitTask() {
	WSTimer(3*time.Second, 30*time.Second, cleanConnection,
		"", nil, nil)
}

func cleanConnection(param interface{}) (result bool) {
	result = true
	gows.ClearTimeoutConnections()
	return

}
