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
	problems, err := parse(lines[:len(lines)-1])
	if err != nil {
		panic(err)
	}

	fmt.Println("sum of solutions:", sumSolutions(problems))

	v2Problems := parseV2(lines[:len(lines)-1])
	fmt.Println("sum v2:", sumV2(v2Problems))
}
