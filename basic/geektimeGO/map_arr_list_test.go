package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 4:01
  @File     : map_arr_list_test.go
  @Desc     :
*/

func TestArr(t *testing.T) {
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	//arr3 := [...]int{1, 2, 3, 4}
	arr4 := [...]int{1, 2, 5}

	fmt.Print(arr1 == arr2)
	//fmt.Print(arr1 == arr3)
	fmt.Print(arr1 == arr4)

}
