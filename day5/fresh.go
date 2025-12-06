package main

func countFresh(freshRanges []freshRange, avail []int) int {
	res := 0
	for _, a := range avail {
		for _, r := range freshRanges {
			if a >= r.low && a <= r.up {
				res++
				break
			}
		}
	}
	return res
}

func collapseRanges(freshRanges []*freshRange) []*freshRange {
	for i, fr := range freshRanges {
		for _, into := range freshRanges {
			if fr == into {
				continue
			}

			if into == nil {
				continue
			}

			if fr.low >= into.low && fr.low <= into.up {
				into.up = max(fr.up, into.up)
				freshRanges[i] = nil
			}

			if fr.up >= into.low && fr.up <= into.up {
				into.low = min(fr.low, into.low)
				freshRanges[i] = nil
			}
		}
	}

	return freshRanges
}

func countAllFresh(freshRanges []freshRange) int {
	freshRangesPointers := make([]*freshRange, len(freshRanges))
	for i, fr := range freshRanges {
		freshRangesPointers[i] = &fr
	}

	collapsed := collapseRanges(freshRangesPointers)
	res := 0
	for _, fr := range collapsed {
		if fr == nil {
			continue
		}
		res += fr.up - fr.low + 1
	}

	return res
}
