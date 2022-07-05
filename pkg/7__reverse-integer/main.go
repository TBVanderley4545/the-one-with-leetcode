package reverseinteger

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	multiplier := 1

	if x < 0 {
		multiplier = -1
	}

	inputAbs := x * multiplier

	intAsRunes := strings.Split(fmt.Sprint(inputAbs), "")

	reversedRunes := ""

	for _, val := range intAsRunes {
		reversedRunes = val + reversedRunes
	}

	reversedInteger, err := strconv.ParseInt(reversedRunes, 10, 32)

	if err != nil {
		return 0
	}

	return multiplier * int(reversedInteger)
}
