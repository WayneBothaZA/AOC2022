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

func ProcessMoves(moves []Move) {
	var head Location
	var tail Location
	var tails [9]Location

	tailKey := make(map[Location]bool)

	// head starts at origin
	head.x = 0
	head.y = 0

	// tail starts with head
	tail.x = head.x
	tail.y = head.y

	// initialise all 10 tails
	for i := 0; i < len(tails); i++ {
		tails[i] = tail
	}

	// record that we were at the start
	tailKey[tail] = true

	for x, m := range moves {
		var _head *Location
		var _tail *Location

		fmt.Println()
		fmt.Printf("%d: %s %d\n", x+1, m.direction, m.count)

		for i := 0; i < m.count; i++ {
			for t := 0; t < len(tails); t++ {
				if t == 0 {
					_head = &head
				} else {
					_head = &tails[t-1]
				}
				_tail = &tails[t]

				// only move head on first round...
				if t == 0 {
					switch m.direction {
					case "R":
						_head.x++
					case "L":
						_head.x--
					case "U":
						_head.y++
					case "D":
						_head.y--
					}
				}

				if math.Abs(_head.x-_tail.x) == 2 {
					if _head.x > _tail.x {
						_tail.x++
					} else {
						_tail.x--
					}
					if _tail.y != _head.y {
						_tail.y = _head.y
					}
				}

				if math.Abs(_head.y-_tail.y) == 2 {
					if _head.y > _tail.y {
						_tail.y++
					} else {
						_tail.y--
					}
					if _tail.x != _head.x {
						_tail.x = _head.x
					}
				}

				if t == len(tails)-1 {
					if _, ok := tailKey[*_tail]; !ok {
						tailKey[*_tail] = true
						fmt.Printf("# %s\n", *_tail)
					}
				}
			}
			fmt.Printf("%d:\n", i+1)
			fmt.Printf("head:    %s\n", head)
			for i := 0; i < len(tails); i++ {
				fmt.Printf("tails[%d] %s\n", i+1, tails[i])
			}
		}
	}

	fmt.Println(len(tailKey))
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
	moves := LoadMoves("../sample.txt")
	ProcessMoves(moves)
}
