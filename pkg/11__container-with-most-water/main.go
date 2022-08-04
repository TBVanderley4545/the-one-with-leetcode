package containerwithmostwater

import "math"

func MaxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		currentArea := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))

		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return maxArea
}
