GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_ERROR_PROTO_FILES=$(shell find api -name "*error.proto")

.PHONY: init
# init env
init:
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/vektra/mockery/v2@latest

ifeq ($(GOHOSTOS), windows)
        #the `find.exe` is different from `find` in bash/shell.
        #to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
        #changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
        #Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
        Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git | grep cmd))))
        INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
        API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
        INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
        API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: api
# generate api、http、grpc、error、swagger
api:
	protoc --proto_path=./api \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
		   --go-grpc_out=paths=source_relative:./api \
		   --validate_out=lang=go,paths=source_relative:./api  \
		   --openapiv2_out=:./api \
		   --openapiv2_opt logtostderr=true \
		   --openapiv2_opt json_names_for_fields=false \
		   --go-errors_out=paths=source_relative:./api \
		  $(API_PROTO_FILES)

.PHONY: buf
# https://docs.buf.build/lint/usage#copy --error-format=config-ignore-yaml
buf:
	docker run --rm --volume "${PWD}:/workspace" --workdir /workspace/ bufbuild/buf lint

.PHONY: lint
# lint
lint:
	docker run --rm -v "${PWD}":/app -w /app golangci/golangci-lint:latest \
			sh -c "GOPROXY=https://goproxy.cn,direct GO111MODULE=on golangci-lint run"

.PHONY: lint-fix
# lint-fix
lint-fix:
	docker run --rm -v "${PWD}":/app -w /app golangci/golangci-lint:latest \
        sh -c "GOPROXY=https://goproxy.cn,direct GO111MODULE=on golangci-lint run --concurrency=2 --fix --timeout 1m"

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make buf;
	make generate;
	make api;

.PHONY: wire
# generate wire
wire:
	find app -maxdepth 2 -mindepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'


.PHONY: docker
docker:
	find app -maxdepth 2 -mindepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) docker'


.PHONY: test
test:
	find app -maxdepth 2 -mindepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) test'