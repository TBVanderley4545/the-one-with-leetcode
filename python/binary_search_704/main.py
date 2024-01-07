class Solution:
    def search(self, nums: list[int], target: int) -> int:
        low, high = 0, len(nums) - 1

        while low <= high:
            midpoint = (low + high) // 2
            val = nums[midpoint]

            if val == target:
                return midpoint

            if val < target:
                low = midpoint + 1
            else:
                high = midpoint - 1

        return -1
