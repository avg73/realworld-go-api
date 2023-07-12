run:
	go run ./cmd/apiserver

build:
	go build ./cmd/apiserver

.DEFAULT_GOAL := run
