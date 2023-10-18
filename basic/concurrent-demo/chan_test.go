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
		for true {
			fmt.Printf("Channel cap: %d, size: %d\n", cap(statsChan), len(statsChan))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		producer(statsChan)
	}()

	go func() {
		defer wg.Done()
		consumer(statsChan)
	}()

	wg.Wait()

	fmt.Println("Service is done")
}

func producer(statsChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for true {
		num := rand.Intn(101)
		statsChan <- num
	}
}

func consumer(statsChan chan int) {
	for true {
		select {
		case val := <-statsChan:
			fmt.Printf("Consumer: %d\n", val)
		default:

		}
		time.Sleep(5 * time.Second)
	}
}
