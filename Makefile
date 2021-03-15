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
	@go build -o ./bin/scalc ./cmd/...

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

.PHONY: generate
generate: mock

.PHONY: mock
mock:
	@echo mock
	@go install github.com/golang/mock/mockgen
	@go generate ./...

IMAGE = scalc

.PHONY: docker
docker:
	@echo docker
	@docker build -t $(IMAGE) -f Dockerfile .

.PHONY: example1
example1: build
	@-cp -r test/*.txt ./bin
	cd ./bin && scalc [ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]

.PHONY: example2
example2: build
	@-cp -r test/*.txt ./bin
	cd ./bin && scalc [ LE 2 a.txt [ GR 1 b.txt c.txt ] ]
