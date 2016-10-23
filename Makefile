.PHONY: default
.DEFAULT_GOAL := test

VERSION := 1.0.2
NAME := github-protobuf
PKG := jhaynie/$(NAME)

SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})

default: all test;

clean:
	@rm -rf build


protoc-go:
	@mkdir -p build/$(VERSION)/go
	@cp src/*.proto build/$(VERSION)/go
	@docker run --rm -v $(BASEDIR):/app -v $(GOPATH):$(GOPATH) -w /app znly/protoc --gofast_out=build/$(VERSION)/go --proto_path=$(GOPATH)/src:$(GOPATH)/src/github.com/gogo/protobuf/protobuf:src src/*.proto

protoc-python:
	@mkdir -p build/$(VERSION)/python
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --python_out=build/$(VERSION)/python -Isrc src/*.proto

protoc-java:
	@mkdir -p build/$(VERSION)/java
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --java_out=build/$(VERSION)/java -Isrc src/*.proto

protoc-js:
	@mkdir -p build/$(VERSION)/javascript
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --js_out=build/$(VERSION)/javascript -Isrc src/*.proto

protoc-ruby:
	@mkdir -p build/$(VERSION)/ruby
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --ruby_out=build/$(VERSION)/ruby -Isrc src/*.proto

protoc-php:
	@mkdir -p build/$(VERSION)/php
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --php_out=build/$(VERSION)/php -Isrc src/*.proto

test: protoc-go
	@cp test/* build/$(VERSION)/go
	@cd build/$(VERSION)/go && go test -v *.pb.go *_test.go

all: protoc-go protoc-python protoc-java protoc-js protoc-ruby protoc-php

release: clean all
	@cd build && rm -rf *.tar *.gz && tar cvf $(NAME)-$(VERSION).tar * && gzip *.tar

default: all
