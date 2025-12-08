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

	d := parse(lines[:len(lines)-1])
	splits, beams := countSplits(d)
	fmt.Println("splits:", splits)
	fmt.Println("timelines:", countTimelines(beams))
}
