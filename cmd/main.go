package main

import (
	"flag"
	"fmt"
	"github.com/project-flotta/powertop_container/pkg/container"
	"log"
	"strings"
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
	path = "/home/sibasishbehera/powertop_stats.csv"
)

var (
	address     = flag.String("address", "0.0.0.0:8886", "bind address")
	metricsPath = flag.String("metrics-path", "/metrics", "metrics path")
	sysInfo     stats.SysInfo
)

func main() {

	err := container.StartPowetopContainer()

	if err != nil {
		log.Printf("Error in starting the container %v", err)
	}

	flag.Parse()

	//register the collector
	err = prometheus.Register(version.NewCollector("powertop_tunable_exporter"))
	if err != nil {
		log.Fatalf("failed to register : %v", err)
	}

	if err != nil {
		log.Fatalf("failed to create collector: %v", err)
	}

	//prometheus http handler
	go func() {
		http.Handle(*metricsPath, promhttp.Handler())
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, err = w.Write([]byte(`<html>}
	fmt.Println("exporter call over")
}
			<head><title>PowerTop Tunable Exporter</title></head>
			<body>
			<h1>Tunable Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
			if err != nil {
				log.Fatalf("failed to write response: %v", err)
			}
		})

		err = http.ListenAndServe(*address, nil)
		if err != nil {
			log.Fatalf("failed to bind on %s: %v", *address, err)
		}
		fmt.Println("exporter call over")
	}()

	//Prometheus Metrics using Gauge
	pt_tu_count := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "powertop_tunables_count",
		Help: "counts the number of tuning available by powertop",
	})

	pt_wakeup_count := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "powertop_wakeup_count",
		Help: "counts the wake up calls per second available by powertop",
	})

	pt_cpu_usage_count := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "powertop_cpu_usage_count",
		Help: "counts the cpu usage in % by powertop",
	})

	pt_baseline_power_count := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "powertop_baseline_power_count",
		Help: "counts the baseline power used available by powertop",
	})

	for {
		data, err := stats.ReadCSV(path)
		if err != nil {
			log.Printf("Failed to fetch data : %v", err)
		}
		var t int
		t = 0
		//fmt.Println(data)
		for _, line := range data {
			fmt.Println(len(line))
			fmt.Println(line)
			t++
			if strings.Contains(line[0], " *  *  *   Device Power Report   *  *  *") {
				fmt.Println("hello")
				baseLinePower := data[t-3][len(data[t-2])-1]

				fmt.Println(baseLinePower)
				fmt.Println(len(baseLinePower))
			}

		}

		sysInfo = sysInfo.ParseSysInfo(data)

		////Fetch wakeup data
		wakeup_calls := sysInfo.Wakeups
		pt_wakeup_count.Set(wakeup_calls)
		fmt.Println("***********************************")
		fmt.Println(wakeup_calls)
		fmt.Println("***********************************")
		////Fetch cpuUsage data
		cpu_usage := sysInfo.CpuUsage
		pt_cpu_usage_count.Set(cpu_usage)
		fmt.Println("***********************************")
		fmt.Println(cpu_usage)
		fmt.Println("***********************************")
		////Fetch baseLine power
		baseLineData := stats.ParseBaseLinePower(data)
		baseLinePower := stats.GetBaseLinePOwer(baseLineData)
		pt_baseline_power_count.Set(baseLinePower)
		fmt.Println("***********************************")
		fmt.Println(baseLinePower)
		fmt.Println("***********************************")
		//Fetch no of tunables
		parsedData := stats.ParseTunables(data)
		tunNum, err := stats.GeNumOfTunables(parsedData)

		//Print logs
		stats.TunableLogs(parsedData)

		//Update the metric
		if err != nil {
			log.Printf("Error fetching no of Tunables %v", err)
		} else {
			pt_tu_count.Set(float64(tunNum))
			fmt.Println("***********************************")
			fmt.Println(tunNum)
			fmt.Println("***********************************")
		}

		//Sleeps for a hour
		time.Sleep(time.Second * 20)
	}

}
