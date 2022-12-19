package main

import (
	"fmt"
	"io"
	"os"
)

func FindBadge(group [3]string) byte {
	for i := 0; i < len(group[0]); i++ {
		for j := 0; j < len(group[1]); j++ {
			for k := 0; k < len(group[2]); k++ {
				if (group[0][i] == group[1][j]) &&
					(group[1][j] == group[2][k]) {
					return group[0][i]
				}
			}
		}
	}
	panic(fmt.Sprintf("FindCommonItem - not found"))
}

func ProcessGroup(group [3]string) byte {
	var badge byte
	for i := 0; i < 3; i++ {
		fmt.Printf("%d,%s\n", i, group[i])
	}
	badge = FindBadge(group)
	fmt.Printf("%c\n", badge)
	return badge
}

func readFile(filePath string) {
	var line int = 0
	var total int = 0
	var priority int = 0
	var group [3]string
	var badge byte

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	for {
		var rucksack string

		_, err := fmt.Fscanf(fd, "%s\n", &rucksack)

		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		group[line] = rucksack

		//fmt.Printf("%d,%s\n", line, rucksack)

		if line == 2 {
			badge = ProcessGroup(group)

			if badge >= 'a' {
				priority = int(badge-'a') + 1
			} else {
				priority = int(badge-'A') + 27
			}

			total += priority
			line = 0
		} else {
			line++
		}
	}
	fmt.Printf("Total: %d\n", total)
}

func main() {
	readFile("../rucksack.txt")
}
