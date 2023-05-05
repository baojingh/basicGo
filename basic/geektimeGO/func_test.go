package main

import (
	"fmt"
	"testing"
	"time"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 5:58
  @File     : func_test.go
  @Desc     :
*/

func slowFunc(param int) int {
	time.Sleep(time.Second * 2)
	return param
}

func TestFunc(t *testing.T) {
	sF := timeSpent(slowFunc)
	rr := sF(11)
	fmt.Print(rr)

}

func timeSpent(inner func(param int) int) func(p int) int {
	return func(p int) int {
		start := time.Now()
		ret := inner(p)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}

}

func Clear() {
	fmt.Println("clear")
}

func TestDefer(t *testing.T) {
	fmt.Println("start **********")
	//defer func() {
	//	fmt.Println("this is defer")
	//}()

	defer Clear()
	panic("this is panic")
	fmt.Println("done**********")

}
