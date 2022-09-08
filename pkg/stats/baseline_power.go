package stats

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	reBaseLinePower    = "The system baseline power is estimated at\\:\\s*(\\d{1,}\\.\\d{1,})\\s*W"
	reBatteryDischarge = "The battery reports a discharge rate of\\:\\s*(\\d{1,}\\.\\d{1,})\\s*W"
)

var (
	baseLinePower string
	reBLP         = regexp.MustCompile(reBaseLinePower)
	reBD          = regexp.MustCompile(reBatteryDischarge)
	blp_value     float64
	matches       [][]string
)

// ParseBaseLinePower parse to get BaseLine data
func ParseBaseLinePower(data [][]string) string {
	for _, line := range data {
		for _, element := range line {
			if strings.Contains(
				element,
				"discharge rate",
			) || (strings.Contains(
				element,
				"baseline power",
			)) {
				baseLinePower = element
				fmt.Println(element)
				return baseLinePower
			}

		}
	}
	return ""
}
func GetBaseLinePower(parsedLine string) float64 {
	matches := reBLP.FindAllStringSubmatch(
		parsedLine,
		-1,
	)
	if len(matches) == 0 {
		matches = reBD.FindAllStringSubmatch(
			parsedLine,
			-1,
		)
	}
	blp := matches[0][1]
	blp_value, _ = strconv.ParseFloat(
		blp,
		8,
	)
	return blp_value
}
