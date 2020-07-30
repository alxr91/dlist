preseq:
	go get "github.com/c9s/goprocinfo/linux"
	go get "github.com/docker/docker/api/types"
	go get "github.com/docker/docker/client"

build:
	go build -o /usr/local/bin/dlist main.go

run:
	dlist

install: preseq build run
