package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
  Author     : He Bao Jing
  Date       : 11/2/2023 3:22 PM
  Description:
*/

func TestTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	ch := make(chan int)

	// 启动一个 Goroutine 执行耗时操作
	go func() {
		fmt.Println("start sub work")
		time.Sleep(4 * time.Second) // 模拟耗时操作
		ch <- 1
	}()

	// 在主 Goroutine 中等待上下文的完成信号
	select {
	case <-ctx.Done():
		fmt.Println("操作已超时，退出")
	case val := <-ch:
		println("get val ", val)
	}
}
