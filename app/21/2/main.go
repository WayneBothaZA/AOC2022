package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var humn int = 1
var answers map[string]int

type Operation struct {
	name1   string
	operand string
	name2   string
}

func (o Operation) Compute() (int, bool) {
	var value1, value2 int
	var found bool
	var answer int

	if unicode.IsDigit(rune(o.name1[0])) {
		value1, _ = strconv.Atoi(o.name1)
	} else {
		value1, found = answers[o.name1]
		if !found {
			return 0, false
		}
	}
	if unicode.IsDigit(rune(o.name2[0])) {
		value2, _ = strconv.Atoi(o.name2)
	} else {
		value2, found = answers[o.name2]
		if !found {
			return 0, false
		}
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
	case "==":
		return value1 - value2, value1 == value2
	default:
		panic(fmt.Sprintf("Compute: invalid operation %s", o.operand))
	}

	return answer, true
}

func (o *Operation) FlipOperand() {
	switch o.operand {
	case "+":
		o.operand = "-"
	case "-":
		o.operand = "+"
	case "*":
		o.operand = "/"
	case "/":
		o.operand = "*"
	}
}

var operations map[string]Operation

func RewriteOperation(lhs string) {
	var o Operation = operations[lhs]
	var n string

	fmt.Printf("%s = %s\t", lhs, o)

	o.FlipOperand()

	if unicode.IsDigit(rune(o.name1[0])) {
		n = o.name2
		o.name2 = o.name1
		o.name1 = lhs
	} else if unicode.IsDigit(rune(o.name2[0])) {
		n = o.name1
		o.name1 = lhs
	}

	fmt.Printf("%s = %s\t\t", n, o)

	a, f := o.Compute()
	if f {
		fmt.Printf("%s: %d\n", n, a)
		answers[n] = a
	} else {
		fmt.Printf("%s: NaN - ", n)
		v1, f1 := answers[o.name1]
		v2, f2 := answers[o.name2]
		fmt.Printf("%v (%v), %v (%v)\n", v1, f1, v2, f2)
	}
}

func ReorderOperation(n string) {
	var next string
	o := operations[n]
	// fmt.Printf("Reorder %s = %s\n", n, o)

	if unicode.IsDigit(rune(o.name1[0])) {
		next = o.name2
	} else if unicode.IsDigit(rune(o.name2[0])) {
		next = o.name1
	} else {
		fmt.Printf("Oops\n")
		return
	}

	RewriteOperation(n)
	if next != "humn" {
		// humn is the last one, we are done
		ReorderOperation(next)
	}

	// fmt.Printf("Next: %s %s\n", next, operations[next])
}

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
			// fmt.Printf("%s: %d\n", name, value)
		} else {
			var o Operation
			var name1, name2, operand string
			fmt.Sscanf(items[1], "%s %s %s", &name1, &operand, &name2)

			if name == "root" {
				operand = "=="
			}
			o.name1 = name1
			o.name2 = name2
			o.operand = operand
			operations[name] = o
		}
	}

	delete(answers, "humn")

	for len(operations) > 0 {

		var oLeft int = len(operations)
		for n, o := range operations {
			var value int
			var found bool

			value, found = o.Compute()
			if found {
				// fmt.Printf("%s: %d\n", n, value)
				answers[n] = value
				delete(operations, n)
			}

		}
		if oLeft == len(operations) {
			fmt.Printf("%d operations left - %v\n", len(operations), operations)
			break
		}
	}

	a1, f1 := answers[operations["root"].name1]
	a2, f2 := answers[operations["root"].name2]
	fmt.Printf("root = %v\n", operations["root"])

	// update the operations with the answers we already have
	for k, v := range operations {
		o := v
		a1, f1 := answers[v.name1]
		a2, f2 := answers[v.name2]

		if f1 {
			o.name1 = strconv.Itoa(a1)
		}
		if f2 {
			o.name2 = strconv.Itoa(a2)
		}
		operations[k] = o
		// fmt.Printf("%s: %s \n", k, operations[k])
	}

	// we found one value for root, set the other value and remap the equations
	fmt.Printf("root = %v\n", operations["root"])
	var next string
	// a1, f1 := answers[operations["root"].name1]
	// a2, f2 := answers[operations["root"].name2]
	if f1 {
		next = operations["root"].name2
		fmt.Printf("%s: %d\n", next, a1)
		answers[next] = a1
		delete(operations, "root")
	} else if f2 {
		next = operations["root"].name1
		fmt.Printf("%s: %d\n", next, a2)
		answers[next] = a2
		delete(operations, "root")
	}

	ReorderOperation(next)

	// take a copy of what we have
	/*
		_answers := make(map[string]int)
		_operations := make(map[string]Operation)
		for k, v := range answers {
			_answers[k] = v
		}
		for k, v := range operations {
			_operations[k] = v
		}

		var a int
		var rootFound bool = false
		for !rootFound {

			// restore what we had and try i
			for k, v := range _answers {
				answers[k] = v
			}
			for k, v := range _operations {
				operations[k] = v
			}
			answers["humn"] = humn

			for len(operations) > 0 {

				var oLeft int = len(operations)
				for n, o := range operations {
					var value int
					var found bool

					value, found = o.Compute()
					if found {
						// fmt.Printf("%s: %d\n", n, value)
						answers[n] = value
						delete(operations, n)
					}
				}

				a, rootFound = answers["root"]
				if a < 0 {
					fmt.Printf("humn = %d\n", humn)
					return
				}

				if oLeft == len(operations) {
					// fmt.Printf("%d operations left - %v\n", len(operations), operations)
					break
				}
			}
			humn = humn + 100000000
		}
	*/

	fmt.Printf("Done: %d\n", answers["humn"])
}

func main() {
	// LoadFile("../sample.txt")
	LoadFile("../puzzle.txt")
}
