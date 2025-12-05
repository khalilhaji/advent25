package main


func countAdjacentRolls(i, j int, d diagram) int {
	count := 0

	if j > 0 {
		left := d[i][j - 1]
		if left == SpacePaper {
			count++
		}
	}

	if j > 0 && i > 0 {
		topleft := d[i - 1][j - 1]
		if topleft == SpacePaper {
			count++
		}
	}

	if j > 0 && i < len(d) - 1 {
		bottomleft := d[i + 1][j - 1]
		if bottomleft == SpacePaper {
			count++
		}
	}

	if j < len(d[0]) - 1 {
		right := d[i][j + 1]
		if right == SpacePaper {
			count++
		}
	}

	if j < len(d[0]) - 1 && i > 0 {
		topright := d[i - 1][j + 1]
		if topright == SpacePaper {
			count++
		}
	}

	if j < len(d[0]) - 1 && i < len(d) - 1 {
		bottomright := d[i + 1][j + 1]
		if bottomright == SpacePaper {
			count++
		}
	}

	if i < len(d) - 1 {
		down := d[i + 1][j]
		if down == SpacePaper {
			count++
		}
	}

	if i > 0 {
		up := d[i - 1][j]
		if up == SpacePaper {
			count++
		}
	}

	return count
}

func traverse(d diagram) int {
	rolls := 0
	for i, y := range d {
		for j, x := range y {
			if x == SpacePaper {
				count := countAdjacentRolls(i, j, d)
				if count < 4 {
					rolls++
				}
			}
		}
	}

	return rolls
}

func traverseAndRemove(d diagram) int {
	removed := 0
	keepLooping := true

	for keepLooping {
		didRemove := false
		for i, y := range d {
			for j, x := range y {
				if x == SpacePaper {
					count := countAdjacentRolls(i, j, d)
					if count < 4 {
						didRemove = true
						removed++
						d[i][j] = SpaceEmpty
					}
				}
			}
		}
		keepLooping = didRemove
	}

	return removed
}

