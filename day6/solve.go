package main

func sumSolutions(problems []problem) int {
	res := 0
	for _, p := range problems {
		res += solve(p)
	}

	return res
}

func solve(p problem) int {
	var res int
	if p.o == OperatorMultiply {
		res = 1
	}

	for _, arg := range p.args {
		switch p.o {
		case OperatorMultiply:
			res *= arg
		case OperatorAdd:
			res += arg
		}
	}

	return res
}

func sumV2(problems []problemV2) int {
	res := 0
	for _, p := range problems {
		res += solveV2(p)
	}

	return res
}

func solveV2(p problemV2) int {
	finalArgs := make([]int, p.maxWidth)

	for _, argStr := range p.args {
		for j, digit := range argStr {
			if digit == ' ' {
				continue
			}

			finalArgs[j] *= 10

			finalArgs[j] += int(digit - '0')
		}
	}

	res := finalArgs[0]
	for _, arg := range finalArgs[1:] {
		switch p.o {
		case OperatorAdd:
			res += arg
		case OperatorMultiply:
			res *= arg
		}
	}

	return res
}
