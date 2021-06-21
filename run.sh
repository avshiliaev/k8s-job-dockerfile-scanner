#!/usr/bin/env bash

kubectl delete job.batch/scanner-job

docker image build . -t avshiliaev/job:0.0.2
docker push avshiliaev/job:0.0.2

# docker-compose build --no-cache
# docker-compose up

kubectl apply -f ./k8s/job.yaml

kubectl get pods --watch
