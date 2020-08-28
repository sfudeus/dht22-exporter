package main // import "github.com/sfudeus/dht22-exporter"

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MichaelS11/go-dht"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sfudeus/dht22-exporter/internal/gathering"
)

var options struct {
	Port      int64  `long:"port" default:"8080" description:"The address to listen on for HTTP requests." env:"EXPORTER_PORT"`
	Interval  int64  `long:"interval" default:"60" env:"INTERVAL" description:"The frequency in seconds in which to gather data"`
	Pin       int64  `long:"pin" default:"4" description:"The GPIO pin to use"`
	MeterName string `long:"metername" description:"The name of your meter, to uniquely name them if you have multiple"`
	Debug     bool   `long:"debug" description:"Activate debug mode"`
}


func main() {
	_, err := flags.Parse(&options)
	if err != nil {
		os.Exit(1)
	}

	err = dht.HostInit()
	if err != nil {
		log.Println("HostInit error:", err)
		return
	}

	go func() {
		for {
			ok := gathering.GatherData(options.MeterName, options.Pin)
			if !ok {
				log.Printf("Data Gathering failed")
			}
			time.Sleep(time.Duration(options.Interval) * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", options.Port), nil))
}
