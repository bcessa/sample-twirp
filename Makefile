.PHONY: proto
default: build
LD_FLAGS=""
BINARY_NAME=sample-twirp

proto: ## Compile protocol files
	@protoc \
	--proto_path=./vendor \
	--proto_path=. \
	--twirp_out=. \
	--go_out=. \
	--descriptor_set_out=proto/service.desc \
	proto/*.proto

build: ## Build for local system
	@go build -v -ldflags $(LD_FLAGS) -o $(BINARY_NAME)

linux: ## Build for AMD64 linux systems
	@env GOOS=linux GOARCH=amd64 go build -v -ldflags $(LD_FLAGS) -o $(BINARY_NAME)-linux-amd64

install: ## Install to local machine
	@go build -v -ldflags $(LD_FLAGS) -i -o ${GOPATH}/bin/$(BINARY_NAME)

deps-list: ## List installed vendor dependencies
	dep status

deps-install: ## Install required vendor dependencies
	dep ensure -v

help: ## Display available make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-16s\033[0m %s\n", $$1, $$2}'