class Solution:
    def twoSum(self, numbers: list[int], target: int) -> list[int]:
        p1, p2 = 0, len(numbers) - 1

        while p1 < p2:
            currVal = numbers[p1] + numbers[p2]

            if currVal < target:
                p1 += 1
            elif currVal > target:
                p2 -= 1
            else:
                return [p1 + 1, p2 + 1]

        return []
