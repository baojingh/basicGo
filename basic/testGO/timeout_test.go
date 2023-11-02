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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 启动一个 Goroutine 执行耗时操作
	go func() {
		fmt.Println("start sub work")
		time.Sleep(10 * time.Second) // 模拟耗时操作
		fmt.Println("finish sub work")
		cancel() // 在耗时操作完成后取消上下文
	}()

	// 在主 Goroutine 中等待上下文的完成信号
	select {
	case <-ctx.Done():
		fmt.Println("操作已超时")
	}
}
