package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Move struct {
	direction string
	count     int
}

type Location struct {
	x float64
	y float64
}

func (l Location) String() string {
	return fmt.Sprintf("(%v,%v)", l.x, l.y)
}

func ProcessMoved(moves []Move) {
	var head, tail Location

	tails := make(map[Location]bool)

	// head starts at origin
	head.x = 0
	head.y = 0

	// tail starts with head
	tail.x = head.x
	tail.y = head.y

	// record that we were at the start
	tails[tail] = true

	for _, m := range moves {
		// fmt.Printf("%s %d %s %s = ", m.direction, m.count, head, tail)

		for i := 0; i < m.count; i++ {
			switch m.direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			if math.Abs(head.x-tail.x) == 2 {
				if head.x > tail.x {
					tail.x++
				} else {
					tail.x--
				}
				if tail.y != head.y {
					tail.y = head.y
				}
			}

			if math.Abs(head.y-tail.y) == 2 {
				if head.y > tail.y {
					tail.y++
				} else {
					tail.y--
				}
				if tail.x != head.x {
					tail.x = head.x
				}
			}

			if _, ok := tails[tail]; !ok {
				tails[tail] = true
			}
		}
		// fmt.Printf("%s %s\n\n", head, tail)
	}

	fmt.Println(len(tails))
}

func LoadMoves(filePath string) (moves []Move) {
	var count int
	var direction string

	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	for {
		_, err := fmt.Fscanf(fd, "%s %d\n", &direction, &count)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Sprintf("fmt.Fscanf %s: %v", filePath, err))
		}

		var move Move
		move.direction = direction
		move.count = count
		moves = append(moves, move)
	}

	return moves
}

func main() {
	moves := LoadMoves("../puzzle.txt")
	ProcessMoved(moves)
}
