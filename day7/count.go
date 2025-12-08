package main

func countSplits(d diagram) (int, map[int]int) {
	beams := map[int]int{
		d.start: 1,
	}
	splits := 0

	for _, line := range d.splitters {
		newBeams := map[int]int{}
		for beam, c := range beams {
			if ok := line[beam]; ok {
				splits++
				newBeams[beam+1] += c

				newBeams[beam-1] += c
			} else {
				newBeams[beam] += c
			}
		}
		beams = newBeams
	}

	return splits, beams
}

func countTimelines(beams map[int]int) int {
	res := 0
	for _, c := range beams {
		res += c
	}

	return res
}
