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

	freshRanges, avail, err := parse(strings.Split(string(file), "\n"))
	if err != nil {
		panic(err)
	}

	fmt.Println("available fresh ingredients:", countFresh(freshRanges, avail))
	fmt.Println("total fresh ingredients:", countAllFresh(freshRanges))
}
