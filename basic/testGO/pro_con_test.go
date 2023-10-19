package main

import (
	"fmt"
	"sync"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 5/5/2023 下午 1:34
  @File     : pro_con_test.go
  @Desc     :
*/

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			data := i + 100
			ch <- data
			fmt.Println("producer: ", data)
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 12; i++ {
			if data, ok := <-ch; ok {
				fmt.Println("consumer: ", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 11)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

}
