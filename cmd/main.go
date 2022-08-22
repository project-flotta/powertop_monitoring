package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/project-flotta/powertop_container/pkg/stats"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	//"log"
	"net/http"
)

const (
	path = "/var/tmp/powertop_report.csv"
)

var (
	address = flag.String(
		"address",
		"0.0.0.0:8886",
		"bind address",
	)
	metricsPath = flag.String(
		"metrics-path",
		"/metrics",
		"metrics path",
	)
	sysInfo stats.SysInfo
	//mountpoint string
	data [][]string
)

func main() {

	flag.Parse()

	//register the collector
	err := prometheus.Register(version.NewCollector("powertop_tunable_exporter"))
	if err != nil {
		log.Fatalf(
			"failed to register : %v",
			err,
		)
	}

	if err != nil {
		log.Fatalf(
			"failed to create collector: %v",
			err,
		)
	}

	//prometheus http handler
	go func() {
		http.Handle(
			*metricsPath,
			promhttp.Handler(),
		)
		http.HandleFunc(
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				_, err = w.Write(
					[]byte(`<html>}
	fmt.Println("exporter call over")
}
			<head><title>PowerTop Tunable Exporter</title></head>
			<body>
			<h1>Tunable Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`),
				)
				if err != nil {
					log.Fatalf(
						"failed to write response: %v",
						err,
					)
				}
			},
		)

		err = http.ListenAndServe(
			*address,
			nil,
		)
		if err != nil {
			log.Fatalf(
				"failed to bind on %s: %v",
				*address,
				err,
			)
		}
		fmt.Println("exporter call over")
	}()

	//Prometheus Metrics using Gauge
	pt_tu_count := promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "powertop_tunables_count",
			Help: "counts the number of tuning available by powertop",
		},
	)

	pt_wakeup_count := promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "powertop_wakeup_count",
			Help: "counts the wake up calls per second available by powertop",
		},
	)

	pt_cpu_usage_count := promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "powertop_cpu_usage_count",
			Help: "counts the cpu usage in % by powertop",
		},
	)

	pt_baseline_power_count := promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "powertop_baseline_power_count",
			Help: "counts the baseline power used available by powertop",
		},
	)

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println(
					"Tick at",
					t,
				)
				fmt.Println("command started")
				cmd := exec.Command(
					"powertop",
					"--csv="+path,
					"--time=4s",
				)
				cmd.Wait()
				out, err := cmd.Output()
				if err != nil {
					log.Printf(
						"%v",
						err,
					)
				}
				fmt.Printf(
					"%s",
					out,
				)
				data, err := stats.ReadCSV(path)
				if err != nil {
					log.Printf(
						"error in opening the csv file %v",
						err,
					)
				}

				// parse_csv_and_publish(path)
				sysInfo, baseLinePower, tunNum := ParseData(data)

				//publish
				////Fetch wakeup data
				pt_wakeup_count.Set(sysInfo.Wakeups)

				////Fetch cpuUsage data
				pt_cpu_usage_count.Set(sysInfo.CpuUsage)

				////Fetch baseLine power
				pt_baseline_power_count.Set(baseLinePower)

				//Fetch no of tunables
				pt_tu_count.Set(float64(tunNum))

			}
		}
	}()
	time.Sleep(20 * time.Second)
}

func ParseData(data [][]string) (stats.SysInfo, float64, uint32) {
	//parsing data
	sysInfo = sysInfo.ParseSysInfo(data)
	baseLineData := stats.ParseBaseLinePower(data)
	parsedTuned := stats.ParseTunables(data)
	tunNum := uint32(0)
	tunNum = stats.GeNumOfTunables(parsedTuned)
	//print tunable logs in console
	//stats.TunableLogs(parsedTuned)
	baseLinePower := stats.GetBaseLinePower(baseLineData)

	//fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	//fmt.Printf(
	//	"%v",
	//	sysInfo,
	//)
	//fmt.Println(baseLinePower)
	//fmt.Println(tunNum)
	//fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	return sysInfo, baseLinePower, tunNum
}
