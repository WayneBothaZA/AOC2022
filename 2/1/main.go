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
		}
	case Paper:
		switch m_item {
		case Rock:
			return Loose
		case Paper:
			return Draw
		case Scissors:
			return Win
		}
	case Scissors:
		switch m_item {
		case Rock:
			return Win
		case Paper:
			return Loose
		case Scissors:
			return Draw
		}
	}
	panic(fmt.Sprintf("CalcWinner %v, %v", o_item, m_item))
}

func CalcScore(o_item, m_item Item) (Result, int, int) {

	result := CalcResult(o_item, m_item)

	return result, m_item.Score, result.Score
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

		result, itemScore, resultScore := CalcScore(o_item, m_item)
		roundScore := itemScore + resultScore

		fmt.Printf("%v,%v,%v,%d+%d=%d,%d\n", o_item.Name, m_item.Name, result.Name, itemScore, resultScore, roundScore, Total)
		Total += roundScore
	}
	fmt.Printf("%d\n", Total)
}

func main() {
	readFile("../sample.txt")
}
