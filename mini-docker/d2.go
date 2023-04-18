package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run() {
	cmd := exec.Command(os.Args[2])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("failed to exec the command ", err)
	}
}

func main() {
	arg := os.Args

	switch arg[1] {
	case "run":
		run()
	default:
		panic("the arg is not defined")
	}

	var a = 2
	fmt.Println(a)

	fmt.Println("##############")

}
