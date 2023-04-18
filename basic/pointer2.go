package main

import (
	"fmt"
)

/**
  @Author   : bob
  @Datetime : 2023/4/17 23:12
  @File     : pointer2.go
  @Desc     :
*/

type calType func(a int, b int) int

func add(a int, b int) int {
	return a + b
}

func multi(a int, b int) int {
	return a * b
}

func cal(a int, b int, f calType) int {
	return f(a, b)
}

func main() {
	res := cal(1, 2, add)
	fmt.Println(res)

	fmt.Printf("############")
}
