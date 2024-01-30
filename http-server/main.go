package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ied", ied)
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		fmt.Println("http server startup failure, ", err)
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from http server!\n")
}

func ied(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		fmt.Println(string(body))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
