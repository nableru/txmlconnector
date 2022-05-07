#!/bin/sh
# docker container rm -f transaq-exporter-1

# building txmlconnector client&server
# docker image rm -f nableru/txmlconnector:local
# docker build -f docker/Dockerfile.go_build_ubuntu -t nableru/txmlconnector:local .

# building transaq-clickhouse-exporter
docker build -f docker/Dockerfile.clickhouse-exporter -t nableru/transaq-clickhouse-exporter:local .

docker-compose -f docker/local-transaq-clickhouse-exporter-compose.yaml -p transaq up


# вход в любой контейнер 
# docker run --rm -it --entrypoint=/bin/bash kmlebedev/txmlconnector

# Clickhouse client
# docker container exec -it transaq-clickhouse-1 clickhouse-client

