#!/bin/bash

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# make a temp directory to hold generated files
mkdir -p tmp/src/github.com/jhaynie/go-github-protobuf/github

# we need to make a temporary copy so that we can run our go program on the fly
cp build/*/go/*{pb.go,proto} tmp/src/github.com/jhaynie/go-github-protobuf/github
cp -R vendor/* tmp/src/

# generate a go program which will reflect our types from our proto types
GOPATH=$BASEDIR/tmp go run gencl.go > tmp/src/github.com/jhaynie/go-github-protobuf/github/gencl.go

# from these types, generate our json wrappers
GOPATH=$BASEDIR/tmp go run genjson.go $1

# cleanup
rm -rf tmp
