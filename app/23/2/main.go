package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 4 move directions in order of attempts
const (
	NORTH = 0
	SOUTH = 1
	WEST  = 2
	EAST  = 3
)

func DirectionString(d int) string {
	switch d {
	case NORTH:
		return "NORTH"
	case SOUTH:
		return "SOUTH"
	case WEST:
		return "WEST"
	case EAST:
		return "EAST"
	}
	return ""
}

func NorthWest(t *Tile) Tile {
	return Tile{t.x - 1, t.y - 1}
}
func NorthEast(t *Tile) Tile {
	return Tile{t.x + 1, t.y - 1}
}
func SouthWest(t *Tile) Tile {
	return Tile{t.x - 1, t.y + 1}
}
func SouthEast(t *Tile) Tile {
	return Tile{t.x + 1, t.y + 1}
}
func North(t *Tile) Tile {
	return Tile{t.x, t.y - 1}
}
func East(t *Tile) Tile {
	return Tile{t.x + 1, t.y}
}
func South(t *Tile) Tile {
	return Tile{t.x, t.y + 1}
}
func West(t *Tile) Tile {
	return Tile{t.x - 1, t.y}
}

func NextDirection(d int) int {
	d++
	if d == 4 {
		d = 0
	}
	return d
}

type Moves [4]bool

type Tile struct {
	x int
	y int
}
type Tiles []Tile
type TilePtr *Tile

func (t Tile) String() string {
	return fmt.Sprintf("(%d,%d)", t.x, t.y)
}

func (t Tiles) Len() int {
	return len(t)
}

func (t Tiles) Less(i, j int) bool {
	if t[i].y < t[j].y {
		return true
	} else if t[i].y > t[j].y {
		return false
	} else {
		if t[i].x < t[j].x {
			return true
		} else if t[i].x >= t[i].x {
			return false
		}
	}
	return false
}

func (t Tiles) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type ProposedMove struct {
	to   Tile
	from []TilePtr
}
type ProposedMoves []ProposedMove

func (p ProposedMove) String() string {
	return fmt.Sprintf("(%d,%d) [%d]", p.to.x, p.to.y, len(p.from))
}

func Less(i, j int) bool {
	if tiles[i].y < tiles[j].y {
		return true
	} else if tiles[i].y > tiles[j].y {
		return false
	} else {
		if tiles[i].x < tiles[j].x {
			return true
		} else if tiles[i].x >= tiles[i].x {
			return false
		}
	}
	return false
}

// hold the min & max values of the map (W, N) is top left and (E, S) is bottom right
var N, S, E, W int

// list of current tile positions
var tiles Tiles

