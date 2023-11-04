package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	command string
	value   int
}

func (i Instruction) isNoop() bool {
	return i.command == "noop"
}

func (i Instruction) String() string {
	if i.isNoop() {
		return fmt.Sprintf("%s", i.command)
	} else {
		return fmt.Sprintf("%s %d", i.command, i.value)
	}
}

func (i Instruction) Cycles() int {
	if i.isNoop() {
		return 1
	} else {
		return 2
	}
}

func ProcessInstructions(instructions []Instruction) {
	var cycle int = 0
	var x int = 1
	var values []int

	for _, instruction := range instructions {
		// fmt.Printf("%d %s (%d)\n", cycle, instruction, x)
		cycle += instruction.Cycles()

		values = append(values, x)
		if !instruction.isNoop() {
			values = append(values, x)
		}

		x += instruction.value
	}

	var pixel int = 0
	for _, v := range values {
		//fmt.Printf("Cycle %d: Pixel: %d, X: %d - Sprite position %d,%d,%d ", cycle+1, pixel, v, v-1, v, v+1)
		if pixel == v-1 || pixel == v || pixel == v+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		// fmt.Println()
		if pixel == 39 {
			fmt.Println()
			pixel = 0
		} else {
			pixel++
		}
	}

	// fmt.Printf("%d\n", values[20]*20+values[60]*60+values[100]*100+values[140]*140+values[180]*180+values[220]*220)
}

func LoadInstructions(filePath string) (instructions []Instruction) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		var instruction Instruction
		var line string = scanner.Text()
		var words []string = strings.Split(line, " ")
		if words[0] == "noop" {
			instruction.command = words[0]
			instruction.value = 0
		} else {
			instruction.command = words[0]
			instruction.value, _ = strconv.Atoi(words[1])
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func main() {
	instructions := LoadInstructions("../puzzle.txt")
	ProcessInstructions(instructions)
}
