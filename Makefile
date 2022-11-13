VERSION=$(shell git describe --abbrev=0 --tags)
LOCALSTACK=DOCKER_HOST=unix://$$XDG_RUNTIME_DIR/podman/podman.sock DOCKER_SOCK=$$XDG_RUNTIME_DIR/podman/podman.sock localstack --profile=2022-11

.PHONY: all
all: test build

.PHONY: test
test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o "bin/server" main.go

.PHONY: debug
debug:
	@echo "VERSION=$(VERSION)"

.PHONY: clean
clean:
	rm -rf bin coverage.out

# Needed on Arch
# https://github.com/nektos/act/issues/303#issuecomment-962403508
.PHONY: act
act:
	act --container-daemon-socket $$XDG_RUNTIME_DIR/podman/podman.sock

.PHONY: localstack-start
localstack-start:
	rm -rf ~/.localstack/2022-11.env
	mkdir -p ~/.localstack
	cp 2022-11.env ~/.localstack/2022-11.env
	$(LOCALSTACK) start --detached

.PHONY: localstack-stop
localstack-stop:
	$(LOCALSTACK) stop
