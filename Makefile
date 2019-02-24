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

binary_name   = bilocation
cover_file    = c.out
grammar_dir   = search/internal/grammar
grammar_files = $(grammar_dir)/Path.interp $(grammar_dir)/Path.tokens $(grammar_dir)/PathLexer.interp $(grammar_dir)/PathLexer.tokens $(grammar_dir)/path_base_listener.go $(grammar_dir)/path_lexer.go $(grammar_dir)/path_listener.go $(grammar_dir)/path_parser.go

# -------
# Targets
# -------

.PHONY: deps
deps:
	go mod vendor

.PHONY: test
test: $(cover_file)

$(cover_file): $(wildcard **/*_test.go)
	go test -covermode=atomic -coverprofile=$@ ./...

.PHONY: build
build: $(binary_name)

$(binary_name): $(wildcard **/*.go)
	go build

.PHONY: protobuf
protobuf: tag/internal/bilocation.pb.go

tag/internal/bilocation.pb.go: bilocation.proto
	mkdir -p $(@D)
	protoc --go_out=$(@D) $<

.PHONY: grammar
grammar: $(grammar_files)

$(grammar_files): $(grammar_dir)/Path.g4
	cd $(<D)
	antlr -Dlanguage=Go -encoding UTF8 -package grammar $<

.PHONY: clean
clean:
	rm -rf $(binary_name) $(cover_file)
