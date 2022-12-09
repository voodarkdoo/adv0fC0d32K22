package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	for i := 0; i < len(input)-13; i++ {
		if check(input[i : i+14]) {
			fmt.Println("RES:", i+14)
			break
		}
	}
}

func check(a string) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}
