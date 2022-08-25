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
	//gpuOperation float64
	//gfx          float64
	//vfs          float64
}

var (
	system_data []string
	wake_up     string
	cpu_usage   string
)

func (sys_info SysInfo) ParseSysInfo(data [][]string) SysInfo {
	//k := 0
	//for _, line := range data {
	//	k++
	//	if strings.Contains(
	//		line[0],
	//		" *  *  *   Top 10 Power Consumers   *  *  *",
	//	) {
	//		wake_up = data[k-3][1]
	//		cpu_usage = data[k-3][2]
	//		fmt.Println(wake_up)
	//		fmt.Println(cpu_usage)
	//		break
	//	}
	//}
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
	//fmt.Println("re starting +++++++++++++++++++++++++++++++++++++++++++++++++")
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

	//fmt.Println("&**&&*&*&&&*&&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*")
	//fmt.Println(wakeUpInfo)
	//fmt.Println(cpuInfo)
	//fmt.Println("&**&&*&*&&&*&&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*")
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

//func SetSysInfo(data []string) SysInfo {
//	sys_info.wakeups, _ = strconv.ParseFloat(data[0], 8)
//	sys_info.cpuUsage, _ = strconv.ParseFloat(data[0], 8)
//
//	return sys_info
//}

func (sys_info SysInfo) GetWakeUpData() float64 {
	//fmt.Println("##############################")
	//fmt.Println(sys_info.CpuUsage)
	//fmt.Println(sys_info.Wakeups)
	//fmt.Println("##############################")

	return sys_info.Wakeups
}
func (sys_info SysInfo) GetCpuUsageData() float64 {
	return sys_info.CpuUsage
}
