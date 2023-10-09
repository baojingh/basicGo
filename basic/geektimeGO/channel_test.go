package main

import (
	"basicGo/ws-demo/wshandle"
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        5/12/2023 11:47 AM
  @Description:
*/

func TestChannel2(t *testing.T) {
	// ch1 := make(chan int) 默认长度是0
	ch1 := make(chan int, 2)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			fmt.Println("producer: ", i)
		}
	}()

	go func() {
		for true {
			msg := <-ch1
			fmt.Println("consumer: ", msg)
			time.Sleep(time.Second * 2)
		}

	}()

	time.Sleep(time.Second * 10)
	fmt.Println("done")

}

func TestChannel3(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			wshandle.MsgChan <- strconv.Itoa(i)
			fmt.Println("producer: ", i)
		}
	}()
}

// 给channel塞入一个值，是否可以支持取两次？
func TestChannel4(t *testing.T) {

	ch := make(chan int, 1)

	go func() {
		ch <- 123 // 将值123发送到通道ch
	}()

	val1 := <-ch // 第一次接收值
	fmt.Println("第一次接收到的值:", val1)

}
