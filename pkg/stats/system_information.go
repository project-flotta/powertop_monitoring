package stats

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	re_wakeup_info    = "System\\:\\s*(\\d{1,}\\.\\d{1,})\\s*wakeup\\/s"
	re_cpu_usage_info = "CPU\\:\\s*(\\d{1,}\\.\\d{1,})%\\s*usage"
)

type SysInfo struct {
	Wakeups  float64
	CpuUsage float64
}

var (
	system_data []string
	wake_up     string
	cpu_usage   string
)

func (sys_info SysInfo) ParseSysInfo(data [][]string) SysInfo {
	for _, line := range data {
		for _, element := range line {
			if strings.Contains(
				element,
				"System:",
			) {
				wake_up = element
				fmt.Println(element)
				continue
			}
			if strings.Contains(
				element,
				"CPU:",
			) {
				cpu_usage = element
				fmt.Println(element)
				break
			}
		}
	}
	reWakeUPInfo := regexp.MustCompile(re_wakeup_info)
	reCpuInfo := regexp.MustCompile(re_cpu_usage_info)

	matches := reWakeUPInfo.FindAllStringSubmatch(
		wake_up,
		-1,
	)
	matches2 := reCpuInfo.FindAllStringSubmatch(
		cpu_usage,
		-1,
	)
	if len(matches) != 0 {
		wake_up = matches[0][1]
	}
	if len(matches2) != 0 {
		cpu_usage = matches2[0][1]
	}
	sys_info.CpuUsage, _ = strconv.ParseFloat(
		wake_up,
		8,
	)
	sys_info.Wakeups, _ = strconv.ParseFloat(
		cpu_usage,
		8,
	)

	return sys_info

}

func (sys_info SysInfo) GetWakeUpData() float64 {
	return sys_info.Wakeups
}

func (sys_info SysInfo) GetCpuUsageData() float64 {
	return sys_info.CpuUsage
}
