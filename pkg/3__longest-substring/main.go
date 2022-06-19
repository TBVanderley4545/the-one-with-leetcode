package longestsubstring

func LengthOfLongestSubstring(s string) int {
	runeMap := make(map[rune]int)

	left := 0
	longest := 0

	for idx, char := range s {
		if runeIdx, ok := runeMap[char]; ok && runeIdx > left-1 {
			left = runeIdx + 1
		}

		if currentLength := idx - left + 1; currentLength > longest {
			longest = currentLength
		}

		runeMap[char] = idx
	}

	return longest
}
