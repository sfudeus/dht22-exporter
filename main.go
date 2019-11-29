package main // import "github.com/sfudeus/dht22-exporter"

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MichaelS11/go-dht"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var options struct {
	Port      int64  `long:"port" default:"8080" description:"The address to listen on for HTTP requests." env:"EXPORTER_PORT"`
	Interval  int64  `long:"interval" default:"60" env:"INTERVAL" description:"The frequency in seconds in which to gather data"`
	Pin       int64  `long:"pin" default:"4" description:"The GPIO pin to use"`
	MeterName string `long:"metername" description:"The name of your meter, to uniquely name them if you have multiple"`
	Debug     bool   `long:"debug" description:"Activate debug mode"`
}

var (
	gaugeTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dht22",
		Name:      "temperature",
		Help:      "Current temperature",
	},
		[]string{
			//manual name of the meter, to distinguish between multiple sensors
			"meter_name",
		})
	gaugeHumidity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "dht22",
		Name:      "humidity",
		Help:      "Current humidity",
	},
		[]string{
			//manual name of the meter, to distinguish between multiple sensors
			"meter_name",
		})
	gatheringDuration = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "dht22",
		Name:      "gatheringduration",
		Help:      "The duration of data gatherings",
	},
		[]string{
			//manual name of the meter, to distinguish between multiple sensors
			"meter_name",
		})
)

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
			ok := gatherData()
			if !ok {
				log.Printf("Data Gathering failed")
			}
			time.Sleep(time.Duration(options.Interval) * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", options.Port), nil))
}

func gatherData() bool {
	timer := prometheus.NewTimer(gatheringDuration.WithLabelValues(options.MeterName))
	defer timer.ObserveDuration()

	log.Println("Gathering metrics")

	dht, err := dht.NewDHT(fmt.Sprintf("GPIO%d", options.Pin), dht.Celsius, "")
	if err != nil {
		log.Println("NewDHT error:", err)
		return false
	}

	humidity, temperature, err := dht.ReadRetry(5)
	if err != nil {
		log.Println("Read error:", err)
		return false
	}

	gaugeTemperature.WithLabelValues(options.MeterName).Set(temperature)
	gaugeHumidity.WithLabelValues(options.MeterName).Set(humidity)

	return true
}

func logDebug(format string, v ...interface{}) {
	if options.Debug {
		log.Printf(format, v...)
	}
}
