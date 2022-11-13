VERSION=$(shell git describe --abbrev=0 --tags)
LOCALSTACK=LOCALSTACK_API_KEY=$$LOCALSTACK_API_KEY CONFIG_PROFILE=2022-11 localstack
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
	docker pull localstack/localstack
	$(LOCALSTACK) start --detached

.PHONY: localstack-stop
localstack-stop: $(HOME)/.localstack/2022-11.env
	$(LOCALSTACK) stop

.PHONY: docker-build
docker-build:
	docker build -t 2022-11 .
	docker tag 2022-11:latest localstack-ec2/alpine-ami:ami-000002
