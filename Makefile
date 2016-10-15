.PHONY: all clean

VERSION := 1.0.0
NAME := github-protobuf
PKG := jhaynie/$(NAME)

SHELL := /bin/bash
BASEDIR := $(shell echo $${PWD})

clean:
	@rm -rf build

protoc-go:
	@mkdir -p build/$(VERSION)/go
	@docker run --rm -v $(BASEDIR):/app -w /app znly/protoc --go_out=build/$(VERSION)/go -Isrc src/*.proto

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

all: protoc-go protoc-python protoc-java protoc-js protoc-ruby protoc-php

release: all
	@cd build && rm -rf *.tar *.gz && tar cvf $(NAME)-$(VERSION).tar * && gzip *.tar

default: all
