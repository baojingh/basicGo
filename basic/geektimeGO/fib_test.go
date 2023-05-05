package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 3:31
  @File     : fib_test.go
  @Desc     :
*/

func TestFib(t *testing.T) {
	var a int = 1
	var b int = 1
	for i := 0; i < 10; i++ {
		tmp := a
		a = b
		b = a + tmp
		t.Log(b, " ")
	}
}

func TestExchage(t *testing.T) {
	a := 1
	b := 2

	tmp := a
	a = b
	b = tmp

	fmt.Print(a, b)
}
