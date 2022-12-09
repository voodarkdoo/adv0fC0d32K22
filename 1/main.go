package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strconv"
)

func main() {
	getTotalCalories()
	getElvesWithMoreFood()
}

func getTotalCalories() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	result := 0
	caloriesOfOneElves := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			caloriesOfOneElves = 0
		} else {
			var singleMealCalories, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			caloriesOfOneElves += singleMealCalories
		}
		if caloriesOfOneElves > result {
			result = caloriesOfOneElves
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Total calories: %d\n", result)
}

type Elve struct {
	carrying int
}

var Elves []Elve

func getElvesWithMoreFood() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	result := 0
	caloriesOfOneElves := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			var newElve Elve
			newElve.carrying = caloriesOfOneElves
			Elves = append(Elves, newElve)

			caloriesOfOneElves = 0
		} else {
			var singleMealCalories, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			caloriesOfOneElves += singleMealCalories
		}
		if caloriesOfOneElves > result {
			result = caloriesOfOneElves
		}
	}
	sortElves(Elves)
	caloriesSum := 0

	for i := 0; i < 3; i++ {
		caloriesSum += Elves[i:][0].carrying
	}

	fmt.Printf("Total calories of 3 elves with more food: %d\n", caloriesSum)
}

func sortElves(elves []Elve) []Elve {
	for i := 0; i < len(elves); i++ {
		for j := 0; j < len(elves); j++ {
			if elves[i].carrying > elves[j].carrying {
				swapElves(i, j, elves)
			}
		}
	}
	return elves
}

func swapElves(i int, j int, elves []Elve) {
	var temp = elves[i]
	elves[i] = elves[j]
	elves[j] = temp
}
