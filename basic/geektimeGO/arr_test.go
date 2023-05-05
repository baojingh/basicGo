package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 4:49
  @File     : arr_test.go
  @Desc     :
*/

func TestArrDemo(t *testing.T) {
	arr1 := [...]int{0, 1, 2, 3}
	arr2 := [4]int{1, 2, 3}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for v, i := range arr2 {
		fmt.Println(v, i)
	}

	fmt.Println(arr1[1:2])

}
