package main

import (
	"fmt"
	"io"
	"os"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Print() {
	for _, c := range *s {
		fmt.Printf("%s", string(c))
	}
}

var stacks int = 0
var stack [9]Stack

func PrintStacks() {
	for i := 0; i < stacks; i++ {
		fmt.Printf("%d: ", i)
		stack[i].Print()
		fmt.Printf("\n")
	}
}

func ProcessStacks(filePath string) {
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

		for _, c := range line {
			fmt.Printf("%s", string(c))
			stack[stacks].Push(string(c))
		}
		fmt.Printf(" %d\n", stacks)
		stacks++
	}
}

func ProcessMoves(filePath string) {
	var count, from, to int

	PrintStacks()
	fmt.Printf("Start\n")

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	for {
		_, err := fmt.Fscanf(fd, "%d %d %d\n", &count, &from, &to)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("fmt.Fscanf %s: %v", filePath, err))
		}

		from--
		to--
		fmt.Printf("move %d from %d to %d", count, from, to)

		fmt.Printf("\n----\n")
		PrintStacks()
		for i := 0; i < count; i++ {
			crate, _ := stack[from].Pop()
			stack[to].Push(crate)
		}
		PrintStacks()
		fmt.Printf("\n----\n")
	}

	for i := 0; i < stacks; i++ {
		crate, _ := stack[i].Pop()
		fmt.Printf("%s", crate)
	}
	fmt.Printf("\n")
}

func main() {
	ProcessStacks("../puzzle_stack.txt")
	ProcessMoves("../puzzle_moves.txt")
}
