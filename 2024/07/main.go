package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func brute_force_solution(current_value int, target_value int, numbers []int, i int, concat bool) int {
	if current_value == target_value {
		return 1
	} else if i >= len(numbers) {
		return 0
	}

	n_solutions := brute_force_solution(current_value+numbers[i], target_value, numbers, i+1, concat)
	n_solutions += brute_force_solution(current_value*numbers[i], target_value, numbers, i+1, concat)
	if concat {
		next_value, _ := strconv.Atoi(strconv.Itoa(current_value) + strconv.Itoa(numbers[i]))
		n_solutions += brute_force_solution(next_value, target_value, numbers, i+1, concat)
	}

	return n_solutions
}

func clip_to_one(x int) int {
	if x >= 1 {
		return 1
	}
	return 0
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	n_solutions_p1 := 0
	n_solutions_p2 := 0
	for _, round := range strings.Split(string(file), "\n") {
		target_value, _ := strconv.Atoi(strings.Split(round, ":")[0])
		var numbers []int
		for _, num := range strings.Split(strings.Split(round, ": ")[1], " ") {
			num, _ := strconv.Atoi(num)
			numbers = append(numbers, num)
		}

		n_solutions_p1 += clip_to_one(brute_force_solution(0, target_value, numbers, 0, false)) * target_value
		n_solutions_p2 += clip_to_one(brute_force_solution(0, target_value, numbers, 0, true)) * target_value
	}
	fmt.Printf("Part 1: %d\n", n_solutions_p1)
	fmt.Printf("Part 2: %d\n", n_solutions_p2)
}
