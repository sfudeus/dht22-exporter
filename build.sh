#!/usr/bin/env bash

docker buildx build --platform linux/arm/v7 -t sfudeus/dht22-exporter:latest --push .
