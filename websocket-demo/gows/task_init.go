package gows

import "time"

/**
  @Author   : bob
  @Datetime : 2023-05-23 下午 10:40
  @File     : task_init.go
  @Desc     :
*/

type TimerFunc func(interface{}) bool

func WSTimer(delay, tick time.Duration, fun TimerFunc, param interface{},
	funcDefer TimerFunc, paramDefer interface{}) {
	go func() {
		defer func() {
			if funcDefer != nil {
				funcDefer(paramDefer)
			}
		}()

		if fun == nil {
			return
		}
		t := time.NewTimer(delay)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				if fun(param) == false {
					return
				}
				t.Reset(tick)
			}
		}
	}()
}
