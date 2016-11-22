.PHONY: default
.DEFAULT_GOAL := test

VERSION := 1.0.9
NAME := github-protobuf
PKG := jhaynie/$(NAME)

SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})

default: all test;

clean:
	@rm -rf build

protoc-go:
	@mkdir -p build/$(VERSION)/go
	@cp $(BASEDIR)/src/*.proto build/$(VERSION)/go
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --gofast_out=build/$(VERSION)/go --proto_path=/app/vendor/:src src/*.proto

protoc-python:
	@$(BASEDIR)/protoc.sh python python_out $(VERSION) ""

protoc-java:
	@$(BASEDIR)/protoc.sh java java_out $(VERSION) ""

protoc-js: protoc-go
	@rm -rf build/$(VERSION)/javascript
	@$(BASEDIR)/protoc.sh javascript js_out $(VERSION) "import_style=commonjs,binary:"
	@$(BASEDIR)/genjson.sh $(VERSION)

protoc-ruby:
	@$(BASEDIR)/protoc.sh ruby ruby_out $(VERSION) ""

protoc-php:
	@$(BASEDIR)/protoc.sh php php_out $(VERSION) ""

test: protoc-go
	@cp test/* build/$(VERSION)/go
	@cd build/$(VERSION)/go && go test -v *.pb.go *_test.go

all: protoc-go protoc-python protoc-java protoc-js protoc-ruby protoc-php

release: clean test clean all
	@cd build && rm -rf *.tar *.gz && tar cvf $(NAME)-$(VERSION).tar --exclude=*_test* * && gzip *.tar

default: all
