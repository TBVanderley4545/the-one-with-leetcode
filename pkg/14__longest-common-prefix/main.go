package longestcommonprefix

import (
	"math"
)

func LongestCommonPrefix(strs []string) string {
	longestPrefix := ""
	shortest := math.MaxInt

	for i := 0; i >= 0; i++ {
		if i >= shortest || len(strs[0]) == 0 || len(longestPrefix) < i {
			break
		}

		targetLetter := string(strs[0][i])
		addLetter := true

		for _, word := range strs {
			if len(word) < shortest {
				shortest = len(word)
			}

			if shortest == 0 || string(word[i]) != targetLetter {
				addLetter = false
				break
			}
		}

		if addLetter {
			longestPrefix += targetLetter
		}
	}

	return longestPrefix
}
