#!/bin/sh

protoc -I ./rpc/ --go_out=./rpc/ --go-grpc_out=./rpc/ rpc/message/*.proto
cp -r rpc/authrpc clients/.