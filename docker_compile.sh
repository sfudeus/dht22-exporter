#!/usr/bin/bash

set -euo pipefail

echo "Building for target $TARGETPLATFORM"
IFS=/ read -a target <<<"$TARGETPLATFORM"

export CGO_ENABLED=0
export GOOS=${target[0]}
export GOARCH=${target[1]}
if [[ ${target[2]:-} =~ v[567] ]]; then
    export GOARM=${target[2]#v}
fi

go build cmd/dht22-exporter/dht22-exporter.go
