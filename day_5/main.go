package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n\n")

	// Extract the page orderings
	// --> x|y: x must be printed before y
	// --> orders[x]: numbers that must come AFTER orders[x]
	orders := make(map[int][]int)
	for _, row := range strings.Split(string(input[0]), "\n") {
		order := strings.Split(string(row), "|")
		a, _ := strconv.Atoi(order[0])
		b, _ := strconv.Atoi(order[1])
		orders[a] = append(orders[a], b)
	}

	// If any b is in orders[a], this violates the ordering condition.
	cmp := func(a int, b int) int {
		values, ok := orders[a]
		if ok {
			for _, v := range values {
				if v == b {
					return -1
				}
			}
		}
		return 0
	}

	part_1_sum := 0
	part_2_sum := 0
	for _, pages := range strings.Split((string(input[1])), "\n") {
		// TODO: refactor to on-the-fly - removes orders map as well.
		var page_numbers []int
		for _, page := range strings.Split(pages, ",") {
			a, _ := strconv.Atoi(page)
			page_numbers = append(page_numbers, a)
		}

		if slices.IsSortedFunc(page_numbers, cmp) {
			part_1_sum += page_numbers[len(page_numbers)/2]
		} else {
			slices.SortFunc(page_numbers, cmp)
			part_2_sum += page_numbers[len(page_numbers)/2]
		}
	}

	fmt.Printf("Part 1: %d\n", part_1_sum)
	fmt.Printf("Part 2: %d\n", part_2_sum)
}
