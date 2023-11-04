package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name           string
	children       []*Directory
	calculatedSize int64
}

func parseDirectory(lines *[]string, currentDir *Directory) {
	if len(*lines) == 0 {
		return
	}

	line := (*lines)[0]
	*lines = (*lines)[1:]

	if line == "$ cd /" {
		return
	} else if strings.HasPrefix(line, "$ cd ") {
		dirName := strings.TrimPrefix(line, "$ cd ")
		if dirName == ".." {
			return
		}
		for _, child := range currentDir.children {
			if child.name == dirName {
				parseDirectory(lines, child)
				break
			}
		}
	} else if len(line) > 0 {
		fields := strings.Fields(line)
		if fields[0] == "dir" {
			newDir := &Directory{name: fields[1], children: []*Directory{}}
			currentDir.children = append(currentDir.children, newDir)
			parseDirectory(lines, newDir)
			currentDir.calculatedSize += newDir.calculatedSize
		} else {
			size, err := strconv.ParseInt(fields[0], 10, 64)
			if err == nil {
				currentDir.calculatedSize += size
			}
		}
	}
	parseDirectory(lines, currentDir)
}

func parseInput(input string) *Directory {
	root := &Directory{name: "/", children: []*Directory{}}
	lines := strings.Split(input, "\n")
	parseDirectory(&lines, root)
	return root
}

func findSmallestDirectory(dir *Directory, minSize int64) *Directory {
	var smallestDir *Directory = nil
	size := dir.calculatedSize

	if size > minSize {
		smallestDir = dir

		for _, child := range dir.children {
			childDir := findSmallestDirectory(child, minSize)
			if childDir != nil && (smallestDir == nil || childDir.calculatedSize < smallestDir.calculatedSize) {
				smallestDir = childDir
			}
		}
	}

	return smallestDir
}

func printDirectorySizes(dir *Directory, level int) {
	indent := strings.Repeat("  ", level)
	fmt.Printf("%s%s (%d)\n", indent, dir.name, dir.calculatedSize)

	for _, child := range dir.children {
		printDirectorySizes(child, level+1)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(inputBytes)

	root := parseInput(input)
	printDirectorySizes(root, 0)
	smallestDir := findSmallestDirectory(root, 8381165)
	fmt.Printf("Smallest Directory: %s (%d)\n", smallestDir.name, smallestDir.calculatedSize)
}
