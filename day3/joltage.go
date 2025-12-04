package main

import "math"

func sumJoltage(banks []bank, n int) int {
	total := 0
	for _, b := range banks {
		total += maxJoltage(b, n)
	}

	return total
}

func maxJoltage(b bank, n int) int {
	if len(b) < n {
		return 0
	}
	picks := []struct{
		idx int
		capacity int
	}{}

	for j := range n {
		minIdx := -1
		if len(picks) > 0 {
			minIdx = picks[len(picks) - 1].idx
		}
		curr := -1
		currIdx := -1

		for i := minIdx + 1; i <= len(b) - n + j; i++ {
			capacity := b[i]
			if capacity > curr {
				curr = capacity
				currIdx = i
			}
		}

		picks = append(picks, struct{idx int; capacity int}{currIdx, curr})
	}

	res := 0
	for i, pick := range picks {
		res += int(math.Pow10(len(picks) - i - 1)) * pick.capacity
	}

	return res
}

