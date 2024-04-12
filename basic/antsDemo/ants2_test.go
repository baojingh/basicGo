package main

import (
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

/**
  @Author:      He Bao Jing
  @Date:        10/19/2023 2:44 PM
  @Description:
*/

func TestAnts2(t *testing.T) {
	var wg sync.WaitGroup

	pool, _ := ants.NewPool(10)
	defer pool.Release()

	for i := 1; i < 11; i++ {
		// 这可能是因为您的代码在并发环境中使用了闭包捕获了循环变量i。
		// 在Go语言中，循环变量i在每次迭代中会被重用，因此闭包捕获的实际上是同一个变量的引用，
		// 而不是它的值。当goroutine实际执行时，它们可能会看到不同的i值，
		// 因为i的值在循环迭代过程中会改变。
		currentI := i

		// wg.Add(1)被放在了goroutine内部，这是不正确的。wg.Add(1)应该在提交任务到池之前调用，
		// 以确保在goroutine实际开始运行之前增加等待组的计数器。否则，如果池中的goroutine在wg.
		// Add(1)调用之前就开始执行，那么wg.Wait()可能会过早地解除阻塞，因为计数器可能没有正确地增加。
		wg.Add(1)
		f := func() {
			defer wg.Done()
			task(currentI)
		}
		_ = pool.Submit(f)
		log.Infof("Running: %d, cap: %d", pool.Running(), pool.Cap())
	}
	wg.Wait()
	println("Task finished")

}

func task(no int) {
	log.Infof("task start: %d", no)
	time.Sleep(time.Second * time.Duration(1))
	log.Infof("task end: %d", no)
}
