class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        output = [0] * len(temperatures)
        stack = []

        for idx, temp in enumerate(temperatures):
            while stack and temp > stack[-1][0]:
                _, stackIdx = stack.pop()
                output[stackIdx] = idx - stackIdx
            stack.append([temp, idx])

        return output
