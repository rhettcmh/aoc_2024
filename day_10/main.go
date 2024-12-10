package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func preprocess_input(topograph_input string) ([][]int, [][]int, [][]int) {
	var topographic_map [][]int
	var starting_locations [][]int
	var peak_locations [][]int

	for r_i, line := range strings.Split(topograph_input, "\n") {
		var row []int
		for c_i, elem := range strings.Split(line, "") {
			elem_int, _ := strconv.Atoi(elem)
			if elem_int == 0 {
				starting_locations = append(starting_locations, []int{r_i, c_i})
			} else if elem_int == 9 {
				peak_locations = append(peak_locations, []int{r_i, c_i})
			}
			row = append(row, elem_int)
		}
		topographic_map = append(topographic_map, row)
	}
	return topographic_map, starting_locations, peak_locations
}

func n_paths_to_peak(topographic_map [][]int, i int, j int, peak_i int, peak_j int) int {
	if i == peak_i && j == peak_j {
		return 1
	}

	n_paths := 0
	for _, move := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		new_i, new_j := move[0]+i, move[1]+j
		if new_i >= 0 && new_i < len(topographic_map) && new_j >= 0 && new_j < len(topographic_map[0]) {
			if topographic_map[new_i][new_j]-topographic_map[i][j] == 1 {
				n_paths += n_paths_to_peak(topographic_map, new_i, new_j, peak_i, peak_j)
			}
		}
	}
	return n_paths
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	topographic_map, starting_locations, peak_locations := preprocess_input(string(file))
	part_1_score := 0
	part_2_score := 0
	for _, start := range starting_locations {
		for _, peak := range peak_locations {
			peak_score := n_paths_to_peak(topographic_map, start[0], start[1], peak[0], peak[1])
			if peak_score > 0 {
				part_1_score++
			}
			part_2_score += peak_score
		}
	}
	fmt.Printf("Part 1: %d\n", part_1_score)
	fmt.Printf("Part 2: %d\n", part_2_score)
}
