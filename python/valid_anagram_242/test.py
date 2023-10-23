from main import Solution


def test():
    assert Solution().isAnagram("anagram", "nagaram") == True, "Should be True"
    assert Solution().isAnagram("rat", "car") == False, "Should be False"


if __name__ == "__main__":
    test()
    print("Everything passed")
