from main import Solution


def test():
    assert Solution().twoSum([2, 7, 11, 15], 9) == [0, 1], "Should be [0, 1]"
    assert Solution().twoSum([3, 2, 4], 6) == [1, 2], "Should be [1, 2]"
    assert Solution().twoSum([3, 3], 6) == [0, 1], "Should be [0, 1]"


if __name__ == "__main__":
    test()
    print("Everything passed")
