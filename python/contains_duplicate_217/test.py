from main import Solution


def test():
    assert Solution().containsDuplicate([1, 2, 3, 4, 5]) == False, "Should be False"
    assert Solution().containsDuplicate([1, 2, 3, 1, 5]) == True, "Should be True"
    assert Solution().containsDuplicate([1, 2, 3, 1, 2]) == True, "Should be True"


if __name__ == "__main__":
    test()
    print("Everything passed")
