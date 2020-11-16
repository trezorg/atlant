# Makefile

PWD 				:= $(shell pwd)
BASE_DIR 			:= $(shell basename $(PWD))
GOPATH_DEFAULT 		:= $(PWD)/.go
export GOPATH 		?= $(GOPATH_DEFAULT)
GOBIN_DEFAULT 		:= $(GOPATH)/bin
export GOBIN 		?= $(GOBIN_DEFAULT)
export GO111MODULE 	:= on
TEST_ARGS_DEFAULT 	:= -v -race
TEST_ARGS 			?= $(TEST_ARGS_DEFAULT)
HAS_GOLANGCI 		:= $(shell command -v golangci-lint;)
HAS_GOIMPORTS 		:= $(shell command -v goimports;)
GOOS				?= $(shell go env GOOS)
DEFAULT_VERSION 	:= 0.0.1
VERSION				?= $(shell git describe --exact-match --tags 2>/dev/null || echo ${DEFAULT_VERSION})
GOARCH				:= amd64
CGO_ENABLED			:= 0
LDFLAGS				:= "-w -s -X 'main.Version=${VERSION}'"
SERVER_PACKAGE 		:= ./cmd/server
CLIENT_PACKAGE 		:= ./cmd/client
BINARY 				?= ./server
CLIENT_BINARY 		?= ./client
PROTOS 				:= $(shell find $(SOURCEDIR) -name '*.proto')

$(GOBIN):
	echo "create gobin"
	mkdir -p $(GOBIN)

work: $(GOBIN)

proto:
	@protoc --go_out=. \
 	--go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    $(PROTOS)

build: clean proto
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
	-ldflags $(LDFLAGS) \
	-o $(BINARY) \
	$(SERVER_PACKAGE)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
	-ldflags $(LDFLAGS) \
	-o $(CLIENT_BINARY) \
	$(CLIENT_PACKAGE)

start: build
	docker-compose up --build -d

stop:
	docker-compose down --remove-orphans -v

install: clean proto check test
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go install \
	-ldflags $(LDFLAGS) \
	$(SERVER_PACKAGE)

test: unit

check: work fmt vet goimports golangci

unit: work proto check
	go test -tags=unit $(TEST_ARGS) ./...

fmt:
	go fmt ./...

goimports:
ifndef HAS_GOIMPORTS
	echo "installing goimports"
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
endif
	goimports -d $(shell find . -iname "*.go")

vet:
	go vet ./...

golangci:
ifndef HAS_GOLANGCI
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.32.2
endif
	golangci-lint run ./...

cover: work
	go test $(TEST_ARGS) -tags=unit -cover -coverpkg=./ ./...

shell:
	$(SHELL) -i

clean: work
	@rm -rf $(BINARY)

version:
	@echo ${VERSION}

.PHONY: install build cover fmt test version clean check proto start
