BIN		= multik8s
GO		= go
VERSION = $(shell git describe --tags --abbrev=0)

LDFLAGS = -ldflags="-X github.com/eliasbokreta/multik8s/cmd.version=$(VERSION)"

.PHONY: tidy fmt lint test build clean

default: all

all: tidy fmt lint test build

tidy:
	$(info ▶ cleaning dependencies...)
	$(GO) mod tidy

fmt:
	$(info ▶ formatting...)
	gofmt -s -w .

lint:
	$(info ▶ running lint tools...)
	golangci-lint run -v

test:
	$(info ▶ running tests...)
	$(GO) test -v -coverprofile cover.out ./...

build:
	$(info ▶ compiling program...)
	$(GO) build $(LDFLAGS)

clean:
	$(info ▶ removing binary...)
	rm $(BIN)
