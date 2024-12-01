package main

import (
	"bufio"
	"fmt"
	"os"
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
	column_a := make(map[int]int)
	column_b := make(map[int]int)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num_a, _ := strconv.Atoi(nums[0])
		column_a[num_a] = num_a

		num_b, _ := strconv.Atoi(nums[1])
		column_b[num_b] += 1
	}

	// Sort slices into smallest -> largest
	sum := 0.0
	for number, _ := range column_a {
		val, ok := column_b[number]
		if ok {
			sum += float64(number * val)
		}
	}
	fmt.Println("The solution is:", int(sum))
}
