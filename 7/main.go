package main

import (
	"bufio"
	"fmt"
	"os"
	// "regexp"
	files "seven/files"
	"strconv"
	"strings"
)

func main() {

	root := files.Dir{
		Name:   "root",
		Parent: nil,
		Size:   0,
	}
	currentDir := &root
	filesystem := []*files.Dir{}
	filesystem = append(filesystem, currentDir)

	// re := regexp.MustCompile(`\d`)
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	count := 0
	cd := 0
	cd_up := 0
	cd_down := 0
	for scanner.Scan() {
		count++
		input := scanner.Text()
		if strings.HasPrefix(input, "dir") {
			newDir := files.InsertDir(currentDir, strings.Split(input, " ")[1])
			filesystem = append(filesystem, newDir)
			continue
		} else if strings.HasPrefix(input, "$ cd") {
			cd++
			fmt.Println(count, input)
			if input == "$ ls" {
				fmt.Println(input == "$ ls")
				continue
			} else if input == "$ cd /" {
				fmt.Println(input == "$ cd /")
				continue
			} else if input == "$ cd .." {
				cd_up++
				// fmt.Println("-->", currentDir.Name)
				currentDir = currentDir.Parent
				// fmt.Println("Problem in line", count)
			} else {
				cd_down++
				// fmt.Println("--", input)
				currentDir = files.FindDir(filesystem, strings.Split(input, " ")[2])
				// fmt.Println("++", currentDir.Name)
			}
			continue
		} else {
			fmt.Println(count, input)
			size, _ := strconv.Atoi(strings.Split(input, " ")[0])
			files.InsertFile(currentDir, size)
		}
	}

	sum := 0
	for _, dir := range filesystem {
		if dir.Size < 100000 {
			sum += dir.Size
		}
		// fmt.Println(dir.Name, dir.Size)
	}
	fmt.Println("RESULT", sum, "filesystem", len(filesystem), cd, cd_up, cd_down)
	fmt.Println("filesystem length:", len(filesystem))
	fmt.Println("cd:", cd)
}
