package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var trees [][]int
	// count := 0
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scanner.Text()
		var newRow []int
		for _, i := range scanner.Text() {
			treeHeight, _ := strconv.Atoi(string(i))
			newRow = append(newRow, treeHeight)
		}
		trees = append(trees, newRow)
	}

	cc := (len(trees[0]) * 2) + (len(trees)-2)*2
	highestScenicScore := 0

	for row := 1; row < len(trees)-1; row++ {
		for col := 1; col < len(trees[0])-1; col++ {

			isVisible := checkVisibility(row, col, trees)
			// highestScenicScore
			temp := calcScore(row, col, trees)
			if temp > highestScenicScore {
				highestScenicScore = temp
			}
			if isVisible {
				cc++
			}
		}
	}
	fmt.Println("RESULT1", cc)
	fmt.Println("RESULT2", highestScenicScore)

}

func checkVisibility(x int, y int, forest [][]int) bool {
	n, o, s, w := true, true, true, true

	for i := 0; i < len(forest[0]); i++ {
		if y == i {
			continue
		} else if forest[x][y] <= forest[x][i] {
			if i < y {
				w = false
			} else {
				o = false
			}
		}
	}

	for j := 0; j < len(forest); j++ {
		if x == j {
			continue
		} else if forest[x][y] <= forest[j][y] {

			if x < j {
				n = false
			} else {
				s = false
			}
		}
	}
	if forest[x][y] == 1 {
		fmt.Println(n, o, s, w)
	}

	return n || o || s || w
}

func calcScore(x int, y int, forest [][]int) int {
	n, o, s, w := 0, 0, 0, 0
	treeHeight := forest[x][y]

	for i := y + 1; i < len(forest[0]); i++ {
		if treeHeight <= forest[x][i] {
			o++
			break
		} else {
			o++
		}
	}

	for i := y - 1; i >= 0; i-- {
		if treeHeight <= forest[x][i] {
			w++
			break
		} else {
			w++
		}
	}

	for i := x - 1; i >= 0; i-- {
		if treeHeight <= forest[i][y] {
			n++
			break
		} else {
			n++
		}
	}

	for i := x + 1; i < len(forest); i++ {
		if treeHeight <= forest[i][y] {
			s++
			break
		} else {
			s++
		}
	}

	if treeHeight == 9 {
		printVisibility(n, o, s, w, treeHeight)
	}
	if treeHeight == 1 {
		printVisibility(n, o, s, w, treeHeight)
	}

	return n * o * s * w
}

func printVisibility(n int, o int, s int, w int, treeHeight int) {
	fmt.Printf("%d'  %d\n", treeHeight, n)
	fmt.Println("    |")
	fmt.Printf("%d--   --%d\n", w, o)
	fmt.Println("    |")
	fmt.Printf("    %d\n", s)
}
