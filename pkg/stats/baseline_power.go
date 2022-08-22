package stats

import (
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
	k := 0
	for _, line := range data {
		k++
		if strings.Contains(
			line[0],
			" *  *  *   Device Power Report   *  *  *",
		) {
			baseLinePower = data[k-3][0]
			if !strings.Contains(
				baseLinePower,
				"baseline",
			) {
				baseLinePower = data[k-4][0]
			}
			//fmt.Println("+++++++++++++++++++++++++++++++++++++")
			//fmt.Println(baseLinePower)
			//fmt.Println("+++++++++++++++++++++++++++++++++++++")
		}
	}
	return baseLinePower

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
	//fmt.Println("111111!!!!!!!!!!!!!1111!!!!!!!!!!!!!!!!!!!!111")
	//fmt.Println(blp_value)
	//fmt.Println("111111!!!!!!!!!!!!!1111!!!!!!!!!!!!!!!!!!!!111")
	return blp_value
}
