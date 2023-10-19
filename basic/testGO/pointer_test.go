package main

import (
	"fmt"
	"testing"
)

/**
  @Author:      He Bao Jing
  @Date:        10/19/2023 5:15 PM
  @Description:
*/

func TestPointer1(t *testing.T) {
	ty := new(int)
	println(ty)
	println(*ty)

	var val int
	val = 1
	println(&val)
	println(val)

	a := 3
	double(a)
	fmt.Printf("double value %d\n", a)

	b := 3
	p := &b
	double2(p)
	fmt.Printf("double2 value %d, %d\n", *p, b)
}

func double(x int) {
	x = x + x
}

func double2(x *int) {
	*x = *x + *x
}
