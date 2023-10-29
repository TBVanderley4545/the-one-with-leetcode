from collections import defaultdict


class Solution:
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        colSet = defaultdict(set)
        rowSet = defaultdict(set)
        gridSet = defaultdict(set)

        for row in range(len(board[0])):
            for column in range(len(board)):
                if board[row][column] == ".":
                    continue

                currVal = board[row][column]

                if (
                    currVal in rowSet[row]
                    or currVal in colSet[column]
                    or currVal in gridSet[(row // 3, column // 3)]
                ):
                    return False

                rowSet[row].add(currVal)
                colSet[column].add(currVal)
                gridSet[(row // 3, column // 3)].add(currVal)

        return True
