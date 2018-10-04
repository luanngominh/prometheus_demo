.PHONY: build deps dev

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

deps:
	go get github.com/prometheus/client_golang/prometheus/promhttp

dev: build
	./server