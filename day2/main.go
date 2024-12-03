package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readReports("input.txt")
	if err != nil {
		log.Fatal("Error reading reports:", err)
	}

	totalSafeReports := 0
	for _, report := range reports {
		if isSafeReport(report) {
			totalSafeReports++
		}
	}

	fmt.Println(totalSafeReports)
}

func readReports(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report, err := parseReport(line)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func parseReport(line string) ([]int, error) {
	fields := strings.Fields(line)
	report := make([]int, len(fields))
	for i, field := range fields {
		value, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		report[i] = value
	}
	return report, nil
}

func isSafeReport(report []int) bool {
	if len(report) < 2 {
		return false
	}

	if checkReport(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := make([]int, 0, len(report)-1)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		if checkReport(modifiedReport) {
			return true
		}
	}

	return false
}

func checkReport(report []int) bool {
	if len(report) < 2 {
		return false
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		diff := int(math.Abs(float64(report[i] - report[i-1])))
		if diff < 1 || diff > 3 {
			return false
		}
		if report[i] > report[i-1] {
			isDecreasing = false
		} else if report[i] < report[i-1] {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
