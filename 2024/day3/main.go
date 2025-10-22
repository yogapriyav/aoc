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

	file, err := os.Open("data")
	if err != nil {
		fmt.Println("Could not open file")
	}

	if file == nil {
		fmt.Println("File opened is empty")
	}

	pattern := `mul\([0-9]+,[0-9]+\)|do\(\)|don\'t\(\)`
	patternRegex := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
	total := 0
	pause := false
	for scanner.Scan() {
		line := scanner.Text()
		match := patternRegex.FindAllString(line, -1)

		lineSum := 0
		for _, v := range match {

			if v == "don't()" {
				if !pause {
					pause = true
				}
				continue
			}

			if v == "do()" {
				if pause {
					pause = false
				}
				continue
			}

			if !pause {
				v1 := strings.Split(v, ",")[0]
				v2 := strings.Split(v, ",")[1]
				n1, _ := strconv.Atoi(strings.Split(v1, `(`)[1])
				n2, _ := strconv.Atoi(strings.Split(v2, `)`)[0])
				lineSum = lineSum + mul(n1, n2)
			}
		}

		total = total + lineSum
		fmt.Println("total: ", total)

	}

	fmt.Println("Total is:", total)

}

func mul(a int, b int) int {
	return a * b
}
