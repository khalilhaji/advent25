package main

func maxArea(coordinates []coordinate) (int, int) {
	maxArea := 0
	maxValidArea := 0

	for i, a := range coordinates {
		for j, b := range coordinates {
			if i == j {
				continue
			}

			currArea := area(a, b)
			if currArea > maxArea {
				maxArea = currArea
			}

			if isValid(a, b, coordinates) && currArea > maxValidArea {
				maxValidArea = currArea
			}

		}
	}

	return maxArea, maxValidArea
}

func isValid(a, b coordinate, coordinates []coordinate) bool {
	for i := 0; i < len(coordinates); i++ {
		first := coordinates[i]
		second := coordinates[(i+1)%len(coordinates)]

		colIntersects := false

		colRangeRectLow := min(a.col, b.col)
		colRangeRectHigh := max(a.col, b.col)

		colRangeLineLow := min(first.col, second.col)
		colRangeLineHigh := max(first.col, second.col)

		if colRangeRectHigh > colRangeLineLow &&
			colRangeRectLow < colRangeLineHigh {
			colIntersects = true
		}

		rowIntersects := false

		rowRangeRectLow := min(a.row, b.row)
		rowRangeRectHigh := max(a.row, b.row)

		rowRangeLineLow := min(first.row, second.row)
		rowRangeLineHigh := max(first.row, second.row)

		if rowRangeRectHigh > rowRangeLineLow &&
			rowRangeRectLow < rowRangeLineHigh {
			rowIntersects = true
		}

		if rowIntersects && colIntersects {
			return false
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func area(a, b coordinate) int {
	return (abs(a.col-b.col) + 1) * (abs(a.row-b.row) + 1)
}

type line struct{ a, b coordinate }
