#! /usr/local/bin env bash

protoc -I .  -I $GOPATH/src/github.com/hyfco/oasis/proto/ --go_out=paths=source_relative:. conf.proto