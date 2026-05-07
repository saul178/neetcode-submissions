class Solution:
    def calPoints(self, operations: List[str]) -> int:
        scoreRecord = []

        for val in range(len(operations)):
            match operations[val]:
                case "+":
                    scoreRecord.append(scoreRecord[-1] + scoreRecord[-2])
                case "D":
                    scoreRecord.append(scoreRecord[-1] * 2)
                case "C":
                    if len(scoreRecord) != 0:
                        scoreRecord.pop()
                case _:
                    scoreRecord.append(int(operations[val]))
        return sum(scoreRecord)