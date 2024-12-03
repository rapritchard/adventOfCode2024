package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the text file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define the regex pattern with capturing groups
	re := regexp.MustCompile(`(mul\(\s*(\d+)\s*,\s*(\d+)\s*\))|(do\(\))|(don't\(\))`)

	total := 0
	mulEnabled := true

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all matches in the current line
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[1] != "" { // mul() match
				if mulEnabled {
					total += convertToInt(match[2]) * convertToInt(match[3])
				}
			} else if match[4] != "" { // do() match
				mulEnabled = true
			} else if match[5] != "" { // don't() match
				mulEnabled = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(total)
}

func convertToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return num
}
