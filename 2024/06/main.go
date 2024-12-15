package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func run_guard_route(initial_map [][]string, i int, j int) (int, []string) {
	directions_map := map[string][]int{
		">": {0, 1},
		"^": {-1, 0},
		"v": {1, 0},
		"<": {0, -1},
	}
	direction_policy := "^>v<"

	// New step count
	step_count := 0

	// Visited cfgs
	var visited []string

	// What is my current heading
	direction_idx := strings.Index(direction_policy, initial_map[i][j])

	for {
		// Which way am I facing?
		heading := directions_map[string(direction_policy[direction_idx])]

		// Record where I've been
		new_key := fmt.Sprintf("%d,%d,%d", i, j, direction_idx)
		if slices.Contains(visited, new_key) {
			return -1, visited
		} else {
			visited = append(visited, new_key)
		}

		// Where do I want to go next?
		row_i := i + heading[0]
		col_i := j + heading[1]

		// Last step out
		if row_i < 0 || row_i >= len(initial_map) || col_i < 0 || col_i >= len(initial_map[0]) {
			return step_count + 1, visited
		}

		// Try to step in front of you
		if initial_map[row_i][col_i] == "#" {
			direction_idx = (direction_idx + 1) % len(direction_policy)
		} else {
			if initial_map[i][j] != "%" {
				initial_map[i][j] = "%"
				step_count++
			}

			i += heading[0]
			j += heading[1]
		}
	}
}

func copy_guard(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func find_all_obstacles(guard_map [][]string, i int, j int, route []string) int {
	count := 0
	cycle_obs_locs := make(map[string]int)
	for _, step := range route {
		steps := strings.Split(step, ",")
		row, _ := strconv.Atoi(steps[0])
		col, _ := strconv.Atoi(steps[1])
		key := fmt.Sprintf("%d,%d", row, col)
		_, ok := cycle_obs_locs[key]
		if !ok {
			cycle_obs_locs[key] = 1
			if !(guard_map[row][col] == "#" || (i == row && j == col)) {
				tmp_guard_map := copy_guard(guard_map)
				tmp_guard_map[row][col] = "#"
				n_steps, _ := run_guard_route(tmp_guard_map, i, j)
				if n_steps == -1 {
					count++
				}

			}
		} else {

		}

	}
	return count

}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	map_rows := strings.Split(string(file), "\n")
	guard_map := make([][]string, len(map_rows))
	start_i, start_j := -1, -1
	for i, row := range map_rows {
		guard_map[i] = strings.Split(row, "")
		idx := strings.Index(row, "^")
		if idx != -1 {
			start_i = i
			start_j = idx
		}
	}

	n_steps, route := run_guard_route(copy_guard(guard_map), start_i, start_j)
	fmt.Println(n_steps)
	fmt.Println(find_all_obstacles(guard_map, start_i, start_j, route))

}
