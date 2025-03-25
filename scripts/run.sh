#!/bin/bash

# eval $(minikube docker-env)

docker build -t grpc-server:latest -f cmd/server/Dockerfile .