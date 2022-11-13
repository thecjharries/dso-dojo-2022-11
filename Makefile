VERSION=$(shell git describe --abbrev=0 --tags)

.PHONY: all
all: test build

.PHONY: test
test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o "bin/server-$(VERSION)" main.go

.PHONY: debug
debug:
	@echo "VERSION=$(VERSION)"

.PHONY: clean
clean:
	rm -rf bin coverage.out
