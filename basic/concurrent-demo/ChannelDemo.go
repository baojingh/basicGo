package main

import (
	"fmt"
	"time"
)

/**
  @Author   : bob
  @Datetime : 2023-10-14 上午 11:27
  @File     : ChannelDemo.go
  @Desc     :
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		count1(c1)
	}()

	go func() {
		count2(c2)
	}()

	for {
		select {
		case d := <-c1:
			fmt.Println(d)
		case d := <-c2:
			fmt.Println(d)
		}
	}
	time.Sleep(10 * time.Second)
}

func count1(c chan string) {
	for {
		c <- "count1"
		time.Sleep(500 * time.Millisecond)
	}
}

func count2(c chan string) {
	for {
		c <- "count2"
		time.Sleep(100 * time.Millisecond)
	}
}
