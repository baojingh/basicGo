package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        10/19/2023 2:44 PM
  @Description:
*/

func main() {
	defer ants.Release()
	var wg sync.WaitGroup
	statsChan := make(chan int, 10)

	con := func() {
		consumer(statsChan)
		wg.Done()
	}

	pro := func() {
		producer(statsChan)
		wg.Done()
	}

	wg.Add(1)
	ants.Submit(pro)

	wg.Add(1)
	ants.Submit(con)

	fmt.Printf("Running: %d, cap: %d", ants.Running(), ants.Cap())
	wg.Wait()
	println("Task finished")

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
