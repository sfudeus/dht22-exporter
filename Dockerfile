FROM golang:1.14 AS builder

WORKDIR /build
ADD . /build

RUN CGO_ENABLED=0 go build -o dht22-exporter .

FROM scratch
COPY --from=builder /build/dht22-exporter /dht22-exporter
LABEL maintainer="Stephan Fudeus <github@mails.fudeus.net>"

ENTRYPOINT ["/dht22-exporter"]
EXPOSE 8080
