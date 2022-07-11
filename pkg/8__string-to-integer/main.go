package stringtointeger

import (
	"math"
	"strings"
)

func MyAtoi(s string) int {
	trimmedStr := strings.Trim(s, " ")

	multiplier := 1.0
	finalLength := 0
	tally := 0.0
	offsetter := 0

	for idx, val := range trimmedStr {
		if idx == 0 && val == '-' {
			offsetter = 1
			multiplier = -1.0
			continue
		}

		if idx == 0 && val == '+' {
			offsetter = 1
			continue
		}

		ascii := int(val) - 48

		if ascii < 0 || ascii > 9 {
			break
		}

		tally += float64(ascii) * math.Pow(0.1, float64(idx+1-offsetter))
		finalLength += 1
	}

	finalVal := math.Round(tally * multiplier * math.Pow(10, float64(finalLength)))

	if finalVal > math.Pow(2, 31)-1 {
		finalVal = math.Pow(2, 31) - 1
	}

	if finalVal < math.Pow(-2, 31) {
		finalVal = math.Pow(-2, 31)
	}

	return int(finalVal)
}
