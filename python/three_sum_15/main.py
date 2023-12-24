class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        result = []
        nums.sort()

        for i, val in enumerate(nums):
            if i > 0 and val == nums[i - 1]:
                continue

            p1, p2 = i + 1, len(nums) - 1

            while p1 < p2:
                threeSum = val + nums[p1] + nums[p2]

                if threeSum < 0:
                    p1 += 1
                    continue
                if threeSum > 0:
                    p2 -= 1
                    continue

                result.append([val, nums[p1], nums[p2]])

                p1 += 1

                while nums[p1] == nums[p1 - 1] and p1 < p2:
                    p1 += 1

        return result
