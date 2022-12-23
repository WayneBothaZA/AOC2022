package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	MOVE  = -1
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
	BLANK = 4
	TILE  = 5
	WALL  = 6
)

type Tile struct {
	t int
	x int
	y int
}

func (t Tile) String() string {
	switch t.t {
	case RIGHT:
		return fmt.Sprintf(">")
	case LEFT:
		return fmt.Sprintf("<")
	case UP:
		return fmt.Sprintf("^")
	case DOWN:
		return fmt.Sprintf("v")
	case BLANK:
		return fmt.Sprintf(" ")
	case TILE:
		return fmt.Sprintf(".")
	case WALL:
		return fmt.Sprintf("#")
	}
	return fmt.Sprintf("?")
}

type Direction struct {
	direction int
	count     int
}

func (d Direction) String() string {
	switch d.direction {
	case MOVE:
		return fmt.Sprintf("%d", d.count)
	case RIGHT:
		return fmt.Sprintf("R")
	case LEFT:
		return fmt.Sprintf("L")
	}
	return fmt.Sprintf("?\n")
}

func CalcIncrements(dir int) (x_inc, y_inc int) {
	switch dir {
	case RIGHT:
		x_inc = 1
		y_inc = 0
	case LEFT:
		x_inc = -1
		y_inc = 0
	case UP:
		x_inc = 0
		y_inc = -1
	case DOWN:
		x_inc = 0
		y_inc = 1
	}
	return x_inc, y_inc
}

func FindFirstTile(tiles [][]Tile, start_x int, start_y int, dir int) (x, y int) {
	x = start_x
	y = start_y
	x_inc, y_inc := CalcIncrements(dir)

	for {
		fmt.Printf("Find first tile at (%d,%d) is [%s]\n", x, y, tiles[y][x])

		// return the first tile we find in that line
		if tiles[y][x].t != BLANK {
			fmt.Printf("First tile at (%d, %d)\n", x, y)
			return x, y
		}

		x += x_inc
		y += y_inc

		if y == len(tiles) {
			y = 0
		}

		if x == len(tiles[y]) {
			x = 0
		}
	}
}

func Teleport(tiles [][]Tile, start_x int, start_y int, dir int) (x int, y int, found bool) {
	x = start_x
	y = start_y

	if dir == RIGHT && x == len(tiles[y]) {
		x = 0
		fmt.Printf("TELEPORT RIGHT to (%d, %d)\n", x, y)
	} else if dir == LEFT && x < 0 {
		x = len(tiles[y]) - 1
		fmt.Printf("TELEPORT LEFT to (%d, %d)\n", x, y)
	} else if dir == DOWN && y == len(tiles) {
		y = 0
		fmt.Printf("TELEPORT DOWN to (%d, %d)\n", x, y)
	} else if dir == UP && y < 0 {
		y = len(tiles) - 1
		fmt.Printf("TELEPORT UP to (%d, %d)\n", x, y)
	}

	if tiles[y][x].t == BLANK {
		fmt.Printf("TELEPORT BLANK, find first tile (%d, %d)\n", x, y)
		x, y = FindFirstTile(tiles, x, y, dir)
	}

	if tiles[y][x].t == WALL {
		fmt.Printf("TELEPORT WALL, stay at (%d, %d)\n", start_x, start_y)
		return start_x, start_y, false
	}

	return x, y, true
}

func FindNextTile(tiles [][]Tile, start_x int, start_y int, dir int) (x, y int) {
	// get the next tile
	x_inc, y_inc := CalcIncrements(dir)

	fmt.Printf("Checking (%d, %d)\n", start_x+x_inc, start_y+y_inc)

	// teleport over edges and blanks
	x, y, found := Teleport(tiles, start_x+x_inc, start_y+y_inc, dir)
	if !found {
		fmt.Printf("Move back from (%d, %d)\n", x, y)
		x += (x_inc * -1)
		y += (y_inc * -1)
	}

	for {
		fmt.Printf("Now at (%d, %d)\n", x, y)
		// return the tile we found (we might already have visited here)
		if tiles[y][x].t != BLANK && tiles[y][x].t != WALL {
			fmt.Printf("Found (%d, %d)\n", x, y)
			return x, y
		}

		x, y, found = Teleport(tiles, start_x+x_inc, start_y+y_inc, dir)
		if !found {
			fmt.Printf("Bail out!")
			return x, y
		}
	}
}

