.PHONY: build deps dev

build:
	go build -o server cmd/main.go

deps:
	go get github.com/prometheus/client_golang/prometheus/promhttp

dev: build
	./server