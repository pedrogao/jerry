#!/usr/bin/env bash

echo building.....

echo may download some library, will cost a little minutes....

export GOPROXY=https://goproxy.io

go build -o jerry -tags=jsoniter jerry.go

echo build finish.

./jerry