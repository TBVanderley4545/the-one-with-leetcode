package palindromenumber

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	remainder := x
	numArr := make([]int, 0)

	for remainder > 0 {
		numArr = append(numArr, remainder%10)

		remainder /= 10
	}

	left := 0
	right := len(numArr) - 1
	inverse := true

	for left <= right {
		if numArr[left] != numArr[right] {
			inverse = false
			break
		}

		left += 1
		right -= 1
	}

	return inverse
}
