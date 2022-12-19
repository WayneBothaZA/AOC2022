package main

import (
	"fmt"
	"io"
	"os"
)

var size_x int = 0
var size_y int = 0
var forrest []string

var heights [][]int
var isVissible [][]bool

func PrintVissible() {
	for y, _ := range isVissible {
		for x, _ := range isVissible[y] {
			if isVissible[y][x] {
				fmt.Print("Y")
			} else {
				fmt.Print("N")
			}
		}
		fmt.Println()
	}
}

func PrintVissibleHeights() int {
	var total int = 0
	for y, _ := range isVissible {
		for x, _ := range isVissible[y] {
			if isVissible[y][x] {
				fmt.Print(heights[y][x])
				//total += heights[y][x]
				total++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return total
}

func ScanColumns() bool {
	var found bool = false

	fmt.Println("ScanColumns")
	for x := 0; x < size_x; x++ {
		var max int = heights[0][x]
		for y := 0; y < size_y; y++ {
			fmt.Printf("(%d,%d) %d %v ", x, y, heights[y][x], isVissible[y][x])
			if !isVissible[y][x] {
				if isVissible[y-1][x] && heights[y-1][x] < heights[y][x] {
					fmt.Printf("[%d (%d,%d)] ", heights[y-1][x], x, y-1)
					isVissible[y][x] = true
				} else if isVissible[y+1][x] && heights[y+1][x] < heights[y][x] {
					fmt.Printf("[%d (%d,%d)] ", heights[y+1][x], x, y+1)
					isVissible[y][x] = true
				} else if heights[y][x] > max {
					fmt.Printf("[%d Height] ", max)
					isVissible[y][x] = true
					max = heights[y][x]
				}
				if isVissible[y][x] {
					found = true
				}
			}
			fmt.Printf("%v\n", isVissible[y][x])
		}
		fmt.Println()
	}
	fmt.Println(found)
	return found
}

func ScanRows() bool {
	var found bool = false

	fmt.Println("\nScanRows")
	for y, _ := range isVissible {
		var max int = heights[y][0]
		for x, _ := range isVissible[y] {
			fmt.Printf("(%d,%d) %d %v ", x, y, heights[y][x], isVissible[y][x])
			if !isVissible[y][x] {
				if isVissible[y][x-1] && heights[y][x-1] < heights[y][x] {
					fmt.Printf("[%d (%d,%d)] ", heights[y][x-1], x-1, y)
					isVissible[y][x] = true
				} else if isVissible[y][x+1] && heights[y][x+1] < heights[y][x] {
					fmt.Printf("[%d (%d,%d)] ", heights[y][x+1], x+1, y)

					isVissible[y][x] = true
				}
				if heights[y][x] > max {
					fmt.Printf("[%d Height] ", max)
					isVissible[y][x] = true
					max = heights[y][x]
				}
				if isVissible[y][x] {
					found = true
				}
			}
			if heights[y][x] > max {
				if !isVissible[y][x] {
					fmt.Printf("[%d Height] ", max)
					isVissible[y][x] = true
					found = true
				}
				max = heights[y][x]
			}
			fmt.Printf("%v\n", isVissible[y][x])
		}
		fmt.Println()
	}
	fmt.Println(found)
	return found
}

func ScanUp(x, y int) bool {
	height := heights[y][x]
	for i := x - 1; i >= 0; i-- {
		if heights[y][i] >= height {
			return false
		}
	}
	return true
}
func ScanDown(x, y int) bool {
	height := heights[y][x]
	for i := x + 1; i < size_x; i++ {
		if heights[y][i] >= height {
			return false
		}
	}
	return true
}

func ScanLeft(x, y int) bool {
	height := heights[y][x]
	for i := y - 1; i >= 0; i-- {
		if heights[i][x] >= height {
			return false
		}
	}
	return true
}

func ScanRight(x, y int) bool {
	height := heights[y][x]
	for i := y + 1; i < size_y; i++ {
		if heights[i][x] >= height {
			return false
		}
	}
	return true
}

func ScanHeights() {
	for x := 1; x < size_x-1; x++ {
		for y := 1; y < size_y-1; y++ {
			fmt.Printf("(%d,%d):", x, y)
			if ScanUp(x, y) || ScanDown(x, y) || ScanLeft(x, y) || ScanRight(x, y) {
				// fmt.Print("X")
				isVissible[y][x] = true
			} else {
				// fmt.Print(" ")
			}
			fmt.Printf("\n")
		}
		fmt.Println()
	}
	fmt.Println()
}

func ProcessForrest() {
	heights = make([][]int, size_y)
	isVissible = make([][]bool, size_y)
	for y := range isVissible {
		heights[y] = make([]int, size_x)
		isVissible[y] = make([]bool, size_x)
	}

	for y, _ := range isVissible {
		for x, _ := range isVissible[y] {
			if (x == 0) || (y == 0) || (x == size_x-1) || (y == size_y-1) {
				isVissible[y][x] = true
			} else {
				isVissible[y][x] = false
			}
			heights[y][x] = int(forrest[y][x]) - '0'
		}
	}

	fmt.Println()

	ScanHeights()
	//PrintVissible()

	/*
		for {
			if !ScanRows() {
				break
			}
		}
		fmt.Println(PrintVissiblHeights())

		for {
			if !ScanColumns() {
				break
			}
		}
	*/

	//fmt.Println()
	//PrintVissible()

	fmt.Println(PrintVissibleHeights())
}

func LoadForrest(filePath string) {
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

		forrest = append(forrest, line)

		if size_x == 0 {
			size_x = len(line)
		}

		size_y++
	}
	fmt.Printf("(%d,%d)\n", size_x, size_y)

	for _, row := range forrest {
		fmt.Println(row)
	}
}

func main() {
	LoadForrest("../sample.txt")
	// LoadForrest("../puzzle.txt")
	ProcessForrest()
}
