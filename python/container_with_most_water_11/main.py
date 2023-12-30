class Solution:
    def maxArea(self, height: list[int]) -> int:
        p1, p2 = 0, len(height) - 1
        largest = 0

        while p1 < p2:
            smallest = min(height[p1], height[p2])
            gap = p2 - p1

            largest = max(largest, smallest * gap)

            if height[p1] > height[p2]:
                p2 -= 1
            else:
                p1 += 1

        return largest
