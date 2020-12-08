#!/usr/bin/env bash
# https://hub.docker.com/r/suyanlong/golang-dev/
# https://github.com/suyanlong/golang-dev
# sudo docker pull suyanlong/golang-dev:latest

sudo docker run -it -p 8801:8801 -p 8802:8802 -p 6060:6060 -p 50051:50051 -l linux-dplatform-build \
    -v "$GOPATH"/src/gitlab.33.cn/dplatform/dplatform:/go/src/gitlab.33.cn/dplatform/dplatform \
    -w /go/src/gitlab.33.cn/dplatform/dplatform suyanlong/golang-dev:latest
