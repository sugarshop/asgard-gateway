#!/usr/bin/env bash

RUN_NAME="asgard.api"
mkdir -p output/bin output/conf
cp script/* output/
cp conf/* output/conf/
chmod +x output/bootstrap.sh

TAG='musl'
if [[ `uname` == 'Darwin' ]]; then
	TAG='dynamic'
fi

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -tags=$TAG,jsoniter -o output/bin/${RUN_NAME}
else
    go test -tags $TAG,jsoniter -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi
