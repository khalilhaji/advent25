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

	banks, err := parse(lines[:len(lines) - 1])
	if err != nil {
		panic(err)
	}

	fmt.Println("max joltage:", sumJoltage(banks, 2))
	fmt.Println("max joltage v2:", sumJoltage(banks, 12))
}
