class Solution:
    def productExceptSelf(self, nums: list[int]) -> list[int]:
        output = [1] * len(nums)

        prefix = 1

        for i, num in enumerate(nums):
            output[i] = prefix
            prefix *= num

        postfix = 1

        for i in range(len(nums) - 1, -1, -1):
            output[i] *= postfix
            postfix *= nums[i]

        return output
