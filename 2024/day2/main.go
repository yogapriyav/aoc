package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Read from a file
	file, err := os.Open("data")

	defer func() {
		_ = file.Close()
	}()

	if err != nil {
		fmt.Println("Error: Exiting..", err)
	}

	// Parse each line
	scanner := bufio.NewScanner(file)

	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		reportStr := strings.Fields(line)

		report := make([]int, 0, len(reportStr))
		for _, levelStr := range reportStr {
			n, _ := strconv.Atoi(levelStr)
			report = append(report, n)
		}

		if !isSafe(report) {

			for i := 0; i < len(report); i++ {
				newReport := removeElement(report, i)
				if isSafe(newReport) {
					safeCount++
					break
				}
			}
		} else {
			safeCount++
		}
	}
	fmt.Println("Safe count is:", safeCount)
}

// If condition check fails or direction check fails, it's not safe
func isSafe(report []int) bool {

	safe := false
	if !isMatchingCondition(report[0], report[1]) {
		return false
	}

	increasing := isIncreasing(report[0], report[1])

	for i := 2; i < len(report); i++ {
		if !isMatchingCondition(report[i-1], report[i]) {
			safe = false
			break
		}
		if isIncreasing(report[i-1], report[i]) != increasing {
			safe = false
			break
		}

		safe = true
	}

	return safe
}

// If diff between 2 adjacent numbers is < 1 or > 3
// condition does not match
func isMatchingCondition(curr int, next int) bool {
	diff := next - curr
	if diff < 0 {
		diff = -diff
	}
	if diff < 1 || diff > 3 {
		return false
	}

	return true
}

// Determines if 2 adjacent levels are in increasing order
func isIncreasing(curr int, next int) bool {
	if next < curr {
		return false
	}
	return true
}

// Remove an element at an index from a report
func removeElement(r []int, i int) []int {
	if i < 0 || i >= len(r) {
		return r
	}
	result := make([]int, 0, len(r)-1)
	result = append(result, r[:i]...)
	result = append(result, r[i+1:]...)
	return result
}
