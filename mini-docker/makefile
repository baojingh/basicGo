CMD=go

all: clean build install

build:
	${CMD} build -o ./md  ./*.go

install:
	cp ./md /usr/local/bin

clean:
	rm -rf ./md
	rm -rf /usr/local/bin/md


