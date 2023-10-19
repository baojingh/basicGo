package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

/**
  @Author   : bob
  @Datetime : 5/5/2023 下午 2:00
  @File     : run_once_test.go
  @Desc     :
*/

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func getSingleton() *Singleton {
	once.Do(func() {
		fmt.Println("hello")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestRunOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := getSingleton()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
		}()
		wg.Done()
	}
	wg.Wait()
}
