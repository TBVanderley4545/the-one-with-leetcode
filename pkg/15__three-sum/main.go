package threesum

import "sort"

func ThreeSumSorted(nums []int) [][]int {
	sort.Ints(nums)

	output := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lowPointer := i + 1
		highPointer := len(nums) - 1

		for lowPointer < highPointer {
			sum := nums[i] + nums[lowPointer] + nums[highPointer]

			if sum < 0 {
				lowPointer++
			}

			if sum > 0 {
				highPointer--
			}

			if sum == 0 {
				output = append(output, []int{nums[i], nums[lowPointer], nums[highPointer]})

				lowPointer++
				highPointer--

				for lowPointer < len(nums) && nums[lowPointer] == nums[lowPointer-1] {
					lowPointer++
				}
			}
		}
	}

	return output
}
