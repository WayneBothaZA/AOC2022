package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	blank = 0
	tile  = 1
	wall  = 2
)

const (
	move  = -1
	right = 0
	down  = 1
	left  = 2
	up    = 3
)

type Tile struct {
	t int
	x int
	y int
}

func (t Tile) String() string {
	switch t.t {
	case blank:
		return fmt.Sprintf(" ")
	case tile:
		return fmt.Sprintf(".")
	case wall:
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
	case move:
		return fmt.Sprintf("%d", d.count)
	case right:
		return fmt.Sprintf("R")
	case left:
		return fmt.Sprintf("L")
	}
	return fmt.Sprintf("?\n")
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
				t.t = blank
			case '.':
				t.t = tile
			case '#':
				t.t = wall
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
					d.direction = move
					d.count, _ = strconv.Atoi(number)
					start = -1
					route = append(route, d)
				}
				if b == 'R' {
					d.direction = right
					d.count = 0
				} else {
					d.direction = left
					d.count = 0
				}
				route = append(route, d)
			}
		}
		if start != -1 {
			var number string = line[start:]
			d.direction = move
			d.count, _ = strconv.Atoi(number)
			route = append(route, d)
		}
	}

	return route
}

func main() {
	tiles := LoadMap("../sample_map.txt")
	route := LoadRoute("../sample_route.txt")
	PrintMap(tiles)
	PrintRoute(route)
}
