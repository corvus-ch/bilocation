MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := test
.DELETE_ON_ERROR:
.SUFFIXES:

# ---------------------
# Environment variables
# ---------------------

GOPATH ?= $(shell go env GOPATH)

# ------------------
# Internal variables
# ------------------

binary_name = bilocation
cover_file  = c.out

# -------
# Targets
# -------

.PHONY: deps
deps:
	go mod vendor

.PHONY: test
test: $(cover_file)

$(cover_file): $(wildcard **/*_test.go)
	go test -coverpkg=./... -covermode=atomic -coverprofile=$@ ./...

.PHONY: build
build: $(binary_name)

$(binary_name): $(wildcard **/*.go)
	go build

.PHONY: protobuf
protobuf: tag/internal/bilocation.pb.go

tag/internal/bilocation.pb.go: bilocation.proto
	mkdir -p $(@D)
	protoc --go_out=$(@D) $<

.PHONY: clean
clean:
	rm -rf $(binary_name) $(cover_file)
