package zigzagconversion

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	container := make([][]rune, numRows)

	currentRow := 0
	increment := 1

	for _, character := range s {
		container[currentRow] = append(container[currentRow], character)

		if currentRow+increment >= len(container) || currentRow+increment < 0 {
			increment *= -1
		}

		currentRow += increment
	}

	output := ""

	for _, val := range container {
		output += string(val)
	}

	return output
}
