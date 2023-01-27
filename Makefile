.SILENT:

BINARY_NAME=apiserver

.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: run
run: build
	./${BINARY_NAME}