func PrintMap() {
	var idx int = 0

	sort.SliceStable(tiles, Less)
	// fmt.Printf("MAP:\n")
	for _, t := range tiles {
		// fmt.Printf("%s ", t)

		if N > t.y {
			N = t.y
		}
		if S < t.y {
			S = t.y
		}
		if W > t.x {
			W = t.x
		}
		if E < t.x {
			E = t.x
		}
	}
	// fmt.Printf("Corners: (%d,%d), (%d,%d)\n", W, N, E, S)
	// fmt.Println()
	for y := N; y <= S; y++ {
		for x := W; x <= E; x++ {
			if (idx < len(tiles)) && (tiles[idx].x == x) && (tiles[idx].y == y) {
				fmt.Print("#")
				idx++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func FindAdjacentElves(t *Tile) Moves {
	// var moves Moves
	var moves = Moves{true, true, true, true}

	nw := NorthWest(t)
	n := North(t)
	ne := NorthEast(t)
	e := East(t)
	w := West(t)
	sw := SouthWest(t)
	s := South(t)
	se := SouthEast(t)

	// fmt.Println("---")
	// fmt.Printf("%v %v %v\n", nw, n, ne)
	// fmt.Printf("%v %v %v\n", w, t, e)
	// fmt.Printf("%v %v %v\n", sw, s, se)
	// fmt.Println("---")

	for _, x := range tiles {
		// skip self
		if x == *t {
			continue
		}
		if moves[NORTH] && (x == nw || x == n || x == ne) {
			// fmt.Printf("Can't go NORTH\n")
			moves[NORTH] = false
		}
		if moves[SOUTH] && (x == se || x == s || x == sw) {
			// fmt.Printf("Can't go SOUTH\n")
			moves[SOUTH] = false
		}
		if moves[WEST] && (x == sw || x == w || x == nw) {
			// fmt.Printf("Can't go WEST\n")
			moves[WEST] = false
		}
		if moves[EAST] && (x == ne || x == e || x == se) {
			// fmt.Printf("Can't go EAST\n")
			moves[EAST] = false
		}
	}

	return moves
}

func WalkMap() {
	var startDir int = NORTH
	var proposedMoves ProposedMoves
	var noOneMoved bool

	PrintMap()
	for r := 0; ; r++ {

		noOneMoved = true
		proposedMoves = make(ProposedMoves, 0)

		// PrintMap()
		// compute proposed moves
		for tIdx := range tiles {

			var t *Tile = &tiles[tIdx]

			// fmt.Printf("%s ", t)
			moves := FindAdjacentElves(t)
			if moves[NORTH] && moves[EAST] && moves[SOUTH] && moves[WEST] {
				// fmt.Printf("STAY\n")
				continue
			}
			noOneMoved = false

			// try each of the 4 directions for this move
			dir := startDir
			for i := 0; i < 4; i++ {
				if moves[dir] {

					var tileTo Tile
					switch dir {
					case NORTH:
						tileTo = North(t)
					case EAST:
						tileTo = East(t)
					case SOUTH:
						tileTo = South(t)
					case WEST:
						tileTo = West(t)
					}
					// fmt.Printf("move %s to %s \n", DirectionString(dir), tileTo)

					var foundIdx int = -1
					for x := range proposedMoves {
						if proposedMoves[x].to == tileTo {
							foundIdx = x
							break
						}
					}

					if foundIdx == -1 {
						var p ProposedMove
						p.to = tileTo
						p.from = make([]TilePtr, 0)
						p.from = append(p.from, t)
						proposedMoves = append(proposedMoves, p)
					} else {
						proposedMoves[foundIdx].from = append(proposedMoves[foundIdx].from, t)
					}

					// we found a direction to move to, don't try others
					break
				}

				// use next direction
				dir = NextDirection(dir)
			}
		}

		if noOneMoved {
			fmt.Printf("No one moved in round %d\n", r+1)
			break
		}

		// fmt.Printf("%d proposed moves:\n", len(proposedMoves))
		for j := range proposedMoves {
			// fmt.Printf("%v ", proposedMoves[j])
			if len(proposedMoves[j].from) == 1 {
				proposedMoves[j].from[0].x = proposedMoves[j].to.x
				proposedMoves[j].from[0].y = proposedMoves[j].to.y
			} else {
				// fmt.Printf("skipping ")
			}
			// fmt.Printf("\n")
		}
		// PrintMap()

		// start next round in next direction
		startDir = NextDirection(startDir)
	}
	// fmt.Println("Done:")
	PrintMap()
	fmt.Printf("Empty tiles = %d * %d - %d = %d\n", S-N+1, E-W+1, len(tiles), (S-N+1)*(E-W+1)-len(tiles))
}

func LoadMap(filePath string) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	N = 99
	W = 99
	S = 0
	E = 0

	tiles = make(Tiles, 0)

	y := 0
	for scanner.Scan() {
		var t Tile
		var line string = scanner.Text()

		x := 0
		for _, b := range line {
			if b == '#' {
				t.x = x + 1
				t.y = y + 1
				tiles = append(tiles, t)

				if N > t.y {
					N = t.y
				}
				if S < t.y {
					S = t.y
				}
				if W > t.x {
					W = t.x
				}
				if E < t.x {
					E = t.x
				}
			}

			x++
		}
		y++
	}
}

func main() {
	// LoadMap("../sample_map.txt")
	LoadMap("../puzzle_map.txt")
	WalkMap()
}
