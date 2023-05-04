#!/bin/sh

export PSM=${PSM:-asgard.api}
CURDIR=$(cd $(dirname $0); pwd)
RUNTIME_ROOT=${CURDIR}
export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

exec "$CURDIR/bin/asgard.api" $*
