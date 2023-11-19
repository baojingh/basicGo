package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
  @Author   : bob
  @Datetime : 2023-11-19 上午 11:23
  @File     : rwlock_test.go
  @Desc     :
*/

func TestUnsafeWrite(t *testing.T) {
	conflictMap := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			conflictMap[1] = i
		}()
	}
}

type SafeMap struct {
	sameMap map[int]int
	sync.Mutex
}

func TestSafeLock(t *testing.T) {
	safe := SafeMap{
		sameMap: make(map[int]int),
		Mutex:   sync.Mutex{},
	}
	for i := 0; i < 10000; i++ {
		go safe.write(1, 1)
	}
	time.Sleep(10 * time.Second)
}

func (s *SafeMap) read(k int) {
	s.Lock()
	defer s.Unlock()
	i2 := 0
	i2 = s.sameMap[k]
	fmt.Println(i2)
}

func (s *SafeMap) write(k int, v int) {
	s.Lock()
	defer s.Unlock()
	s.sameMap[k] = v
}
