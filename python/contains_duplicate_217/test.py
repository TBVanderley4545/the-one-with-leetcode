from main import Solution


def test():
    assert Solution().containsDuplicate([1, 2, 3, 4, 5]) is False, "Should be False"
    assert Solution().containsDuplicate([1, 2, 3, 1, 5]) is True, "Should be True"
    assert Solution().containsDuplicate([1, 2, 3, 1, 2]) is True, "Should be True"


if __name__ == "__main__":
    test()
    print("Everything passed")
