class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        seenMap = {}

        for i, val in enumerate(nums):
            currentTarget = target - val

            if currentTarget in seenMap:
                return [seenMap[currentTarget], i]

            seenMap[val] = i

        return [-1, -1]
