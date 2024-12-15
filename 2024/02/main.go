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

func is_safe(values []string) bool {
	increasing := false
	for i := range len(values) - 1 {
		a, _ := strconv.Atoi(values[i+1])
		b, _ := strconv.Atoi(values[i])
		diff := a - b

		abs_diff := math.Abs(float64(diff))
		if i == 0 {
			increasing = diff > 0
		}

		// Test bounds
		if (abs_diff < 1) || (abs_diff > 3) || (increasing != (diff > 0)) {
			return false
		}
	}
	return true
}

func is_mostly_safe(values []string) bool {
	for i := range len(values) {
		values_less_i := slices.Clone(values)
		values_less_i = append(values_less_i[:i], values_less_i[i+1:]...)
		if is_safe(values_less_i) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var n_safe int
	var n_mostly_safe int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		if is_safe(values) {
			n_safe++
		}

		if is_mostly_safe(values) {
			n_mostly_safe++
		}
	}
	fmt.Printf("N safe: %d\n", n_safe)
	fmt.Printf("N mostly safe: %d\n", n_mostly_safe)
}
