package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        10/18/2023 5:19 PM
  @Description:
*/

func TestFunc(t *testing.T) {
	statsChan := make(chan int)
	fmt.Printf("Channel cap is %d, size is %d\n", cap(statsChan), len(statsChan))

}

func producer(statsChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for true {
		num := rand.Intn(101)
		statsChan <- num
	}
}

func consumer(statsChan chan int) {

}
