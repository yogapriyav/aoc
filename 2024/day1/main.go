package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("data")

	defer func() {
		file.Close()
	}()

	if err != nil {
		fmt.Println("Error reading the file.")
	}

	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "   ")

		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])

		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	// countMap for # of times number in list1 occurs in list2
	countMap := make(map[int]int)
	for _, num := range list2 {
		countMap[num]++
	}

	// similarity score
	score := 0
	for _, v := range list1 {
		score = score + v*countMap[v]
	}
	fmt.Println("Similarity Score is: ", score)

	// Sort list
	sort.Ints(list1)
	sort.Ints(list2)

	// Find distance and total
	total := 0
	for i, v := range list1 {
		diff := v - list2[i]
		if diff < 0 {
			diff = -diff
		}
		total = total + diff
	}

	fmt.Println("total distance: ", total)

}
