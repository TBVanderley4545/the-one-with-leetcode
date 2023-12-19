class Solution:
    def carFleet(self, target: int, position: list[int], speed: list[int]) -> int:
        pair = list(zip(position, speed))

        fleets = []

        for p, s in sorted(pair)[::-1]:
            fleets.append((target - p) / s)

            if len(fleets) >= 2 and fleets[-1] <= fleets[-2]:
                fleets.pop()

        return len(fleets)