func Move(tiles [][]Tile, start_x, start_y, dir, count int) (x, y int) {
	x = start_x
	y = start_y

	tiles[y][x].t = dir
	for i := 0; i < count; i++ {
		x, y = FindNextTile(tiles, x, y, dir)
		tiles[y][x].t = dir
	}

	return x, y
}

func Turn(start_dir, turn int) (dir int) {
	switch turn {
	case LEFT:
		switch start_dir {
		case RIGHT:
			dir = UP
		case DOWN:
			dir = RIGHT
		case LEFT:
			dir = DOWN
		case UP:
			dir = LEFT
		}
	case RIGHT:
		switch start_dir {
		case RIGHT:
			dir = DOWN
		case DOWN:
			dir = LEFT
		case LEFT:
			dir = UP
		case UP:
			dir = RIGHT
		}
	}
	return dir
}

func WalkMap(tiles [][]Tile, route []Direction) {
	var x, y int
	var t *Tile
	var dir int

	PrintMap(tiles)

	// find starting tile, face right starting from 0,0
	dir = RIGHT
	x, y = FindFirstTile(tiles, 0, 0, dir)

	for _, d := range route {
		t = &tiles[y][x]
		t.t = dir
		fmt.Printf("(%d,%d) %s: %s\n", x, y, t, d)
		switch d.direction {
		case MOVE:
			x, y = Move(tiles, x, y, dir, d.count)
		case RIGHT:
			fallthrough
		case LEFT:
			dir = Turn(dir, d.direction)
			tiles[y][x].t = dir
		}
		PrintMap(tiles)
	}

	fmt.Println()
	PrintMap(tiles)

	fmt.Printf("%d\n", (y+1)*1000+(x+1)*4+dir)
}

func PrintMap(tiles [][]Tile) {
	for _, tline := range tiles {
		for _, t := range tline {
			fmt.Print(t)
		}
		fmt.Println()
	}
}

func LoadMap(filePath string) (tiles [][]Tile) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	tiles = make([][]Tile, 0)
	y := 0
	for scanner.Scan() {
		var t Tile
		var tline []Tile
		var line string = scanner.Text()

		tline = make([]Tile, 0)
		x := 0
		for _, b := range line {
			switch b {
			case ' ':
				t.t = BLANK
			case '.':
				t.t = TILE
			case '#':
				t.t = WALL
			}
			t.x = x + 1
			t.y = y + 1

			tline = append(tline, t)
			x++
		}
		y++

		tiles = append(tiles, tline)
	}

	return tiles
}

func PrintRoute(route []Direction) {
	for _, r := range route {
		fmt.Print(r)
	}
	fmt.Println()
}

func LoadRoute(filePath string) (route []Direction) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	route = make([]Direction, 0)
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		var line string = scanner.Text()
		var d Direction

		var start int = -1
		for i, b := range line {
			if unicode.IsDigit(b) {
				if start == -1 {
					start = i
				}
			} else {
				if start != -1 {
					var number string = line[start:i]
					d.direction = MOVE
					d.count, _ = strconv.Atoi(number)
					start = -1
					route = append(route, d)
				}
				if b == 'R' {
					d.direction = RIGHT
					d.count = 0
				} else {
					d.direction = LEFT
					d.count = 0
				}
				route = append(route, d)
			}
		}
		if start != -1 {
			var number string = line[start:]
			d.direction = MOVE
			d.count, _ = strconv.Atoi(number)
			route = append(route, d)
		}
	}

	return route
}

func main() {
	// tiles := LoadMap("../sample_map.txt")
	// route := LoadRoute("../sample_route.txt")
	tiles := LoadMap("../puzzle_map.txt")
	route := LoadRoute("../puzzle_route.txt")
	WalkMap(tiles, route)
}
