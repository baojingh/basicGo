package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

/**
  @Author   : bob
  @Datetime : 7/5/2023 下午 3:47
  @File     : select_demo_test.go
  @Desc     :
*/

/**
select就是用来监听和channel有关的IO操作，当 IO 操作发生时，触发相应的动作。

select {
case <- chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程

*/

/*
*
如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行
否则的话，如果有default分支，则执行default分支语句，
如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行
*/
func TestSelect1(t *testing.T) {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		close(c)
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- 2
	}()

	go func() {
		time.Sleep(time.Second * 3)
		ch2 <- 3
	}()
	fmt.Println("Blocking")

	select {
	case <-c:
		fmt.Printf("unblocked %v later \n", time.Since(start))
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
		//default:
		//	fmt.Println("default")
	}
	fmt.Println("main finished")
}

/*
*
测试超时机制
*/
func TestTimeoutDemo(t *testing.T) {
	ch := make(chan int)
	quitCh := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				log.Info("consumer: ", num)
			case <-time.After(time.Second * 3):
				log.Info("timeout")
				quitCh <- true
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
		log.Info("producer: ", i)
	}

	// 如果这个channel中没有数据，main会一直阻塞在这里
	// 当超时控制中执行了quitCh <- true，说明channel中存在数据
	// main就不在阻塞，就继续执行，实现了程序的超时控制
	<-quitCh

	log.Info("main finished")

}
