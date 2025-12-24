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
	nodes, err := parse(lines[:len(lines)-1])
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, node := range nodes {
		sum += pushes(node)
	}
	fmt.Println("total pushes:", sum)

	sumV2 := 0
	for _, node := range nodes {
		sumV2 += 
	}
}
