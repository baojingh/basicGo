package main

import (
	"fmt"
	"testing"
)

/**
  @Author   : bob
  @Datetime : 4/5/2023 下午 5:27
  @File     : map_test.go
  @Desc     :
*/

func TestMap(t *testing.T) {
	var mm = map[string]int{"k1": 1, "k2": 2}
	for k, v := range mm {
		fmt.Printf("%s:%d\n", k, v)
	}

	var mm2 = map[string]int{"k1": 1}
	fmt.Println(mm2["k1"])
	if v, isExit := mm2["k1"]; isExit != true {
		fmt.Println("no found", v)
	} else {
		fmt.Println("existing")
	}

}

func TestMapFactory(t *testing.T) {
	m := map[string]func(param int) int{}
	m["a"] = func(param int) int {
		return 1
	}

	m["b"] = func(param int) int {
		return 2
	}

	m["c"] = func(param int) int {
		return 3
	}

	fmt.Println(m["a"](2))
}

func TestSet(t *testing.T) {
	set := map[string]bool{}
	set["ss"] = true

	fmt.Println()

}
