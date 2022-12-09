package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	name   string
	parent *Dir
	size   int
	dirs   []Dir
}

func main() {

	var start Dir

	start.name = "start"
	start.parent = nil
	start.size = 0
	start.dirs = nil

	currDir := start

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var input string
	count := 0
	for scanner.Scan() {
		input = scanner.Text()
		if strings.HasPrefix(input, "$ cd /") {
			// insertDir(currDir, input)
		}
		if strings.HasPrefix(input, "dir") {
			insertDir(&currDir, input)
		}
		if !strings.HasPrefix(input, "$") && !strings.HasPrefix(input, "dir") {
			insertFile(&currDir, input)
		}
		if strings.HasPrefix(input, "$ cd ..") {
			if currDir.parent != nil {
				currDir = *currDir.parent
			}
		}
		if strings.HasPrefix(input, "$ cd ") && !strings.HasPrefix(input, "$ cd ..") {
			currDir = findDir(currDir, input)
		}
		// fmt.Println(currDir.size)
	}
	fmt.Println("N dirs", count)
	printDirs(parent(currDir), "-")
}

func insertFile(d *Dir, s string) {
	num, _ := strconv.Atoi(strings.Split(s, " ")[0])
	d.size = d.size + num
	if d.parent != nil {
		d.parent.size = d.parent.size + num
	}
}

func insertDir(d *Dir, s string) {
	fmt.Println("start - INSERT", d, s)
	var dirs []Dir
	var newDir Dir
	newDir.name = strings.Split(s, " ")[1]
	newDir.size = 0
	newDir.parent = d
	newDir.dirs = dirs
	fmt.Println("new dir", newDir)
	d.dirs = append(d.dirs, newDir)
	fmt.Println("end - ISNERT", d)
}

func findDir(d Dir, s string) Dir {
	fmt.Println("FIND:", d, s)
	name := strings.Split(s, " ")[2]
	for _, dd := range d.dirs {
		fmt.Println(dd.name, "=>", name)
		if dd.name == name {
			return dd
		}
	}
	return d
}

func printDirs(d Dir, s string) int {
	fmt.Println(d)
	fmt.Print(s, d.name, d.size)
	s = s + s
	totSum := 0
	if d.size < 100000 {
		totSum = d.size
	}
	for _, dd := range d.dirs {
		sum := printDirs(dd, s)
		if sum < 100000 {
			totSum += sum
		}
	}
	return totSum
}

func parent(d Dir) Dir {
	if d.parent == nil {
		return d
	} else {
		return parent(*d.parent)
	}
}
