package main

import (
	"fmt"
	"os"
	"strings"
)

func preprocess_map(antenna_locations string) ([][]string, map[string][][]int) {
	var antenna_map [][]string
	antenna_locs := make(map[string][][]int)
	for r_i, row := range strings.Split(antenna_locations, "\n") {
		antenna_map = append(antenna_map, strings.Split(row, ""))
		for c_i, char := range strings.Split(row, "") {
			if char != "." {
				_, ok := antenna_locs[char]
				if !ok {
					antenna_locs[char] = [][]int{{r_i, c_i}}
				} else {
					antenna_locs[char] = append(antenna_locs[char], []int{r_i, c_i})
				}
			}
		}
	}
	return antenna_map, antenna_locs
}

func in_range(antenna_map [][]string, location []int) bool {
	if location[0] >= 0 && location[1] >= 0 && location[0] < len(antenna_map) && location[1] < len(antenna_map[0]) {
		return true
	}
	return false
}

func add_antinodes(antenna_map [][]string, location []int, slope []int, op int, until_ubound bool) int {
	n_new_added := 0
	location = []int{location[0] + op*slope[0], location[1] + op*slope[1]}
	for {
		if in_range(antenna_map, location) {
			// Part 1 -- antennas are only antinodes if other antennas cause it.
			// Part 2 -- all antennas are antinodes, so we ignore antennas.
			if (!until_ubound && antenna_map[location[0]][location[1]] != "#") || (until_ubound && antenna_map[location[0]][location[1]] == ".") {
				antenna_map[location[0]][location[1]] = "#"
				n_new_added++
			}
		} else {
			return n_new_added
		}

		if until_ubound {
			location = []int{location[0] + op*slope[0], location[1] + op*slope[1]}
		} else {
			return n_new_added
		}
	}
}

func compute_n_unique_antinodes(antenna_map [][]string, antenna_locs map[string][][]int, until_ubound bool) int {
	n_unique_antinodes := 0
	n_antenna_locs := 0
	for _, locations := range antenna_locs {
		n_antenna_locs += len(locations)
		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if i != j {
					slope := []int{locations[i][0] - locations[j][0], locations[i][1] - locations[j][1]}
					n_unique_antinodes += add_antinodes(antenna_map, locations[j], slope, -1, until_ubound)
					n_unique_antinodes += add_antinodes(antenna_map, locations[i], slope, 1, until_ubound)
				}
			}
		}
	}

	if until_ubound {
		n_unique_antinodes += n_antenna_locs
	}

	return n_unique_antinodes
}

func copy_map(matrix [][]string) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	antenna_map, antenna_locs := preprocess_map(string(file))
	fmt.Printf("Part 1: %d\n", compute_n_unique_antinodes(copy_map(antenna_map), antenna_locs, false))
	fmt.Printf("Part 2: %d\n", compute_n_unique_antinodes(copy_map(antenna_map), antenna_locs, true))
}
