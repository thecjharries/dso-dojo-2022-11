VERSION=$(shell git describe --abbrev=0 --tags)
PACKER_FILES=dojo.pkr.hcl

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

.PHONY: act
act:
	act

.PHONY: packer
packer:
	packer init $(PACKER_FILES)
	packer fmt $(PACKER_FILES)
	packer validate $(PACKER_FILES)
	packer build $(PACKER_FILES)

.PHONY: dojo
dojo:
	$(MAKE) -f ./Makefile packer
	cd terraform && npm install && $(MAKE) -f ./Makefile test
