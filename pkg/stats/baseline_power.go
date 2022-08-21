package stats

import (
	"fmt"
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
	blp_value     float64
)

// ParseBaseLinePower parse to get BaseLine data
func ParseBaseLinePower(data [][]string) string {
	k := 0
	for _, line := range data {
		k++
		if strings.Contains(line[0], " *  *  *   Device Power Report   *  *  *") {
			baseLinePower = data[k-3][len(data[k-2])-1]
			fmt.Println("+++++++++++++++++++++++++++++++++++++")
			fmt.Println(baseLinePower)
			fmt.Println("+++++++++++++++++++++++++++++++++++++")
		}
	}
	return baseLinePower

}
func GetBaseLinePOwer(parsedLine string) float64 {
	matches := reBLP.FindAllStringSubmatch(parsedLine, -1)
	blp := matches[0][1]
	blp_value, _ = strconv.ParseFloat(blp, 8)
	return blp_value
}
