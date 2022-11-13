VERSION=$(shell git describe --abbrev=0 --tags)
LOCALSTACK=CONFIG_PROFILE=2022-11 DOCKER_HOST=unix://$$XDG_RUNTIME_DIR/podman/podman.sock DOCKER_SOCK=$$XDG_RUNTIME_DIR/podman/podman.sock localstack
HOME=$(shell echo $$HOME)

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
	rm -rf $(HOME)/.localstack/2022-11.env

# Needed on Arch
# https://github.com/nektos/act/issues/303#issuecomment-962403508
.PHONY: act
act:
	act --container-daemon-socket $$XDG_RUNTIME_DIR/podman/podman.sock

$(HOME)/.localstack/2022-11.env:
	rm -rf $(HOME)/.localstack/2022-11.env
	mkdir -p $(HOME)/.localstack
	cp 2022-11.env $(HOME)/.localstack/2022-11.env

.PHONY: localstack-start
localstack-start: $(HOME)/.localstack/2022-11.env
	$(LOCALSTACK) start --detached

.PHONY: localstack-stop
localstack-stop: $(HOME)/.localstack/2022-11.env
	$(LOCALSTACK) stop

.PHONY: localstack-status
localstack-status: $(HOME)/.localstack/2022-11.env
	$(LOCALSTACK) status
