#!/bin/bash

LANG=$1
OUTCMD=$2
VERSION=$3
CMD=${4:-}

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

## Make all relative imports inside github package to absolute to allow easier importing into other packages
## Remove GoGo specific stuff for python
## Copy misc testing files

DIR=$BASEDIR/tmp

rm -rf $DIR
mkdir -p $DIR

function finish {
	rm -rf $DIR
}
trap finish EXIT

for f in `find src -name "*.proto"`; do
	tmpfile=$(mktemp /tmp/ghp.XXXXXX)
	tmpfile2=$(mktemp /tmp/ghp.XXXXXX)
	tmpfile3=$(mktemp /tmp/ghp.XXXXXX)
	tmpfile4=$(mktemp /tmp/ghp.XXXXXX)
	sed -e 's/import "github.com\/jhaynie\/go-github-protobuf\/github\/\([a-z]*\)\.proto"/import "\1.proto"/g' $f > $tmpfile
	sed -e 's/import "github.com\/gogo\/protobuf\/gogoproto\/gogo\.proto";//g' $tmpfile > $tmpfile2
	sed -e 's/option/\/\/options/g' $tmpfile2 > $tmpfile3
	sed -e 's/\[/;\/\/\[/g'  $tmpfile3 > $tmpfile4
	cp $tmpfile4 $DIR/$(basename $f)
	rm -rf $tmpfile $tmpfile2 $tmpfile3 $tmpfile4
done


mkdir -p build/$VERSION/$LANG
cp $DIR/*.proto build/$VERSION/$LANG
docker run --rm -v $BASEDIR:/app -w /app znly/protoc --$OUTCMD=${CMD}build/$VERSION/$LANG --proto_path=/app/vendor/:tmp tmp/*.proto
