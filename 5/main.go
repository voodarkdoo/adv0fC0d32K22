package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	firstChallange()
	// secondClallange()

}

type Command struct {
	move int
	from int
	to   int
}

func firstChallange() {

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(f)

	var commands []Command
	var arr [][]string

	for scanner.Scan() {

		if strings.HasPrefix(scanner.Text(), "move") {
			if strings.HasPrefix(scanner.Text(), "move") {
				var command Command
				num := re.FindAll([]byte(scanner.Text()), -1)
				command.move, _ = strconv.Atoi(string(num[0]))
				command.from, _ = strconv.Atoi(string(num[1]))
				command.to, _ = strconv.Atoi(string(num[2]))
				command.from--
				command.to--

				commands = append(commands, command)
			}
		}
		if strings.HasPrefix(scanner.Text(), "   ") ||
			strings.HasPrefix(scanner.Text(), "[") {
			var temp []string
			for i := 1; i < len(scanner.Text()); i = i + 4 {
				temp = append(temp, string(scanner.Text()[i]))
			}
			arr = append(arr, temp)
			temp = nil
		}

	}

	start := processArr(arr)
	for _, command := range commands {
		fmt.Println("Command:", command)
		for i, c := range start {
			fmt.Println(i, len(c), c)
		}
		start = moveCargo(start, command.move, command.from, command.to)
	}
	for _, ss := range start {
		fmt.Print(ss[len(ss)-1:])
	}
}

func processArr(arr [][]string) []string {
	var rows, cols = len(arr), len(arr[0])
	var result []string

	var temp string
	for i := 0; i <= cols-1; i++ {
		for j := rows - 1; j >= 0; j-- {
			temp = temp + string(arr[j][i])
		}
		result = append(result, strings.Trim(temp, " "))
		temp = ""
	}
	return result
}

func moveCargo(s []string, move int, from int, to int) []string {

	temp := s[from][len(s[from])-move:]
	// temp = reverse(temp) // toggle comment for first versionn b.
	s[from] = s[from][:len(s[from])-move]
	s[to] = s[to] + temp

	return s
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
