#!/bin/sh
protoc --go_out=plugins=grpc:../ticket ./*.proto --proto_path=. --proto_path="$GOPATH/src/github.com/33cn/dplatformos/types/proto/"
