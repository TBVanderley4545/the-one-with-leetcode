class Solution:
    def isValid(self, s: str) -> bool:
        checkMap = {"}": "{", "]": "[", ")": "("}
        stack = []

        for val in s:
            if val in checkMap:
                if not stack or stack.pop() != checkMap[val]:
                    return False
            else:
                stack.append(val)

        return not stack
