GO_SOURCES := $(shell find . -type f -name '*.go')

.all: \
	aoc2024

.DEFAULT_GOAL := .all

.PHONY: aoc2024
pyrolite: build/aoc2024

build/aoc2024: $(GO_SOURCES)
	go get ./...
	go build -o build/aoc2024 ./cmd/aoc2024


.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...

.PHONY: clean
clean:
	rm -rf build
