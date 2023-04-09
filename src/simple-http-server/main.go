package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
		fmt.Println("sent")
	})
	http.ListenAndServe(":8080", nil)
}
