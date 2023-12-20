class Solution:
    def isPalindrome(self, s: str) -> bool:
        normS = s.casefold()
        p1 = 0
        p2 = len(s) - 1

        while p1 < p2:
            p1Char = normS[p1]
            p2Char = normS[p2]

            if not p1Char.isalnum():
                p1 += 1
                continue

            if not p2Char.isalnum():
                p2 -= 1
                continue

            if p1Char != p2Char:
                return False

            p1 += 1
            p2 -= 1

        return True
