package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

type Elve struct {
	carrying int
}

var Elves []Elve

func main() {
	// f, err := os.Open("input.txt") //input for the first challange
	f, err := os.Open("input2.txt") //input for the second challange
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	result1 := 0
	result2 := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result1 += getResult(scanner.Text())
		result2 += getResult2(scanner.Text())
	}
	fmt.Printf("1: %d\n2: %d\n", result1, result2)
}

func getResult(s string) int {

	var sign = make(map[string]string)
	sign["A"] = "X"
	sign["B"] = "Y"
	sign["C"] = "Z"

	var base_score = make(map[string]int)
	base_score["X"] = 1
	base_score["Y"] = 2
	base_score["Z"] = 3
	ns := strings.Replace(s, s[0:1], sign[s[0:1]], 1)

	switch ns {
	case "X X", "Y Y", "Z Z":
		return 3 + base_score[ns[len(ns)-1:]]
	case "X Y", "Y Z", "Z X":
		return 6 + base_score[ns[len(ns)-1:]]
	default:
		return base_score[ns[len(ns)-1:]]
	}
}

func getResult2(s string) int {
	var outcome = make(map[string]string)
	var win = make(map[string]string)
	var draw = make(map[string]string)
	var lose = make(map[string]string)
	// A rock
	// B paper
	// C sissors

	win["A"] = "Y"
	win["B"] = "Z"
	win["C"] = "X"

	draw["A"] = "X"
	draw["B"] = "Y"
	draw["C"] = "Z"

	lose["A"] = "Z"
	lose["B"] = "X"
	lose["C"] = "Y"

	outcome["X"] = lose[s[0:1]]
	outcome["Y"] = draw[s[0:1]]
	outcome["Z"] = win[s[0:1]]

	ns := strings.Replace(s, s[len(s)-1:], outcome[s[len(s)-1:]], 1)
	return getResult(ns)
}

// This strategy guide predicts and recommends the following:

//     In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
//     In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
//     The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.

// In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).
