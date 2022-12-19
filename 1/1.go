package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(filePath string) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var elf int = 1
	var elfTotal = 0
	for {
		var line int
		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			fmt.Printf("%d,%d\n", elf, elfTotal)
			if err == io.EOF {
				return
			}
			elf++
			elfTotal = 0
		} else {
			elfTotal += line
		}
	}
}

func main() {
	readFile("numbers.txt")
}
