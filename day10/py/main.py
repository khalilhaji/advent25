from z3 import *
import sys


def main():
    res = 0
    with open(sys.argv[1], 'r') as file:
        for line in file:
            buttons: list[list[int]] = []
            joltages = []
            parts = line.split(" ")
            for part in parts:
                if part[0] == '(':
                    buttons.append([int(i) for i in part[1:-1].split(",")])
                elif part[0] == '{':
                    joltages = [int(i) for i in part[1:-2].split(",")]

            vars = [Int(f"x_{i}") for i in range(len(buttons))]

            equations: list[list[ArithRef]] = [[] for _ in joltages]

            for i, button in enumerate(buttons):
                for idx in button:
                    equations[idx].append(vars[i])

            s = Optimize()
            for var in vars:
                s.add(var >= 0)

            for i, equation in enumerate(equations):
                s.add(Sum(*equation) == joltages[i])

            s.minimize(Sum(*vars))
            s.check()
            m = s.model()
            res += sum([m[x].as_long() for x in vars])

    print(res)
 


if __name__ == "__main__":
    main()
