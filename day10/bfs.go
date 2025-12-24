package main

func pushes(root node) int {
	res := 0

	seen := map[string]bool{}
	seen[stateString(root.state)] = true

	currLevel := []node{root}

	for {
		nextLevel := []node{}
		for _, n := range currLevel {
			if stateEqual(n.state, n.target) {
				return res
			}
			for _, edge := range n.edges {
				child := node{
					state:   push(n.state, edge),
					joltage: setJoltage(n.joltage, edge),
					target:  n.target,
					edges:   n.edges,
					costs:   n.costs,
				}

				if _, ok := seen[stateString(child.state)]; ok {
					continue
				}
				seen[stateString(child.state)] = true
				nextLevel = append(nextLevel, child)
			}
		}

		currLevel = nextLevel
		res++
	}
}

func stateEqual(s0, s1 []state) bool {
	if len(s0) != len(s1) {
		return false
	}

	for i := range s0 {
		if s0[i] != s1[i] {
			return false
		}
	}

	return true
}

func joltageEqual(n node) bool {
	if len(n.joltage) != len(n.costs) {
		return false
	}

	for i := range n.joltage {
		if n.joltage[i] != n.costs[i] {
			return false
		}
	}

	return true
}

func push(s0 []state, button []int) []state {
	newState := make([]state, len(s0))

	copy(newState, s0)

	for _, idx := range button {
		newState[idx] = toggle(newState[idx])
	}

	return newState
}

func setJoltage(j []int, button []int) []int {
	newJoltage := make([]int, len(j))
	copy(newJoltage, j)

	for _, idx := range button {
		newJoltage[idx]++
	}
	return newJoltage
}

func toggle(s state) state {
	if s == On {
		return Off
	}
	return On
}

func stateString(s []state) string {
	out := make([]byte, len(s))

	for i, s0 := range s {
		out[i] = byte(s0)
	}

	return string(out)
}

func joltageString(j []int) string {
	out := make([]byte, len(j))

	for i, s0 := range j {
		out[i] = byte(s0)
	}

	return string(out)

}
