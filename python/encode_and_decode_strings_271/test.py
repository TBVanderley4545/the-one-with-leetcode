from main import Solution

inputArr = ["hi", "this", "is", "the", "input", "value", "for", "this", "function"]
encodedInputArr = "2#hi4#this2#is3#the5#input5#value3#for4#this8#function"


def test():
    assert (
        Solution().encode(inputArr) == encodedInputArr
    ), f"The encoded input should be {encodedInputArr}"
    assert (
        Solution().decode(encodedInputArr) == inputArr
    ), f"The decoded input should be {inputArr}"


if __name__ == "__main__":
    test()
    print("Everything passed")
