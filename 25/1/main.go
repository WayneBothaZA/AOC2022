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

func Decimal2SNAFU(decimal int) string {
	fmt.Printf("\nConvert [%d]\n", decimal)

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
		fmt.Printf("(%d) %d: [%v]\n", i, int(p), a)
		number[i] = a
		v = int(math.Mod(float64(v), p))
	}

	// convert to -2 offset range
	for x := range number {
		fmt.Printf("%d", number[x])
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
		}
	}
	fmt.Printf("\n")

	fmt.Printf("Decimal values: ")
	for x := range number {
		fmt.Printf("%d ", number[len(number)-x-1])
	}
	fmt.Printf("\n")

	fmt.Printf("SNAFU values: ")
	for x := range number {
		fmt.Printf("%c", EncodeSNAFU(number[len(number)-x-1]))
	}
	fmt.Printf("\n")

	return ""
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
		fmt.Printf("%s: ", line)

		var decimal int = 0
		for i, j := len(line)-1, 0; i >= 0; i, j = i-1, j+1 {
			decimal += DecodeSNAFU(line[i]) * int(math.Pow(5, float64(j)))
			// fmt.Printf("%d * %d\n", DecodeSNAFU(line[i]), int(math.Pow(5, float64(j))))
		}
		fmt.Printf("%d\n", decimal)
		// Decimal2SNAFU(decimal)
		// fmt.Printf("\n")
		total += decimal
	}
	fmt.Printf("TOTAL: %d\n", total)
	// Decimal2SNAFU(total)

	// fmt.Printf("\nExpecting 2=0=  2 (125) -2 (25) 0 (5) -2 (1)\n\n")
	// Decimal2SNAFU(198)
	fmt.Printf("\n\n\n")
	Decimal2SNAFU(22)
	Decimal2SNAFU(23)
	Decimal2SNAFU(24)
	/*
			Decimal2SNAFU(1)
			Decimal2SNAFU(2)
			Decimal2SNAFU(3)
			Decimal2SNAFU(4)
			Decimal2SNAFU(5)
			Decimal2SNAFU(6)
			Decimal2SNAFU(7)
			Decimal2SNAFU(8)
			Decimal2SNAFU(9)
			Decimal2SNAFU(10)
			Decimal2SNAFU(15)
			Decimal2SNAFU(20)
			Decimal2SNAFU(2022)
			Decimal2SNAFU(12345)
			Decimal2SNAFU(314159265)

			Decimal2SNAFU(1747)
			Decimal2SNAFU(906)
			Decimal2SNAFU(198)
			Decimal2SNAFU(11)

			1=-0-2     1747
		 12111      906
		  2=0=      198
		    21       11
		  2=01      201
		   111       31
		 20012     1257
		   112       32
		 1=-1=      353
		  1-12      107
		    12        7
		    1=        3
		   122       37
	*/
}

func main() {
	LoadNumbers("../sample.txt")
	// LoadNumbers("../puzzle.txt")
}
