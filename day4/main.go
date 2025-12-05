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

	d, err := parse(lines[:len(lines)-1])
	if err != nil {
		panic(err)
	}

	fmt.Println("number of rolls:", traverse(d))
	fmt.Println("number of rolls with removing:", traverseAndRemove(d))
}
