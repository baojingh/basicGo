package main

import (
	"fmt"
	"math"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 3:49
  @File     : type_test.go
  @Desc     :
*/

func TestType(t *testing.T) {
	fmt.Print(math.MaxFloat64)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	fmt.Printf("%T", a)
	fmt.Printf("%T", aPtr)
}

func TestString(t *testing.T) {
	var s string
	var a int = 1
	a++

	fmt.Print(a)
	fmt.Printf("%T", s)
	fmt.Println("*" + s + "*")

}
