package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	MOVE       = 0
	TURN_RIGHT = 1
	TURN_LEFT  = 2
)

const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
	BLANK = 4
	TILE  = 5
	WALL  = 6
)

var tiles [][]Tile

type Tile struct {
	t int
	x int
	y int
	z int
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

func (t Tile) zString() string {
	if t.z == 0 {
		return fmt.Sprintf(" ")

	} else {
		return fmt.Sprintf("%d", t.z)
	}
}

var instructions []Instruction

type Instruction struct {
	action int
	count  int
}

func (i Instruction) String() string {
	switch i.action {
	case MOVE:
		return fmt.Sprintf("%d", i.count)
	case TURN_RIGHT:
		return fmt.Sprintf("R")
	case TURN_LEFT:
		return fmt.Sprintf("L")
	}
	return fmt.Sprintf("?\n")
}

func DirStr(dir int) string {
	switch dir {
	case RIGHT:
		return fmt.Sprintf(">")
	case LEFT:
		return fmt.Sprintf("<")
	case UP:
		return fmt.Sprintf("^")
	case DOWN:
		return fmt.Sprintf("v")
	}
	return "?"
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

func TurnEdge(x, y, x_inc, y_inc *int) {
	if *y == len(tiles) {
		*y = 0
	}
	if *y == -1 {
		*y = len(tiles) - 1
	}

	if *x == len(tiles[*y]) {
		*x = 0
	}

	if *x == -1 {
		*x = len(tiles[*y]) - 1
	}
}

func FindFirstTile(start_x int, start_y int, dir int) (x, y int) {
	x = start_x
	y = start_y
	x_inc, y_inc := CalcIncrements(dir)

	for {
		// fmt.Printf("Find first tile at (%d,%d) is [%s]\n", x, y, tiles[y][x])

		// return the first tile we find in that line
		if tiles[y][x].t != BLANK {
			// fmt.Printf("First tile at (%d, %d)\n", x, y)
			return x, y
		}

		x += x_inc
		y += y_inc

		TurnEdge(&x, &y, &x_inc, &y_inc)
	}
}

func Teleport(start_x int, start_y int, start_dir int) (x int, y int, dir int) {
	x_inc, y_inc := CalcIncrements(start_dir)
	dir = start_dir

	x = start_x + x_inc
	y = start_y + y_inc

	fmt.Printf("TELEPORT %s from (%d, %d)\n", DirStr(start_dir), start_x, start_y)
	fmt.Printf("TELEPORT checking (%d, %d)\n", x, y)
	if start_dir == UP && y < 0 {
		y = len(tiles) - 1
	} else if start_dir == RIGHT && x == len(tiles[y]) {
		x = 0
	} else if start_dir == LEFT && x < 0 {
		x = len(tiles[y]) - 1
	} else if start_dir == DOWN && y == len(tiles) {
		y = 0
	}
	fmt.Printf("Now at (%d, %d)\n", x, y)
	fmt.Printf("[%s]\n", tiles[y][x])

	if tiles[y][x].t == BLANK {
		fmt.Printf("TELEPORT BLANK, find first tile (%d, %d)\n", x, y)
		x, y = FindFirstTile(x, y, dir)
	}

	if tiles[y][x].t == WALL {
		// fmt.Printf("TELEPORT WALL, stay at (%d, %d)\n", start_x, start_y)
		return start_x, start_y, dir
	}

	return x, y, dir
}

func FindNextTile(start_x, start_y, start_dir int) (x, y, dir int) {
	x_inc, y_inc := CalcIncrements(start_dir)
	dir = start_dir
	next_x := start_x + x_inc
	next_y := start_y + y_inc

	// fmt.Printf("Checking (%d, %d)\n", next_x, next_y)

	// map check boundaries
	if next_y < 0 || next_x < 0 || next_y == len(tiles) || next_x == len(tiles[y]) {
		// fmt.Printf("BOUNDARY, teleport...\n")
		next_x, next_y, dir = Teleport(start_x, start_y, start_dir)
	}

	if tiles[next_y][next_x].t == BLANK {
		// fmt.Printf("BLANK, teleport...\n")
		next_x, next_y, dir = Teleport(start_x, start_y, start_dir)
	}

	if tiles[next_y][next_x].t == WALL {
		// fmt.Printf("WALL, stay at (%d, %d)\n", start_x, start_y)
		return start_x, start_y, dir
	}

	if tiles[next_y][next_x].t == TILE {
		// fmt.Printf("Found tile at (%d, %d)\n", next_x, next_y)
		return next_x, next_y, dir
	}

	// fmt.Printf("Found tile we were already at (%d, %d)\n", next_x, next_y)
	return next_x, next_y, dir
}

func Move(start_x, start_y, start_dir, count int) (x, y, dir int) {
	x = start_x
	y = start_y
	dir = start_dir

	tiles[y][x].t = dir
	for i := 0; i < count; i++ {
		x, y, dir = FindNextTile(x, y, dir)
		tiles[y][x].t = dir
	}

	return x, y, dir
}

func Turn(start_dir, turn int) (dir int) {
	switch turn {
	case TURN_LEFT:
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
	case TURN_RIGHT:
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

func WalkMap() {
	var x, y int
	var dir int

	PrintZone(tiles)
	fmt.Println()
	PrintMap(tiles)

	// find starting tile, face right starting from 0,0
	dir = RIGHT
	x, y = FindFirstTile(0, 0, dir)

	for _, i := range instructions {
		// fmt.Printf("I: (%d,%d) %s : %s\n", x, y, DirStr(dir), i)
		switch i.action {
		case MOVE:
			x, y, dir = Move(x, y, dir, i.count)
		case TURN_RIGHT:
			fallthrough
		case TURN_LEFT:
			dir = Turn(dir, i.action)
			tiles[y][x].t = dir
			// fmt.Printf("After turn: %d %s\n", dir, tiles[y][x])
		default:
			fmt.Printf("WHY?!\n")
		}
		// PrintMap(tiles)
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

func PrintZone(tiles [][]Tile) {
	for _, tline := range tiles {
		for _, t := range tline {
			fmt.Print(t.zString())
		}
		fmt.Println()
	}
}

func LoadZones(filePath string) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	y := 0
	for scanner.Scan() {
		var t *Tile
		var tline []Tile
		var line string = scanner.Text()

		x := 0
		for _, b := range line {
			t = &tiles[y][x]

			switch b {
			case ' ':
				t.z = 0
			default:
				t.z = int(b - '0')
			}
			t.x = x + 1
			t.y = y + 1

			x++
		}
		// pad lines with blank tiles
		for ; x < 150; x++ {
			var t Tile
			t.x = x + 1
			t.y = y + 1
			t.t = BLANK
			t.z = 0
			tline = append(tline, t)
		}
		y++
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
		// pad lines with blank tiles
		/*
			for ; x < 150; x++ {
				t.x = x + 1
				t.y = y + 1
				t.t = BLANK
				tline = append(tline, t)
			}
		*/
		y++

		tiles = append(tiles, tline)
	}

	return tiles
}

func PrintRoute(route []Instruction) {
	for _, r := range route {
		fmt.Print(r)
	}
	fmt.Println()
}

func LoadRoute(filePath string) (route []Instruction) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("os.Open %s: %v", filePath, err))
	}
	defer fd.Close()

	route = make([]Instruction, 0)
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		var line string = scanner.Text()
		var d Instruction

		var start int = -1
		for i, b := range line {
			if unicode.IsDigit(b) {
				if start == -1 {
					start = i
				}
			} else {
				if start != -1 {
					var number string = line[start:i]
					d.action = MOVE
					d.count, _ = strconv.Atoi(number)
					start = -1
					route = append(route, d)
				}
				if b == 'R' {
					d.action = TURN_RIGHT
					d.count = 0
				} else {
					d.action = TURN_LEFT
					d.count = 0
				}
				route = append(route, d)
			}
		}
		if start != -1 {
			var number string = line[start:]
			d.action = MOVE
			d.count, _ = strconv.Atoi(number)
			route = append(route, d)
		}
	}

	return route
}

func main() {
	tiles = LoadMap("../sample_map.txt")
	LoadZones("../sample_zones.txt")
	instructions = LoadRoute("../sample_route.txt")
	// tiles = LoadMap("../puzzle_map.txt")
	// instructions = LoadRoute("../puzzle_route.txt")
	WalkMap()
}
