#!/usr/bin/env bash

set -eu

if [ $# -lt 1 ]; then
  echo "Usage: $0 <go_version>";
  exit 1
fi

GO_VERSION=$1
PUSH=${2:-}

PRG_VERSION=$(git describe --tags --dirty)
IMAGE_VERSION=${PRG_VERSION}_${GO_VERSION}

echo "Building version $IMAGE_VERSION"
docker buildx build --build-arg "GO_VERSION=${GO_VERSION}" --platform linux/amd64 --platform linux/arm/v7 -t "sfudeus/dht22-exporter:${IMAGE_VERSION}" .

if [[ "${PUSH}" == "--push" ]]; then
  echo "Pushing version $IMAGE_VERSION"
  docker buildx build --build-arg "GO_VERSION=${GO_VERSION}" --platform linux/amd64 --platform linux/arm/v7 -t "sfudeus/dht22-exporter:${IMAGE_VERSION}" .
fi
