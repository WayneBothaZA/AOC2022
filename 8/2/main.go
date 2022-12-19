package main

import (
	"fmt"
	"io"
	"os"
)

func ScanLeft(x, y int, array [][]byte) int {
	vissible := 0
	for i := x - 1; i >= 0; i-- {
		vissible++
		if array[i][y] >= array[x][y] {
			break
		}
	}
	return vissible
}

func ScanRight(x, y int, array [][]byte) int {
	vissible := 0
	for i := x + 1; i < len(array); i++ {
		vissible++
		if array[i][y] >= array[x][y] {
			break
		}
	}
	return vissible
}

func ScanUp(x, y int, array [][]byte) int {
	vissible := 0
	for j := y - 1; j >= 0; j-- {
		vissible++
		if array[x][j] >= array[x][y] {
			break
		}
	}
	return vissible
}

func ScanDown(x, y int, array [][]byte) int {
	vissible := 0
	for j := y + 1; j < len(array[0]); j++ {
		vissible++
		if array[x][j] >= array[x][y] {
			break
		}
	}
	return vissible
}

func ProcessArray(array [][]byte) {
	max := 0
	vissible := 0
	max_x := len(array) - 1
	max_y := len(array[0]) - 1
	for y := 0; y <= max_y; y++ {
		for x := 0; x <= max_x; x++ {
			vissible = ScanLeft(x, y, array) * ScanRight(x, y, array) * ScanUp(x, y, array) * ScanDown(x, y, array)
			fmt.Printf("%d", vissible)
			if vissible > max {
				max = vissible
			}
		}
		fmt.Println()
	}
	fmt.Printf("Max: %d\n", max)
}

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
	var array [][]byte = Load2DArray("../puzzle.txt")
	//var array [][]byte = Load2DArray("../sample.txt")
	PrintArray(array)
	ProcessArray(array)
}
