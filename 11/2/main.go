package main

import (
	"fmt"
	"math/big"
)

type Monkey struct {
	id             int
	items          []*big.Int
	operation      string
	operationValue *big.Int
	test           *big.Int
	next           [2]int
	inspections    int
}

func (m Monkey) CalcNewValue(old *big.Int) *big.Int {
	switch m.operation {
	case "*":
		return old.Mul(old, m.operationValue)
	case "+":
		return old.Add(old, m.operationValue)
	case "square":
		return old.Mul(old, old)
	}
	return big.NewInt(0)
}

func (m *Monkey) Inspect(index int) {
	m.items[index] = m.CalcNewValue(m.items[index])
	m.inspections++
}

func (m *Monkey) Throw(index int) {
	var nextMonkey int

	x := big.NewInt(0)
	x = x.Mod(m.items[index], m.test)

	if x.Cmp(big.NewInt(0)) == 0 {
		nextMonkey = m.next[1]
	} else {
		nextMonkey = m.next[0]
	}

	monkeys[nextMonkey].Catch(m.items[index])
}

func (m *Monkey) Catch(item *big.Int) {
	m.items = append(m.items, item)
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

	for x := 0; x < 10000; x++ {
		for m, _ := range monkeys {
			for i, _ := range monkeys[m].items {
				monkeys[m].Inspect(i)
				monkeys[m].Throw(i)
			}
			monkeys[m].items = nil
		}
		if (x % 100) == 0 {
			fmt.Printf("%d\n", x)
		}
		/*
			fmt.Printf("-- After round %d ==\n", x+1)
			for _, monkey := range monkeys {
				fmt.Printf("Monkey %d inspected items %d times.\n", monkey.id, monkey.inspections)
			}
		*/
	}

	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", monkey.id, monkey.inspections)
	}
}

func LoadSample() {
	var m Monkey

	m.id = 0
	m.items = nil
	m.items = append(m.items, big.NewInt(79))
	m.items = append(m.items, big.NewInt(98))
	m.operation = "*"
	m.operationValue = big.NewInt(19)
	m.test = big.NewInt(23)
	m.next = [2]int{3, 2}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 1
	m.items = nil
	m.items = append(m.items, big.NewInt(54))
	m.items = append(m.items, big.NewInt(65))
	m.items = append(m.items, big.NewInt(75))
	m.items = append(m.items, big.NewInt(74))
	m.operation = "+"
	m.operationValue = big.NewInt(6)
	m.test = big.NewInt(19)
	m.next = [2]int{0, 2}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 2
	m.items = nil
	m.items = append(m.items, big.NewInt(79))
	m.items = append(m.items, big.NewInt(60))
	m.items = append(m.items, big.NewInt(97))
	m.operation = "square"
	m.operationValue = big.NewInt(0)
	m.test = big.NewInt(13)
	m.next = [2]int{3, 1}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 3
	m.items = nil
	m.items = append(m.items, big.NewInt(74))
	m.operation = "+"
	m.operationValue = big.NewInt(3)
	m.test = big.NewInt(17)
	m.next = [2]int{1, 0}
	m.inspections = 0
	monkeys = append(monkeys, m)
}

func LoadPuzzle() {
	var m Monkey

	m.id = 0
	m.items = nil
	m.items = append(m.items, big.NewInt(74))
	m.items = append(m.items, big.NewInt(64))
	m.items = append(m.items, big.NewInt(74))
	m.items = append(m.items, big.NewInt(63))
	m.items = append(m.items, big.NewInt(53))
	m.operation = "*"
	m.operationValue = big.NewInt(7)
	m.test = big.NewInt(5)
	m.next = [2]int{6, 1}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 1
	m.items = nil
	m.items = append(m.items, big.NewInt(69))
	m.items = append(m.items, big.NewInt(99))
	m.items = append(m.items, big.NewInt(95))
	m.items = append(m.items, big.NewInt(62))
	m.operation = "square"
	m.operationValue = big.NewInt(0)
	m.test = big.NewInt(17)
	m.next = [2]int{5, 2}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 2
	m.items = nil
	m.items = append(m.items, big.NewInt(59))
	m.items = append(m.items, big.NewInt(81))
	m.operation = "+"
	m.operationValue = big.NewInt(8)
	m.test = big.NewInt(7)
	m.next = [2]int{3, 4}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 3
	m.items = nil
	m.items = append(m.items, big.NewInt(50))
	m.items = append(m.items, big.NewInt(67))
	m.items = append(m.items, big.NewInt(63))
	m.items = append(m.items, big.NewInt(57))
	m.items = append(m.items, big.NewInt(63))
	m.items = append(m.items, big.NewInt(83))
	m.items = append(m.items, big.NewInt(97))
	m.operation = "+"
	m.operationValue = big.NewInt(4)
	m.test = big.NewInt(17)
	m.next = [2]int{7, 0}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 4
	m.items = nil
	m.items = append(m.items, big.NewInt(61))
	m.items = append(m.items, big.NewInt(94))
	m.items = append(m.items, big.NewInt(85))
	m.items = append(m.items, big.NewInt(52))
	m.items = append(m.items, big.NewInt(81))
	m.items = append(m.items, big.NewInt(90))
	m.items = append(m.items, big.NewInt(94))
	m.items = append(m.items, big.NewInt(70))
	m.operation = "+"
	m.operationValue = big.NewInt(3)
	m.test = big.NewInt(19)
	m.next = [2]int{3, 7}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 5
	m.items = nil
	m.items = append(m.items, big.NewInt(69))
	m.operation = "+"
	m.operationValue = big.NewInt(5)
	m.test = big.NewInt(3)
	m.next = [2]int{2, 4}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 6
	m.items = nil
	m.items = append(m.items, big.NewInt(54))
	m.items = append(m.items, big.NewInt(55))
	m.items = append(m.items, big.NewInt(58))
	m.operation = "+"
	m.operationValue = big.NewInt(7)
	m.test = big.NewInt(11)
	m.next = [2]int{5, 1}
	m.inspections = 0
	monkeys = append(monkeys, m)

	m.id = 7
	m.items = nil
	m.items = append(m.items, big.NewInt(79))
	m.items = append(m.items, big.NewInt(51))
	m.items = append(m.items, big.NewInt(83))
	m.items = append(m.items, big.NewInt(88))
	m.items = append(m.items, big.NewInt(93))
	m.items = append(m.items, big.NewInt(76))
	m.operation = "*"
	m.operationValue = big.NewInt(3)
	m.test = big.NewInt(2)
	m.next = [2]int{6, 0}
	m.inspections = 0
	monkeys = append(monkeys, m)
}

func main() {
	//LoadSample()
	LoadPuzzle()
	ProcessMonkeys()
}
