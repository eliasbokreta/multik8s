BIN	= multik8s
GO	= go

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
	$(GO) build

clean:
	$(info ▶ removing binary...)
	rm $(BIN)
