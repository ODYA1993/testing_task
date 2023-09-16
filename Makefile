.PHONY: build
build:
	go build -v ./cmd

.PHONY: test
test:
	go test

.DEFAULT_GOAL := build