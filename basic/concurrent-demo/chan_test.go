package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        10/18/2023 5:19 PM
  @Description:
*/

func TestFunc(t *testing.T) {
	statsChan := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(statsChan)
	}()
	go func() {
		defer wg.Done()
		consumer(statsChan)
	}()
	wg.Wait()
}

func producer(statsChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for true {
		num := rand.Intn(3)
		statsChan <- num
	}
}

func consumer(statsChan chan int) {
	for true {
		select {
		case val := <-statsChan:
			fmt.Printf("Consumer: %d\n", val)
		}
	}
}

// channel中本身只有一个元素，只能消费一个元素，不能继续消费，否则报错
func TestFunc2(t *testing.T) {
	ch1 := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go func() {
			ch1 <- "a"
		}()
	}

	for true {
		// ERROR : fatal error: all goroutines are asleep - deadlock!
		//当进入第5个循环时，由于ch1通道内的数字都输出完毕，<-ch1一直等待新的数据，
		//所以变为阻塞状态，但是此时所有goroutine都阻塞了，所以发生了死锁，
		v1 := <-ch1
		fmt.Println(v1)
	}

	time.Sleep(1 * time.Second)
}

// 方法1: 一个goroutine发送，一个接收
// 方法2: 给channel设置缓冲区，而不需要接收端同时接收 (缓冲区如果超出大小同样会造成死锁)
func TestFunc3(t *testing.T) {
	ch := make(chan string)
	ch <- "a"

	//fatal error: all goroutines are asleep - deadlock!
	v1 := <-ch
	println(v1)
}

func TestFunc4(t *testing.T) {
	c := make(chan int, 2)
	c <- 3
	c <- 5
	close(c)
	println(len(c), cap(c))
	x, ok := <-c
	println(x, ok)

	println()

	x, ok = <-c
	println(len(c), cap(c))
	println(x, ok)

	println()

	x, ok = <-c
	println(len(c), cap(c))
	println(x, ok)

	println()

	x, ok = <-c
	println(len(c), cap(c))
	println(x, ok)

	//the 2 lines will cause deadlock
	c <- 6
	close(c)

}
