VERSION=$(shell git describe --abbrev=0 --tags)

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o "bin/server-$(VERSION)" main.go

debug:
	@echo "VERSION=$(VERSION)"

clean:
	rm -rf bin coverage.out
