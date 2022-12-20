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

func PrintKnots(knots [10]Location) {
	var min, max Location

	min.x = 0.0
	min.y = 0.0
	max.x = 0.0
	max.y = 0.0

	fmt.Println("------------------------------")
	// fmt.Printf("%d knots\n", len(knots))
	for _, k := range knots {
		// fmt.Println(k)
		if k.x > max.x {
			max.x = k.x
		} else if k.x < min.x {
			min.x = k.x
		}

		if k.y > max.y {
			max.y = k.y
		} else if k.y < min.y {
			min.y = k.y
		}
	}
	// fmt.Println()
	// fmt.Printf("%s %s\n", min, max)

	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			var k Location
			var p, i int
			found := false
			for i, k = range knots {
				if k.x == x && k.y == y {
					if !found {
						p = i
					}
					found = true
				}
			}

			if found {
				if p == 0 {
					fmt.Printf("H")
				} else {
					fmt.Printf("%d", p)
				}
			} else {
				if x == 0 && y == 0 {
					fmt.Printf("s")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("------------------------------")
}

func ProcessMoves(moves []Move) {
	var knots [10]Location

	tailKey := make(map[Location]bool)

	// head starts at origin
	knots[0].x = 0
	knots[0].y = 0

	// initialise all tails
	for i := 1; i < len(knots); i++ {
		knots[i].x = knots[0].x
		knots[i].y = knots[0].y
	}

	// record that we were at the start
	tailKey[knots[len(knots)-1]] = true

	for _, m := range moves {
		var _head *Location
		var _tail *Location

		fmt.Println()
		fmt.Printf("%s %d\n", m.direction, m.count)

		for i := 0; i < m.count; i++ {

			_head = &knots[0]
			// fmt.Printf("%d [%d]:       %s -> ", i+1, 0, *_head)

			// move the head
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
			// fmt.Printf("%s\n", *_head)

			// move the tails
			for t := 1; t < len(knots); t++ {
				_tail = &knots[t]
				// fmt.Printf("%d [%d]: %s %s -> ", i+1, t, *_head, *_tail)

				// special case for diagnal moves with multiple tails
				if math.Abs(_head.x-_tail.x) == 2 && math.Abs(_head.y-_tail.y) == 2 {
					if _head.x > _tail.x {
						_tail.x++
					} else {
						_tail.x--
					}
					if _head.y > _tail.y {
						_tail.y++
					} else {
						_tail.y--
					}
				} else {
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
				}

				// if t == 0 {
				// fmt.Printf("%s\n", *_head)
				// } else {
				// fmt.Printf("%s\n", *_tail)
				// }

				// skip the head as well
				if t == len(knots)-1 {
					if _, ok := tailKey[*_tail]; !ok {
						tailKey[*_tail] = true
						// fmt.Printf("# %s\n", *_tail)
					}
				}

				// move new head ?!
				_head = &knots[t]
			}
			// for i := 0; i < len(knots); i++ {
			// fmt.Printf("knots[%d] %s\n", i, knots[i])
			// }
			PrintKnots(knots)
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
	moves := LoadMoves("../sample2.txt")
	ProcessMoves(moves)
}
