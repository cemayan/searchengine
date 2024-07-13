#!/bin/bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install redis-sentinel bitnami/redis --values redis/values.yaml


#export REDIS_PASSWORD=$(kubectl get secret --namespace default redis-sentinel -o jsonpath="{.data.redis-password}" | base64 -d)
#kubectl run --namespace default redis-client --restart='Never'  --env REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:7.2.5-debian-12-r2 --command -- sleep infinity


