class Solution:
    def trap(self, height: list[int]) -> int:
        if not height:
            return 0

        p1, p2 = 0, len(height) - 1
        maxL, maxR = height[p1], height[p2]

        total = 0

        while p1 < p2:
            if maxL < maxR:
                p1 += 1
                maxL = max(maxL, height[p1])
                total += maxL - height[p1]
            else:
                p2 -= 1
                maxR = max(maxR, height[p2])
                total += maxR - height[p2]

        return total
