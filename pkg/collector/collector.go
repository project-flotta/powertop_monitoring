package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promauto"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"github.com/prometheus/common/version"
)

var (
	// custom collector
	reg = prometheus.NewRegistry()
	// some metrics
	myGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gauge_name",
			Help: "guage_help",
		},
		[]string{"l"},
	)
)

func init() {
	// register metrics to my collector
	reg.MustRegister(myGauge)
}
