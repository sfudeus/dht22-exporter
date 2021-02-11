Prometheus Exporter for DHT22 sensors
====


```
Usage:
  dht22-exporter [OPTIONS]

Application Options:
      --port=      The address to listen on for HTTP requests. (default: 8080) [$EXPORTER_PORT]
      --interval=  The frequency in seconds in which to gather data (default: 60) [$INTERVAL]
      --pin=       The GPIO pin to use (default: 4)
      --metername= The name of your meter, to uniquely name them if you have multiple
      --debug      Activate debug mode

Help Options:
  -h, --help       Show this help message
```

Docker image
---

A recent docker image (multiarch amd64,arm/v7) is automatically built for each release at sfudeus/dht22-exporter:$TAG.
