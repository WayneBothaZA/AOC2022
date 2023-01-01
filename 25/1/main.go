package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func DecodeSNAFU(snafu byte) int {
	switch snafu {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '-':
		return -1
	case '=':
		return -2
	}
	return 0
}

func EncodeSNAFU(decimal int) byte {
	switch decimal {
	case -2:
		return '='
	case -1:
		return '-'
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	}
	// fmt.Printf("ERROR CONVERTING %d to SNAFU", decimal)
	return '?'
}

func Decimal2SNAFU(decimal int) (s []byte) {
	var i int
	var p float64

	// find the highest group
	for i = 0; ; i++ {
		p = math.Pow(5.0, float64(i))
		if int(p) > decimal {
			break
		}
	}

	var number []int = make([]int, i)

	// divide number into group, using remainder for rest of groups
	var v int = decimal
	for i = i - 1; i >= 0; i-- {
		p = math.Pow(5.0, float64(i))
		var a int = int(math.Floor(float64(v) / p))
		// fmt.Printf("(%d) %d: [%v]\n", i, int(p), a)
		number[i] = a
		v = int(math.Mod(float64(v), p))
	}

	// convert to -2 offset range
	for x := range number {
		switch number[x] {
		case 0:
		case 1:
		case 2:
		case 3:
			number[x] = -2
			if x == len(number)-1 {
				number = append(number, 0)
			}
			number[x+1] += 1
		case 4:
			number[x] = -1
			if x == len(number)-1 {
				number = append(number, 0)
			}
			number[x+1] += 1
		case 5:
			number[x] = 0
			if x == len(number)-1 {
				number = append(number, 0)
			}
			number[x+1] += 1
		}
	}

	for x := range number {
		s = append(s, EncodeSNAFU(number[len(number)-x-1]))
	}

	return s
}

func LoadNumbers(filePath string) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var total int = 0
	for scanner.Scan() {
		var line string = scanner.Text()
		// fmt.Printf("%s: ", line)

		var decimal int = 0
		for i, j := len(line)-1, 0; i >= 0; i, j = i-1, j+1 {
			decimal += DecodeSNAFU(line[i]) * int(math.Pow(5, float64(j)))
		}
		// fmt.Printf("%d\n", decimal)
		total += decimal
	}
	fmt.Printf("TOTAL: %d\n", total)
	fmt.Printf("%s\n", Decimal2SNAFU(total))
}

func main() {
	// LoadNumbers("../sample.txt")
	LoadNumbers("../puzzle.txt")
}
