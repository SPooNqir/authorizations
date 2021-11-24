#!/bin/bash

set -e

GROUPS_VERSION="master"
curl -o proto/groups.proto "https://raw.githubusercontent.com/slavayssiere-spoon/groups/$GROUPS_VERSION/proto/groups.proto"

USERS_VERSION="master"
curl -o proto/users.proto "https://raw.githubusercontent.com/slavayssiere-spoon/users/$USERS_VERSION/proto/users.proto"

GEN_PATH="."
GO_LIB_PATH=$(go env GOPATH)/src
GOPATH=$(go env GOPATH)

echo "gen protoc"
protoc \
        -I proto \
        -I $GO_LIB_PATH/include \
        --go_out=$GEN_PATH      --go_opt=paths=source_relative \
        --go-grpc_out=$GEN_PATH --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=logtostderr=true,paths=source_relative:$GEN_PATH \
        --openapiv2_out=logtostderr=true:$GEN_PATH \
        proto/authorizations.proto

protoc-go-inject-tag -input=./authorizations.pb.go

echo "mod tidy"
go mod tidy

echo "update"
go get -u ./...



