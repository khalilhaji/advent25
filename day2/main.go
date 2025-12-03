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

	ranges, err := parse(strings.TrimSpace(string(file)))
	if err != nil {
		panic(err)
	}

	fmt.Println("sum of invalid ids:", sumInvalid(ranges, invalidIDs))
	fmt.Println("sum of invalid ids (v2):", sumInvalid(ranges, invalidIDsV2))
}
