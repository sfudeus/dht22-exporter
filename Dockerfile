FROM golang:alpine AS builder

ARG OS=linux
ARG ARCH=arm

WORKDIR /build
ADD . /build

RUN apk add --no-cache git
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -o dht22-exporter .

FROM scratch
COPY --from=builder /build/dht22-exporter /dht22-exporter
LABEL maintainer="Stephan Fudeus <github@mails.fudeus.net>"

ENTRYPOINT ["/dht22-exporter"]
EXPOSE 8080
