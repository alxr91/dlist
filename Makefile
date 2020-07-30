preseq:
	go get "github.com/c9s/goprocinfo/linux"

build:
	go build -o /usr/local/bin/superawesome main.go

run:
	superawesome

install: preseq build run
