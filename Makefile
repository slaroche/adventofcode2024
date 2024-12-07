GO_SOURCES := $(shell find . -type f -name '*.go')

.all: \
	hello

.DEFAULT_GOAL := .all

.PHONY: hello
hello: build/hello

build/hello: $(GO_SOURCES)
	go get ./...
	go build -o build/hello ./cmd/hello

.PHONY: lint
lint:
	go fmt ./...
	goimports -l -w .
	golint ./...
	go vet ./...

.PHONY: clean
clean:
	rm -rf build
