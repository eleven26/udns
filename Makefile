.PHONY: init
init:
	go build -modfile=tools/go.mod -o bin/gofumpt mvdan.cc/gofumpt
	go build -modfile=tools/go.mod -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: check
check:
	bin/golangci-lint run

FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: format
format:
	go mod tidy
	bin/gofumpt -w $(FILES)

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: integration
integration:
	go test -v ./... -cover -tags=integration

.PHONY: build
build:
	go build -o udns

.PHONY: release
release:
	goreleaser release --snapshot --rm-dist

.PHONY: all
all:
	make check
	make format
	make test
	make integration