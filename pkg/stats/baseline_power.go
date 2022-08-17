package stats

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	reBaseLinePower = "The system baseline power is estimated at\\:\\s*(\\d{1,}\\.\\d{1,})\\s*W"
)

var (
	baseLinePower string
	reBLP         = regexp.MustCompile(reBaseLinePower)
)

// ParseBaseLinePower parse to get BaseLine data
func ParseBaseLinePower(data [][]string) string {
	k := 0
	for _, line := range data {
		k++
		if strings.Contains(line[0], " *  *  *   Device Power Report   *  *  *") {
			baseLinePower = data[k-3][len(data[k-2])-1]
		}
	}
	return baseLinePower

}
func GetBaseLinePOwer(parsedLine string) float64 {
	reBLP.Find([]byte(parsedLine))
	blp, err := strconv.ParseFloat(baseLinePower, 8)

	if err != nil {
		return 0.0
	}
	return blp
}
