package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	// Read columns into slices & counts
	scanner := bufio.NewScanner(file)
	var column_a []int
	var column_b []int
	counts := make(map[int]int)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num_a, _ := strconv.Atoi(nums[0])
		column_a = append(column_a, num_a)

		num_b, _ := strconv.Atoi(nums[1])
		column_b = append(column_b, num_b)
		counts[num_b]++
	}

	// Sorted slices
	slices.Sort(column_a)
	slices.Sort(column_b)

	part_1 := 0.0
	part_2 := 0.0
	for i, number := range column_a {
		part_1 += math.Abs(float64(column_a[i] - column_b[i]))
		part_2 += float64(number * counts[number])
	}

	fmt.Println("Part 1:", int(part_1))
	fmt.Println("Part 2:", int(part_2))
}
