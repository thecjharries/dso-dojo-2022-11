VERSION=$(shell git describe --abbrev=0 --tags)

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

debug:
	@echo "VERSION=$(VERSION)"
