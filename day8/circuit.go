package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func distance(a, b box) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

type directConnection struct {
	distance float64
	a, b     int
}

func computeDirectConnections(boxes []box) []directConnection {
	res := []directConnection{}
	for i, a := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			b := boxes[j]
			res = append(res, directConnection{
				distance: distance(a, b),
				a:        i,
				b:        j,
			})
		}
	}

	slices.SortFunc(res, func(a, b directConnection) int {
		if a.distance < b.distance {
			return -1
		}

		return 1
	})

	return res
}

func mergeCircuits(a, b int, boxes []box) {
	for i := range boxes {
		if boxes[i].circuit == a {
			boxes[i].circuit = b
		}
	}
}

func makeConnections(n int, boxes []box, connections []directConnection) error {
	if len(connections) < n {
		return fmt.Errorf("cannot make %v connections, not enough pairs", n)
	}

	for i := range n {
		c := connections[i]
		mergeCircuits(boxes[c.a].circuit, boxes[c.b].circuit, boxes)
	}

	return nil
}

func multiplyLongestCircuits(boxes []box) int {
	circuits := map[int]int{}
	for _, b := range boxes {
		circuits[b.circuit]++
	}

	keys := make([]int, 0, len(circuits))
	for k := range circuits {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return circuits[keys[i]] > circuits[keys[j]]
	})

	return circuits[keys[0]] * circuits[keys[1]] * circuits[keys[2]]
}

func countCircuits(boxes []box) int {
	c := map[int]bool{}
	for _, b := range boxes {
		c[b.circuit] = true
	}
	return len(c)
}

func makeConnectionsV2(boxes []box, connections []directConnection) int {
	var lastBoxes []box

	for i, n := 0, countCircuits(boxes); n > 1; i, n = i+1, countCircuits(boxes) {
		c := connections[i]
		mergeCircuits(boxes[c.a].circuit, boxes[c.b].circuit, boxes)
		lastBoxes = []box{boxes[c.a], boxes[c.b]}
	}

	return lastBoxes[0].x * lastBoxes[1].x
}
