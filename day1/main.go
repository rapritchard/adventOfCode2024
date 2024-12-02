package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	// Initialise two empty slices
	var leftList, rightList []int
	totalDistance, similar := 0, 0

	// Open the file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line into two parts
		line := scanner.Text()
		parts := strings.Fields(line)

		// Convert the strings to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting left to integer", err)
			return
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting right to integer", err)
			return
		}

		// Append the integers to the slices
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	// check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	// sort slices by size of int
	sort.Ints(leftList)
	sort.Ints(rightList)

	dict := make(map[int]int)
	for i := range leftList {
		totalDistance += abs(leftList[i] - rightList[i])
		if _, exists := dict[leftList[i]]; !exists {
			dict[leftList[i]] = 0
		}
	}

	// Increment the value in dict if an element in rightList matches a key
	for _, value := range rightList {
		if _, exists := dict[value]; exists {
				dict[value]++
		}
	}

	for key, value := range dict {
		similar += key * value
	}
	
	fmt.Println("Total Distance:", totalDistance)
	fmt.Println("Similar:", similar)
}