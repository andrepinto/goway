#!/usr/bin/env bash

protoc -I/usr/local/include -I. -I./vendor \
    -I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:. domain/*.proto

protoc -I/usr/local/include -I. -I./vendor \
    -I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. domain/*.proto

#protoc -I/usr/local/include -I. -I./vendor \
#    -I./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#     --swagger_out=logtostderr=true:. domain/*.proto
