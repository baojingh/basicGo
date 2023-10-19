package main

import (
	"fmt"
	"os"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 12:46
  @File     : hello_world.go
  @Desc     :
*/

func main() {
	fmt.Println(os.Args)

	fmt.Println("hello")
	os.Exit(-1)
}
