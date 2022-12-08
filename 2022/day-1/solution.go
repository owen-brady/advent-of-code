package main

import (
	"fmt"
	"log"
	"math"
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

	// Iterate over strings, treating empty strings as the end of an elf's list
	currCalories := 0
	maxCalories := 0

	for i := 0; i < len(split); i++ {
		if split[i] == "" {
			// end of elf's list, increment and track next elf
			// study assignment
			maxAsFloat := math.Max(float64(currCalories), float64(maxCalories))
			maxCalories = int(math.Round(maxAsFloat))
			currCalories = 0

		} else {
			// elf has more snacks to add
			// study, decide on error handling
			increemnt, _ := strconv.ParseInt(split[i], 10, 0)
			currCalories += int(increemnt)
		}
	}

	maxAsFloat := math.Max(float64(currCalories), float64(maxCalories))
	maxCalories = int(math.Round(maxAsFloat))

	fmt.Println(maxCalories)
}
