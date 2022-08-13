package romannumeraltointeger

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

	total := 0

	for i := 0; i < len(s); i++ {
		currentVal := numeral[string(s[i])]
		nextVal := 0

		if i+1 < len(s) {
			nextVal = numeral[string(s[i+1])]
		}

		if currentVal < nextVal {
			total += nextVal - currentVal

			i++
		} else {
			total += currentVal
		}
	}

	return total
}
