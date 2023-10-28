class Solution:
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        occurenceMap = {}
        freq = [[] for i in range(len(nums) + 1)]

        for val in nums:
            occurenceMap[val] = occurenceMap.get(val, 0) + 1

        for key, val in occurenceMap.items():
            freq[val].append(key)

        res = []

        for i in range(len(freq) - 1, 0, -1):
            for value in freq[i]:
                res.append(value)
                if len(res) == k:
                    return res

        return res

    def topKFrequentUsingSorted(self, nums: list[int], k: int) -> list[int]:
        occurenceMap = {}

        for val in nums:
            occurenceMap[val] = occurenceMap.get(val, 0) + 1

        return sorted(occurenceMap, key=occurenceMap.get)[-k:]
