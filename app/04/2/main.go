package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var totalOverlaps int = 0

func Overlaps(pairs []string) bool {
	var min []int
	var max []int

	min = []int{0, 0}
	max = []int{0, 0}

	for i := 0; i < 2; i++ {
		minmax := strings.Split(pairs[i], "-")
		if len(minmax) != 2 {
			panic(fmt.Sprintf("Too many values - %s", pairs[i]))
		}
		min[i], _ = strconv.Atoi(minmax[0])
		max[i], _ = strconv.Atoi(minmax[1])
	}

	if (min[1] >= min[0]) && min[1] <= max[0] {
		return true
	}
	if (min[0] >= min[1]) && min[0] <= max[1] {
		return true
	}

	return false
}

func ProcessLine(line string) {
	var pairs []string = strings.Split(line, ",")

	if len(pairs) != 2 {
		panic(fmt.Sprintf("Too many pairs - %s", line))
	}

	if Overlaps(pairs) {
		totalOverlaps++
	}

	fmt.Printf("%s: %s %s, %d\n", line, pairs[0], pairs[1], totalOverlaps)
}

func ProcessFile(filePath string) {
	var line string

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	for {
		_, err := fmt.Fscanf(fd, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("fmt.Fscanf %s: %v", filePath, err))
		}

		ProcessLine(line)
	}
	fmt.Printf("%d\n", totalOverlaps)
}

func main() {
	ProcessFile("../sample.txt")
}
