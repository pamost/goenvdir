# Basic go commands
GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: gofmt test linter build

install:
	go build -o $(GOPATH)/bin/goenvdir

gofmt:
	gofmt -w .

test:
	go test -count=1 -race -cover -v ./...

linter:
	golangci-lint run --enable-all

build:
	go build -o goenvdir

clean:
	go clean -i ./...