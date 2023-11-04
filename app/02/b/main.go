package main

import (
	"fmt"
	"io"
	"os"
)

type Item struct {
	Name  string
	Score int
}

var (
	Rock     Item = Item{"Rock", 1}
	Paper    Item = Item{"Paper", 2}
	Scissors Item = Item{"Scissors", 3}
)

type Result struct {
	Name  string
	Score int
}

var (
	Loose Result = Result{"Loose", 0}
	Draw  Result = Result{"Draw", 3}
	Win   Result = Result{"Win", 6}
)

func CalcResult(o_item, m_item Item) Result {
	switch o_item {
	case Rock:
		switch m_item {
		case Rock:
			return Draw
		case Paper:
			return Win
		case Scissors:
			return Loose
		default:
			panic(fmt.Sprintf("CalcWinner %v,%v", o_item, m_item))
		}
	case Paper:
		switch m_item {
		case Rock:
			return Loose
		case Paper:
			return Draw
		case Scissors:
			return Win
		default:
			panic(fmt.Sprintf("CalcWinner %v,%v", o_item, m_item))
		}
	case Scissors:
		switch m_item {
		case Rock:
			return Win
		case Paper:
			return Loose
		case Scissors:
			return Draw
		default:
			panic(fmt.Sprintf("CalcWinner %v,%v", o_item, m_item))
		}
	}
	panic(fmt.Sprintf("CalcWinner %v, %v", o_item, m_item))
}

func CalcMyItem(o_item Item, m_result Result) Item {
	switch o_item {
	case Rock:
		switch m_result {
		case Loose:
			return Scissors
		case Draw:
			return Rock
		case Win:
			return Paper
		}
	case Paper:
		switch m_result {
		case Loose:
			return Rock
		case Draw:
			return Paper
		case Win:
			return Scissors
		}
	case Scissors:
		switch m_result {
		case Loose:
			return Paper
		case Draw:
			return Scissors
		case Win:
			return Rock
		}
	}
	panic(fmt.Sprintf("CalcWinner %v,%v", o_item, m_result))
}

func CalcScore(o_item Item, m_result Result) (int, int) {
	m_item := CalcMyItem(o_item, m_result)
	result := CalcResult(o_item, m_item)
	return m_item.Score, result.Score
}

func ConvertItem(item string) Item {
	switch item[0] {
	case 'A':
		return Rock
	case 'B':
		return Paper
	case 'C':
		return Scissors
	case 'X':
		return Rock
	case 'Y':
		return Paper
	case 'Z':
		return Scissors
	}
	panic(fmt.Sprintf("ConvertItem %v", item))
}

func readFile(filePath string) {
	var Total int = 0

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	for {
		var o, m string
		var o_item, m_item Item

		_, err := fmt.Fscanf(fd, "%s %s\n", &o, &m)
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		o_item = ConvertItem(o)
		m_item = ConvertItem(m)

		itemScore, resultScore := CalcScore(o_item, m_item)
		roundScore := itemScore + resultScore

		fmt.Printf("%v,%v,%v,%d+%d=%d,%d\n", o_item.Name, m_item.Name, result.Name, itemScore, resultScore, roundScore, Total)
		Total += roundScore
	}
	fmt.Printf("%d\n", Total)
}

func main() {
	readFile("../strategy.txt")
}
