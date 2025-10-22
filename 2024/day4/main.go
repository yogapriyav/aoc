package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.Open("data")
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	// Build the 2D matrix
	count := 0
	scanner := bufio.NewScanner(file)
	grid := make([]string, 0)
	hCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)

		hCount = hCount + countPattern(line)
	}

	// Vertical Count
	verGrid := make([]string, 0)
	vCount := 0
	for col := 0; col < len(grid[0]); col++ {
		var sb strings.Builder
		for row := 0; row < len(grid); row++ {
			sb.WriteByte(grid[row][col])
		}
		verGrid = append(verGrid, sb.String())

		vCount = vCount + countPattern(verGrid[col])
	}

	// Diagonal count
	diagGrid := make([]string, 0)
	dCount := 0

	// Diagonal right count
	// Top row, move right
	for startCol := 0; startCol < len(grid[0]); startCol++ {
		var sb strings.Builder
		row := 0
		col := startCol
		for row < len(grid) && col < len(grid[0]) {
			sb.WriteByte(grid[row][col])
			row++
			col++
		}
		diagGrid = append(diagGrid, sb.String())
	}

	// Left col, move down (skip 0,0 already done above)
	for startRow := 1; startRow < len(grid); startRow++ {
		var sb strings.Builder
		row := startRow
		col := 0
		for row < len(grid) && col < len(grid[0]) {
			sb.WriteByte(grid[row][col])
			row++
			col++
		}
		diagGrid = append(diagGrid, sb.String())
	}

	// Diagonal left count
	// Top row, move left
	for startCol := 0; startCol < len(grid[0]); startCol++ {
		var sb strings.Builder
		row := 0
		col := startCol
		for row < len(grid) && col >= 0 {
			sb.WriteByte(grid[row][col])
			row++
			col--
		}
		diagGrid = append(diagGrid, sb.String())
	}

	// Right col, move down (skip top right corner, already done)
	for startRow := 1; startRow < len(grid); startRow++ {
		var sb strings.Builder
		row := startRow
		col := len(grid[0]) - 1
		for row < len(grid) && col >= 0 {
			sb.WriteByte(grid[row][col])
			row++
			col--
		}
		diagGrid = append(diagGrid, sb.String())
	}

	// Total Diagonal count
	for _, value := range diagGrid {
		dCount = dCount + countPattern(value)
	}

	// Total count
	count = hCount + vCount + dCount

	fmt.Println("Total count (horizontal, vertical, diagonal): ", count)

	// X-MAS count
	// isXPattern logic looks 2 rows and cols ahead
	// Hence loop bound has -2
	xCount := 0
	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			if isXPattern(row, col, grid) {
				xCount++
			}
		}
	}
	fmt.Println("Total X-MAS count: ", xCount)
}

func countPattern(s string) int {

	pattern1 := "XMAS"
	pattern2 := "SAMX"
	patternRegex1 := regexp.MustCompile(pattern1)
	patternRegex2 := regexp.MustCompile(pattern2)

	len1 := len(patternRegex1.FindAllString(s, -1))
	len2 := len(patternRegex2.FindAllString(s, -1))

	return len1 + len2
}

func isXPattern(row int, col int, grid []string) bool {
	if string(grid[row+1][col+1]) != "A" {
		return false
	}

	/*
		---
		M . M
		. A .
		S . S
		---
		M . S
		. A .
		M . S
	*/
	if string(grid[row][col]) == "M" {
		if string(grid[row+2][col+2]) != "S" {
			return false
		}

		if (string(grid[row][col+2]) == "M" && string(grid[row+2][col]) == "S") || (string(grid[row][col+2]) == "S" && string(grid[row+2][col]) == "M") {
			return true
		}
	}

	/*
		---
		S . M
		. A .
		S . M
		---
		S . S
		. A .
		M . M
	*/

	if string(grid[row][col]) == "S" {
		if string(grid[row+2][col+2]) != "M" {
			return false
		}

		if (string(grid[row][col+2]) == "S" && string(grid[row+2][col]) == "M") || (string(grid[row][col+2]) == "M" && string(grid[row+2][col]) == "S") {
			return true
		}
	}

	return false
}
