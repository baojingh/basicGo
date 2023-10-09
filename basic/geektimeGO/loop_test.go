package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 4:38
  @File     : loop_test.go
  @Desc     :
*/

func TestFor(t *testing.T) {
	res := 0
	for i := 0; i < 101; i++ {
		res = res + i
	}
	//fmt.Print(res)

	res = 0
	n := 0
	for n < 5 {
		res = res + n
		n++
		fmt.Println(n)
	}

}

func TestSwitch(t *testing.T) {
	a := "ab"
	switch a {
	case "a", "ab":
		fmt.Println("this is a")
	case "b":
		fmt.Println("this is b")
	default:
		fmt.Println("this is default")
	}

}
