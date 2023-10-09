package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 7:56
  @File     : goroutine_test.go
  @Desc     :
*/

func TestGoroutine(t *testing.T) {
	var mut sync.Mutex
	var res int = 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			res++
		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(res)
}

func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	var counter int = 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)
}
