all: clean build run

run:
	./http-server

build:
	GOOS=linux GOARCH=amd64 go build  -o http-server main.go

clean:
	rm -rf http-server

