package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var answers map[string]int

type Operation struct {
	name1   string
	name2   string
	operand string
}

func (o Operation) Compute() (int, bool) {
	var value1, value2 int
	var found bool
	var answer int

	value1, found = answers[o.name1]
	if !found {
		return 0, false
	}
	value2, found = answers[o.name2]
	if !found {
		return 0, false
	}

	switch o.operand {
	case "+":
		answer = value1 + value2
	case "-":
		answer = value1 - value2
	case "*":
		answer = value1 * value2
	case "/":
		answer = int(value1 / value2)
	default:
		panic(fmt.Sprintf("Compute: invalid operation %s", o.operand))
	}

	return answer, true
}

var operations map[string]Operation

func LoadFile(filePath string) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	answers = make(map[string]int)
	operations = make(map[string]Operation)

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		var line string = scanner.Text()
		var items []string = strings.Split(line, ": ")
		var name string = items[0]
		var value int = 0
		if unicode.IsDigit(rune(items[1][0])) {
			value, _ = strconv.Atoi(items[1])
			answers[name] = value
			fmt.Printf("%s: %d\n", name, value)
		} else {
			var o Operation
			var name1, name2, operand string
			fmt.Sscanf(items[1], "%s %s %s", &name1, &operand, &name2)
			o.name1 = name1
			o.name2 = name2
			o.operand = operand
			operations[name] = o
		}
	}

	for len(operations) > 0 {

		for n, o := range operations {
			var value int
			var found bool

			value, found = o.Compute()
			if found {
				fmt.Printf("%s: %d\n", n, value)
				answers[n] = value
				delete(operations, n)
			}
		}
	}

	fmt.Printf("Done: %d\n", answers["root"])
}

func main() {
	// LoadFile("../sample.txt")
	LoadFile("../puzzle.txt")
}
