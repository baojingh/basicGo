package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
  @Author   : bob
  @Datetime : 2023-11-19 上午 11:56
  @File     : cus_consumer_producer_test.go
  @Desc     :
*/

type Queue struct {
	Data []int
	Cond *sync.Cond
}

func TestConPro(t *testing.T) {
	q := &Queue{
		Data: []int{},
		Cond: sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		for i := 0; i < 10000; i++ {
			q.Enqueue(i)
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		q.Dequeue()
		time.Sleep(1 * time.Second)
	}
}

func (q *Queue) Enqueue(val int) {
	q.Cond.L.Lock()
	defer q.Cond.L.Unlock()
	q.Data = append(q.Data, val)
	q.Cond.Broadcast()
	fmt.Printf("Produce val %v \n", val)
}

func (q *Queue) Dequeue() int {
	q.Cond.L.Lock()
	defer q.Cond.L.Unlock()
	if len(q.Data) == 0 {
		q.Cond.Wait()
	}
	val := q.Data[0]
	q.Data = q.Data[1:]
	fmt.Printf("Consume val %v \n", val)
	return val
}
