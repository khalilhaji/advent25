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

	instructions, err := Parse(lines[:len(lines) - 1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Count of end state = 0:", CountZeroEndStates(instructions))
	fmt.Println("Count of rotations past 0:", CountZeroPasses(instructions))
}
