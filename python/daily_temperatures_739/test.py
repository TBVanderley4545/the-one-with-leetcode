from main import Solution


def test():
    assert Solution().dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]) == [
        1,
        1,
        4,
        2,
        1,
        1,
        0,
        0,
    ], "The first test case failed."

    assert Solution().dailyTemperatures([30, 40, 50, 60]) == [
        1,
        1,
        1,
        0,
    ], "The second test case failed."

    assert Solution().dailyTemperatures([30, 60, 90]) == [
        1,
        1,
        0,
    ], "The third test case failed."


if __name__ == "__main__":
    test()
    print("Everything passed")
