package main

import (
	"fmt"
	"io"
	"os"
)

func FindCommonItem(c1, c2 string) byte {
	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {
			if c1[i] == c2[j] {
				return c1[i]
			}
		}
	}
	panic(fmt.Sprintf("FinCommonItem - not found"))
}

func readFile(filePath string) {
	var total int = 0

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	for {
		var rucksack string
		var priority int

		_, err := fmt.Fscanf(fd, "%s\n", &rucksack)

		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		c1 := rucksack[:len(rucksack)/2]
		c2 := rucksack[len(rucksack)/2:]

		commonItem := FindCommonItem(c1, c2)

		if commonItem >= 'a' {
			priority = int(commonItem-'a') + 1
		} else {
			priority = int(commonItem-'A') + 27
		}
		total += priority

		fmt.Printf("%s,%s,%s,%d(%c),%d\n", rucksack, c1, c2, priority, commonItem, total)
	}
}

func main() {
	readFile("../sample.txt")
}
