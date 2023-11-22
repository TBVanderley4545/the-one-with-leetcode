from main import Solution


def test():
    assert Solution().evalRPN(["2", "1", "+", "3", "*"]) == 9, "The answer should be 9"
    assert Solution().evalRPN(["4", "13", "5", "/", "+"]) == 6, "The answer should be 6"
    assert (
        Solution().evalRPN(
            ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
        )
        == 22
    ), "The answer should be 22"


if __name__ == "__main__":
    test()
    print("Everything passed")
