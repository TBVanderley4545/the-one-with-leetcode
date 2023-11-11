class Solution:
    def encode(self, strs: list[str]) -> str:
        res = ""

        for s in strs:
            res += f"{len(s)}#{s}"

        return res

    def decode(self, encodedVal: str) -> list[str]:
        res, p1 = [], 0

        while p1 < len(encodedVal):
            p2 = p1

            while encodedVal[p2] != "#":
                p2 += 1

            length = int(encodedVal[p1:p2])

            res.append(encodedVal[p2 + 1 : (p2 + 1 + length)])

            p1 = p2 + 1 + length

        return res
