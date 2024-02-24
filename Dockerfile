FROM --platform=$BUILDPLATFORM golang:1.22 AS builder
ARG TARGETARCH
ARG TARGETOS

WORKDIR /build
ADD . /build

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build cmd/dht22-exporter/dht22-exporter.go

FROM scratch
COPY --from=builder /build/dht22-exporter /dht22-exporter
LABEL maintainer="Stephan Fudeus <github@mails.fudeus.net>"

ENTRYPOINT ["/dht22-exporter"]
EXPOSE 8080
