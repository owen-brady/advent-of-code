package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Get file path
	dir, _ := filepath.Abs(".")
	// file := "2022/day-1/simple-input.txt"
	file := "2022/day-1/input.txt"
	path := dir + "/" + file

	// Read content
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert input to array of strings
	split := strings.Split(string(content), "\n")

	// Keep track of top three elves' total calories
	topThree := [3]int{0, 0, 0}

	// Iterate over strings, treating empty strings as the end of an elf's list
	currCalories := 0

	for i := 0; i < len(split); i++ {
		if split[i] == "" {
			// end of an elf's list, compare and move to next elf
			topThree = updateCalorieList(topThree, currCalories)
			currCalories = 0

		} else {
			// elf has more snacks to add
			increment, _ := strconv.ParseInt(split[i], 10, 0)
			currCalories += int(increment)
		}
	}

	fmt.Println("Top Calories: " + strconv.Itoa(topThree[0]))
	fmt.Println("Total Calories: " + strconv.Itoa((topThree[0] + topThree[1] + topThree[2])))
}

func updateCalorieList(topThree [3]int, curr int) [3]int {
	if curr > topThree[0] {
		topThree[2] = topThree[1]
		topThree[1] = topThree[0]
		topThree[0] = curr
	} else if curr > topThree[1] {
		topThree[2] = topThree[1]
		topThree[1] = curr
	} else if curr > topThree[2] {
		topThree[2] = curr
	}

	return topThree
}
