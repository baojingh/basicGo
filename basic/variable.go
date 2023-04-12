package main

import (
	"fmt"
)

func setup(task string) func() {
	fmt.Println("do the task ", task)
	return func() {
		println("do other task ", task)
	}
}

func main() {
	ss := setup("he")
	defer ss()

	fmt.Println("*************")
}
