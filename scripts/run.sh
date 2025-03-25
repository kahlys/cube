#!/bin/bash

# eval $(minikube docker-env)

go build -o server cmd/server/main.go
go build -o client cmd/client/main.go

if docker image inspect grpc-server:latest > /dev/null 2>&1; then
    docker rmi grpc-server:latest
fi
docker build -t grpc-server:latest -f cmd/server/Dockerfile .
if docker image inspect grpc-client:latest > /dev/null 2>&1; then
    docker rmi grpc-client:latest
fi
docker build -t grpc-client:latest -f cmd/client/Dockerfile .

rm server client