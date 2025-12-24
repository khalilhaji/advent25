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

	coordinates, err := parse(lines[:len(lines)-1])
	if err != nil {
		panic(err)
	}

	m, m2 := maxArea(coordinates)
	fmt.Println("max area:", m)
	fmt.Println("max area part 2:", m2)
}
