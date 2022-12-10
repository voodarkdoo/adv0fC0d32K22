package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
}

var positions []Pos
var head = Pos{0, 0}
var prev_head = Pos{0, 0}
var tail = Pos{0, 0}

func main() {

	f, _ := os.Open("test.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	positions = append(positions, tail)

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		direction := command[0]
		number, _ := strconv.Atoi(command[1])
		for x := 0; x < number; x++ {
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}
			if haveToMove(head, tail) {
				positions = append(positions, tail)
				tail = prev_head
			}
			prev_head = head
		}
	}
	for _, x := range positions {
		fmt.Println("-", x)
	}
	fmt.Println(len(positions))
}

func isNotIn(positions []Pos, start Pos) bool {
	for _, p := range positions {
		if start == p {
			return false
		}
	}
	return true
}

func haveToMove(t Pos, h Pos) bool {
	xx := (t.x - h.x)
	yy := t.y - h.y
	if math.Sqrt(float64((xx*xx)+(yy*yy))) > 1 {
		return true
	} else {
		return false
	}
}
