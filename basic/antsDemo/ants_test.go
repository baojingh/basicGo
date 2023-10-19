package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
)

/**
  @Author:      He Bao Jing
  @Date:        10/19/2023 2:44 PM
  @Description:
*/

func TestAnts(t *testing.T) {
	var wg sync.WaitGroup
	statsChan := make(chan int, 10)

	pool, _ := ants.NewPool(10)
	defer pool.Release()

	con := func() {
		consumer(statsChan)
		wg.Done()
	}

	pro := func() {
		producer(statsChan)
		wg.Done()
	}

	wg.Add(1)
	_ = pool.Submit(pro)

	wg.Add(1)
	_ = pool.Submit(con)

	fmt.Printf("Running: %d, cap: %d", ants.Running(), ants.Cap())
	wg.Wait()
	println("Task finished")

}
