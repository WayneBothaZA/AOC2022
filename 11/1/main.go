package main

import (
	"fmt"
	"math"
)

const (
	Multiply int = 0
	Square       = 1
	Add          = 2
	Subtract     = 3
)

type MonkeyOperation interface {
	DoOperation(int) int
}

type MonkeyTest interface {
	RunTest() bool
}

type Monkey struct {
	id             int
	items          []int
	operation      string
	operationValue int
	test           int
	next           [2]int
	inspections    int
}

func (m Monkey) CalcNewValue(old int) int {
	var new int
	switch m.operation {
	case "*":
		new = old * m.operationValue
	case "+":
		new = old + m.operationValue
	case "square":
		new = old * old
	}
	return new
}

func (m *Monkey) Inspect(index int) {
	var old int = m.items[index]
	var new int
	new = m.CalcNewValue(old)
	fmt.Printf("    Worry level is %s by %d to %d.\n", m.operation, m.operationValue, new)
	new = int(math.Floor(float64(new) / 3.0))
	fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", new)
	m.items[index] = new
	m.inspections++
}

func (m *Monkey) Throw(index int) {
	var testPass bool = ((m.items[index] % m.test) == 0)
	var testPassString string
	var nextMonkey int

	if testPass {
		testPassString = ""
		nextMonkey = m.next[1]
	} else {
		testPassString = " not"
		nextMonkey = m.next[0]
	}
	fmt.Printf("    Current worry level is%s divisible by %d.\n", testPassString, m.test)
	fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", m.items[index], nextMonkey)

	monkeys[nextMonkey].Catch(m.items[index])
	// fmt.Printf("(%d)\n", len(monkeys[nextMonkey].items))

	// fmt.Printf("x=> Monkey %d catches %d (%d --> ", monkeys[nextMonkey].id, m.items[index], len(monkeys[nextMonkey].items))
	// m.items = append(m.items, m.items[index])
	// fmt.Printf("%d)\n", len(monkeys[nextMonkey].items))
}

func (m *Monkey) Catch(item int) {
	// fmt.Printf("y=> Monkey %d catches %d (%d --> ", m.id, item, len(m.items))
	m.items = append(m.items, item)
	// fmt.Printf("%d)\n", len(m.items))
}

func (m Monkey) String() string {
	var _operation string = m.operation
	var _operationValue string

	_operationValue = fmt.Sprintf("%d", m.operationValue)

	if _operation == "square" {
		_operation = "*"
		_operationValue = "old"
	}

	return fmt.Sprintf(
		`Monkey %d:  
  Starting items %v  
  Operation: new = old %s %s
  Test: divisible by %d
    if true: throw to monkey %d
    if false: throw to monkey %d
  `, m.id, m.items, _operation, _operationValue, m.test, m.next[1], m.next[0])
}

var monkeys []Monkey

func PrintItems() {
	for _, monkey := range monkeys {
		fmt.Printf("%d: %v\n", monkey.id, monkey.items)
	}
}

func ProcessMonkeys() {
	for _, monkey := range monkeys {
		fmt.Printf("%s\n", monkey)
	}

	PrintItems()
	for x := 0; x < 20; x++ {
		for m, _ := range monkeys {
			fmt.Printf("Monkey %d (%d):\n", monkeys[m].id, len(monkeys[m].items))
			for i, item := range monkeys[m].items {
				fmt.Printf("  Monkey inspects an item with worry level of %d.\n", item)
				monkeys[m].Inspect(i)
				monkeys[m].Throw(i)
			}
			monkeys[m].items = nil
		}
	}
	PrintItems()
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", monkey.id, monkey.inspections)
	}
}

func LoadSample() {
	var m Monkey

	m.id = 0
	m.items = []int{79, 98}
	m.operation = "*"
	m.operationValue = 19
	m.test = 23
	m.next = [2]int{3, 2}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 1
	m.items = []int{54, 65, 75, 74}
	m.operation = "+"
	m.operationValue = 6
	m.test = 19
	m.next = [2]int{0, 2}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 2
	m.items = []int{79, 60, 97}
	m.operation = "square"
	m.operationValue = 0
	m.test = 13
	m.next = [2]int{3, 1}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 3
	m.items = []int{74}
	m.operation = "+"
	m.operationValue = 3
	m.test = 17
	m.next = [2]int{1, 0}
	m.inspections = 0
	monkeys = append(monkeys, m)
}

func LoadPuzzle() {
	var m Monkey

	m.id = 0
	m.items = []int{74, 64, 74, 63, 53}
	m.operation = "*"
	m.operationValue = 7
	m.test = 5
	m.next = [2]int{6, 1}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 1
	m.items = []int{69, 99, 95, 62}
	m.operation = "square"
	m.operationValue = 0
	m.test = 17
	m.next = [2]int{5, 2}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 2
	m.items = []int{59, 81}
	m.operation = "+"
	m.operationValue = 8
	m.test = 7
	m.next = [2]int{3, 4}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 3
	m.items = []int{50, 67, 63, 57, 63, 83, 97}
	m.operation = "+"
	m.operationValue = 4
	m.test = 17
	m.next = [2]int{7, 0}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 4
	m.items = []int{61, 94, 85, 52, 81, 90, 94, 70}
	m.operation = "+"
	m.operationValue = 3
	m.test = 19
	m.next = [2]int{3, 7}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 5
	m.items = []int{69}
	m.operation = "+"
	m.operationValue = 5
	m.test = 3
	m.next = [2]int{2, 4}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 6
	m.items = []int{54, 55, 58}
	m.operation = "+"
	m.operationValue = 7
	m.test = 11
	m.next = [2]int{5, 1}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 7
	m.items = []int{79, 51, 83, 88, 93, 76}
	m.operation = "*"
	m.operationValue = 3
	m.test = 2
	m.next = [2]int{6, 0}
	m.inspections = 0
	monkeys = append(monkeys, m)
}

func main() {
	// LoadSample()
	LoadPuzzle()
	ProcessMonkeys()
}
