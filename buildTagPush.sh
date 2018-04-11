#!/usr/bin/env bash
# docker login --username=shotcounterapp --email=shotcounterapp@gmail.com

docker build -t mongo-exporter -f Dockerfile .

docker tag mongo-exporter shotcounterapp/mongo-exporter
docker push shotcounterapp/mongo-exporter
