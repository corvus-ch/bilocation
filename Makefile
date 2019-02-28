MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := test
.DELETE_ON_ERROR:
.SUFFIXES:

# ---------------------
# Environment variables
# ---------------------

GOPATH     ?= $(shell go env GOPATH)
GORELEASER ?= goreleaser

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

.PHONY: release
release:
	$(GORELEASER) --rm-dist
	package_cloud push corvus-ch/tools/ubuntu/xenial dist/bilocation_*_linux_*.deb
	package_cloud push corvus-ch/tools/ubuntu/bionic dist/bilocation_*_linux_*.deb
	package_cloud push corvus-ch/tools/debian/stretch dist/bilocation_*_linux_armv6.deb
	package_cloud push corvus-ch/tools/debian/buster dist/bilocation_*_linux_*.deb
	package_cloud push corvus-ch/tools/raspbian/stretch dist/bilocation_*_linux_armv6.deb
	package_cloud push corvus-ch/tools/raspbian/buster dist/bilocation_*_linux_armv6.deb
	package_cloud push corvus-ch/tools/el/6 dist/bilocation_*_linux_*.rpm
	package_cloud push corvus-ch/tools/el/7 dist/bilocation_*_linux_*.rpm
	package_cloud push corvus-ch/tools/fedora/28 dist/bilocation_*_linux_*.rpm
	package_cloud push corvus-ch/tools/fedora/29 dist/bilocation_*_linux_*.rpm

.PHONY: snapshot
snapshot:
	$(GORELEASER) --rm-dist --snapshot

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
