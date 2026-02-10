package gathering

import (
	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sfudeus/go-dht"
)

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

// GatherData does the real data gathering and stores it into prometheus structures
func GatherData(meterName string, pin int64) bool {
	timer := prometheus.NewTimer(gatheringDuration.WithLabelValues(meterName))
	defer timer.ObserveDuration()

	log.Println("Gathering metrics")

	dht, err := dht.NewDHT(fmt.Sprintf("GPIO%d", pin), dht.Celsius, "")
	if err != nil {
		log.Println("NewDHT error:", err)
		return false
	}

	humidity, temperature, err := dht.ReadRetry(5)
	if err != nil {
		log.Println("Read error:", err)
		return false
	}

	gaugeTemperature.WithLabelValues(meterName).Set(temperature)
	gaugeHumidity.WithLabelValues(meterName).Set(humidity)

	return true
}
