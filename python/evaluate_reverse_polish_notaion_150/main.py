class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        valStack = []

        for t in tokens:
            if t == "+":
                valStack.append(valStack.pop() + valStack.pop())
            elif t == "-":
                a, b = valStack.pop(), valStack.pop()

                valStack.append(b - a)
            elif t == "*":
                valStack.append(valStack.pop() * valStack.pop())
            elif t == "/":
                a, b = valStack.pop(), valStack.pop()

                valStack.append(int(float(b) / a))
            else:
                valStack.append(int(t))

        return valStack[0]
