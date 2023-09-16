package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 3:44
  @File     : constant_test.go
  @Desc     :
*/

func TestConstant(t *testing.T) {
	const (
		Monday  = 1
		Tuesday = 2
	)
	fmt.Print(Monday, Tuesday)

}
