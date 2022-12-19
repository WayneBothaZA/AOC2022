package main

import (
	"fmt"
	"io"
	"os"
)

func PrintArray(array [][]byte) {
	for y := 0; y < len(array[0]); y++ {
		for x := 0; x < len(array); x++ {
			fmt.Printf("%d", array[x][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

func Load2DArray(filePath string) (array [][]byte) {
	var line string
	var lines []string
	var x, y int

	//var array [][]byte

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}

	// load all the lines
	x = 0
	y = 0
	for {
		_, err := fmt.Fscanf(fd, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("fmt.Fscanf %s: %v", filePath, err))
		}

		lines = append(lines, line)

		if x == 0 {
			x = len(line)
		}

		y++
	}

	// allocate the space
	array = make([][]byte, x)
	for i := range array {
		array[i] = make([]byte, y)
	}

	// flip the data
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			array[i][j] = lines[j][i] - '0'
		}
	}

	return array
}

func main() {
	// var array [][]byte = Load2DArray("../puzzle.txt")
	var array [][]byte = Load2DArray("../sample.txt")
	PrintArray(array)
}
