package main

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

func main() {

	reader, writer := io.Pipe()
	go func() {
		writer.Write([]byte("hello"))
		defer writer.Close()
	}()

	buffer := make([]byte, 20)
	reader.Read(buffer)

	log.Info(string(buffer))

	fmt.Println(string(buffer))

	log.Info("#############")

}
