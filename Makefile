VERSION := $(shell git describe --tags --always --long --dirty)
PROJECT_NAME := $(shell basename "$(PWD)")
GO_SRC_FILES := $(shell find . -type f -name '*.go')

GO_ENVFLAGS=CGO_ENABLED=0

all: tidy vet fmt test build

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -l -w $(GO_SRC_FILES)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: build
build:
	go build -v -o $(PROJECT_NAME)
