package mediansortedarrays

import (
	"math"
)

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	combinedLength := len(nums1) + len(nums2)

	if combinedLength <= 2 {
		sum := 0.0

		for _, val := range append(nums1, nums2...) {
			sum += float64(val)
		}

		return sum / float64(combinedLength)
	}

	midpoint := float64(combinedLength+1) / 2
	lowerMidpoint := int(math.Floor(midpoint))
	upperMidpoint := int(math.Ceil(midpoint))

	lookupTable := make(map[int]int)
	lookupTable[lowerMidpoint] = 0
	lookupTable[upperMidpoint] = 0
	lookupTableSum := 0.0

	pointer1 := 0
	pointer2 := 0

	for i := 0; i < int(upperMidpoint); i++ {
		smallerNum := 0
		pointerIncremented := false

		if pointer1 >= len(nums1) {
			smallerNum = nums2[pointer2]
			pointer2++
			pointerIncremented = true
		}

		if pointer2 >= len(nums2) {
			smallerNum = nums1[pointer1]
			pointer1++
			pointerIncremented = true
		}

		if !pointerIncremented {
			smallerNum = int(math.Min(float64(nums1[pointer1]), float64(nums2[pointer2])))

			if int(smallerNum) == nums1[pointer1] {
				pointer1++
			} else {
				pointer2++
			}
		}

		if _, ok := lookupTable[i+1]; ok {
			lookupTable[i+1] = smallerNum
		}

	}

	for _, val := range lookupTable {
		lookupTableSum += float64(val)
	}

	return lookupTableSum / float64(len(lookupTable))
}
