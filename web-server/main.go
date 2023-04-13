package main

import (
	"basicGo/web-server/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
