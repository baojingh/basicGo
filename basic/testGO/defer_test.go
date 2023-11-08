package main

import (
	"fmt"
	"testing"
)

/**
  Author     : He Bao Jing
  Date       : 11/8/2023 4:59 PM
  Description:
*/

func TestDeferDemo(t *testing.T) {
	println("before defer")
	defer deferFunc1()
	println("after defer")
	deferFunc2()
}

func deferFunc1() {
	fmt.Println("hello, this is deferFunc1")
}

func deferFunc2() {
	fmt.Println("hello, this is deferFunc2")
}
