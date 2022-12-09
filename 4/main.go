package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strconv"
)

func main() {

	firstChallange()
	// secondClallange()

}

func firstChallange() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	result := 0
	scanner := bufio.NewScanner(f)
	cc := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		cc++
		var temp []int
		for _, item := range s {
			for _, n := range strings.Split(item, "-") {
				number, err := strconv.Atoi(n)
				if err != nil {
					fmt.Println("ERRORE:", err)
				} else {
					temp = append(temp, number)
				}
			}
		}
		// result += isFullOvelap(temp)
		if overlapAtAll(temp) {
			result += 1
		} else {
			result += 0
		}

	}
	fmt.Println(cc, result)
}

func isFullOvelap(r []int) bool {
	if (r[0] <= r[2] && r[1] >= r[3]) || (r[0] >= r[2] && r[1] <= r[3]) {
		return true
	} else {
		return false
	}
}

func overlapAtAll(r []int) bool {
	if isFullOvelap(r) {
		return true
	} else if (r[1] >= r[2] && r[0] <= r[3]) || (r[2] >= r[1] && r[3] <= r[0]) {
		return true
	} else {
		return false
	}
}
