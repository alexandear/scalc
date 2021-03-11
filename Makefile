MAKEFILE_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
PATH := $(MAKEFILE_PATH):$(PATH)

export GOBIN := $(MAKEFILE_PATH)/bin
export GOFLAGS = -mod=vendor

PATH := $(GOBIN):$(PATH)

.PHONY: default
default: build lint test

.PHONY: build
build:
	@echo build
	@go build -o ./bin/scalc .

.PHONY: format
format:
	@echo format
	@go fmt ./...

.PHONY: vendor
vendor:
	@echo vendor
	@-rm -rf vendor/
	@go mod vendor

.PHONY: lint
lint:
	@echo lint
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@$(GOBIN)/golangci-lint run

.PHONY: test
test:
	@echo test
	@go test -race -v -count=1 ./...

IMAGE = scalc

.PHONY: docker
docker:
	@echo docker
	@docker build -t $(IMAGE) -f Dockerfile .
