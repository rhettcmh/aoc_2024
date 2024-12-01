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

	// Read columns into slices
	scanner := bufio.NewScanner(file)
	var column_a []int
	var column_b []int
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num_a, _ := strconv.Atoi(nums[0])
		column_a = append(column_a, num_a)

		num_b, _ := strconv.Atoi(nums[1])
		column_b = append(column_b, num_b)
	}

	// Sort slices into smallest -> largest
	slices.Sort(column_a)
	slices.Sort(column_b)

	// Compute the differences
	solution := 0.0
	if len(column_a) == len(column_b) {
		for i := range column_a {
			solution += math.Abs(float64(column_a[i] - column_b[i]))
		}
	} else {
		panic("The input data is malformed - column lengths are not equal.")
	}

	fmt.Println("The solution is:", int(solution))

}
