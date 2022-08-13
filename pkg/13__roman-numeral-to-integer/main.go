package romannumeraltointeger

import "strings"

func RomanToInt(s string) int {
	numeral := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	splitString := strings.Split(s, "")

	pointer := 0
	total := 0

	for pointer < len(splitString) {
		currentVal := numeral[splitString[pointer]]
		nextVal := 0

		if pointer+1 < len(splitString) {
			nextVal = numeral[splitString[pointer+1]]
		}

		if currentVal < nextVal {
			total += nextVal - currentVal

			pointer += 2
		} else {
			total += currentVal

			pointer += 1
		}
	}

	return total
}
