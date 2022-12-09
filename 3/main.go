package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	firstChallange()
	secondClallange()

}

func firstChallange() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	result := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result += findItem(scanner.Text())
	}
}

func secondClallange() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	result := 0
	scanner := bufio.NewScanner(f)
	counter := 0
	var group []string
	var emptyArray []string
	for scanner.Scan() {
		if counter == 2 {
			group = append(group, scanner.Text())
			result += findGroupBadge(group)
			group = emptyArray
			counter = 0
		} else {
			group = append(group, scanner.Text())
			counter++
		}
	}
	fmt.Printf("Result2: %d\n", result)
}

func findGroupBadge(s []string) int {
	for i := 0; i < len(s[0]); i++ {
		if strings.ContainsAny(s[1], string(s[0][i])) {
			if strings.ContainsAny(s[2], string(s[0][i])) {
				return getItemScore(s[0][i])
			}
		}
	}
	return 0
}

func findItem(s string) int {
	fmt.Println(s)
	for i := 0; i < len(s)/2; i++ {
		for j := len(s) / 2; j < len(s); j++ {
			if s[i] == s[j] {
				return getItemScore(s[i])
			}
		}
	}
	return 0
}

func getItemScore(c byte) int {
	ascii := int(c)
	if ascii > 96 {
		fmt.Println(c, string(c), ascii-97+1)
		return ascii - 97 + 1
	} else {
		fmt.Println(c, string(c), ascii-65+27)
		return ascii - 65 + 27
	}
}
