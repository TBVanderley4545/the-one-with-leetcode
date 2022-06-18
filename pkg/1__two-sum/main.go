package twosum

func TwoSum(numberSlice []int, target int) []int {
	lookupTable := make(map[int]int)

	for idx, num := range numberSlice {
		complement := target - num

		if _, ok := lookupTable[complement]; ok {
			return []int{lookupTable[complement], idx}
		}

		lookupTable[num] = idx
	}

	return make([]int, 2)
}
