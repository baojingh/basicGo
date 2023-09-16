package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 5:08
  @File     : slice_test.go
  @Desc     :
*/

func TestSlice(t *testing.T) {
	var s0 []int

	fmt.Println(len(s0), cap(s0))

	s0 = append(s0, 111)
	fmt.Println(len(s0), cap(s0))

	s1 := make([]int, 3, 5)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 222)
	fmt.Println(len(s1), cap(s1))

	fmt.Println("***************")

	var s3 []int
	for i := 0; i < 20; i++ {
		s3 = append(s3, 1)
		fmt.Println(len(s3), cap(s3))
	}

}
