package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	boxes, err := parse(lines[:len(lines)-1])
	if err != nil {
		panic(err)
	}
	boxesCopy := make([]box, len(boxes))
	copy(boxesCopy, boxes)

	connections := computeDirectConnections(boxes)
	err = makeConnections(1000, boxes, connections)
	if err != nil {
		panic(err)
	}

	fmt.Println("result:", multiplyLongestCircuits(boxes))
	fmt.Println("result v2:", makeConnectionsV2(boxesCopy, connections))
}
