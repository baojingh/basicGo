package main

import (
	"fmt"
	"os"

	_ "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("this is init3...")
}

func init() {
	fmt.Println("this is init1...")
}

func init() {
	fmt.Println("this is init2...")
}

func main() {
	logrus.Println("hello, gopath mode")
	logrus.Println(uuid.NewString())
	fmt.Println("hello, ncv9uioasnio")

	fmt.Println("cavca", os.Getenv("GOPATH"))
}
