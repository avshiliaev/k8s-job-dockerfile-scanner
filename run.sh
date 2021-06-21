#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -o scanner-job .
docker image build . -t avshiliaev/job:0.0.1
docker push avshiliaev/job:0.0.1

# docker-compose build --no-cache
# docker-compose up

kubectl apply -f ./k8s/job.yaml

kubectl get pods --watch
