package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"k1": 111,
		"k2": 222,
		"k3": 333,
		"k4": 444,
		"k5": 555,
	}

	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}

	fmt.Println("*************")
}
