FROM --platform=$BUILDPLATFORM golang:1.23 AS builder
ARG TARGETARCH
ARG TARGETOS
ARG TARGETPLATFORM

WORKDIR /build
ADD . /build

RUN ./docker_compile.sh

FROM scratch
COPY --from=builder /build/dht22-exporter /dht22-exporter
LABEL maintainer="Stephan Fudeus <github@mails.fudeus.net>"

ENTRYPOINT ["/dht22-exporter"]
EXPOSE 8080
