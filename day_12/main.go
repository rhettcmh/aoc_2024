package main

import (
	"fmt"
	"os"
	"strings"
)

// Ordered: Up, Left, Down, Right
var directions = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func in_range(garden [][]string, row int, col int) bool {
	return row >= 0 && row < len(garden) && col >= 0 && col < len(garden[0])
}

func count_corners(garden [][]string, row int, col int) int {
	n_corners := 0
	for i := 0; i < len(directions); i++ {
		dir_i := directions[i]
		dir_j := directions[(i+1)%len(directions)]

		var edge_a, edge_b, diag string
		if in_range(garden, row+dir_i[0], col+dir_i[1]) {
			edge_a = garden[row+dir_i[0]][col+dir_i[1]]
		}
		if in_range(garden, row+dir_j[0], col+dir_j[1]) {
			edge_b = garden[row+dir_j[0]][col+dir_j[1]]
		}
		if in_range(garden, row+dir_i[0]+dir_j[0], col+dir_i[1]+dir_j[1]) {
			diag = garden[row+dir_i[0]+dir_j[0]][col+dir_i[1]+dir_j[1]]
		}

		if (edge_a != garden[row][col] && edge_b != garden[row][col]) || (edge_a == garden[row][col] && edge_b == garden[row][col] && diag != garden[row][col]) {
			n_corners++
		}
	}
	return n_corners
}

func compute_perimeter_and_area(garden [][]string, visited_plots [][]bool, row int, col int, letter string) (int, int, int) {
	// Base case: outside map, not on letter or already visited.
	if !in_range(garden, row, col) || garden[row][col] != letter || visited_plots[row][col] {
		return 0, 0, 0
	}

	area := 1
	perimeter := 0
	num_sides := count_corners(garden, row, col)
	visited_plots[row][col] = true
	for _, move := range directions {
		next_row, next_col := move[0]+row, move[1]+col
		if !in_range(garden, next_row, next_col) || garden[next_row][next_col] != letter {
			perimeter++
		}
		p, ns, a := compute_perimeter_and_area(garden, visited_plots, next_row, next_col, letter)
		perimeter += p
		num_sides += ns
		area += a

	}
	return perimeter, num_sides, area
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var garden [][]string
	var visited_plots [][]bool
	for _, row := range strings.Split(string(file), "\n") {
		garden = append(garden, strings.Split(row, ""))
		visited_plots = append(visited_plots, make([]bool, len(row)))
	}

	part_1_cost := 0
	part_2_cost := 0
	for row := 0; row < len(garden); row++ {
		for col := 0; col < len(garden[0]); col++ {
			if !visited_plots[row][col] {
				p, ns, a := compute_perimeter_and_area(garden, visited_plots, row, col, garden[row][col])
				part_1_cost += p * a
				part_2_cost += ns * a
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part_1_cost)
	fmt.Printf("Part 2: %d\n", part_2_cost)
}
