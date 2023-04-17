package main

import (
	"basicGo/web-server/router"
	"fmt"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(":8081")
	fmt.Println(err)
}
