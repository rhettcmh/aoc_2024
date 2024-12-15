package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func is_int(val float64) bool {
	return val == float64(int(val))
}

func solveCramers(a_x, b_y, e, c_x, d_y, f int) (float64, float64) {
	/* Uses Cramer's rule to solve the set of equations
	ax + by = e
	cx + dy = f
	*/
	denom := a_x*d_y - c_x*b_y
	if denom == 0 {
		return -1, -1
	}

	x := float64(e*d_y-f*b_y) / float64(denom)
	y := float64(a_x*f-c_x*e) / float64(denom)

	if x < 0 || y < 0 || !is_int(x) || !is_int(y) {
		return -1, -1
	}

	return x, y
}

func solve(equation string, part_2 bool) float64 {
	eq_a_regex, _ := regexp.Compile(`A: X\+(\d+), Y\+(\d+)`)
	eq_a := eq_a_regex.FindAllStringSubmatch(equation, -1)[0]
	a_x, c_y := eq_a[1], eq_a[2]
	A_x, _ := strconv.Atoi(a_x)
	C_y, _ := strconv.Atoi(c_y)

	eq_b_regex, _ := regexp.Compile(`B: X\+(\d+), Y\+(\d+)`)
	eq_b := eq_b_regex.FindAllStringSubmatch(equation, -1)[0]
	b_x, d_y := eq_b[1], eq_b[2]
	B_x, _ := strconv.Atoi(b_x)
	D_y, _ := strconv.Atoi(d_y)

	target_eq_regex, _ := regexp.Compile(`Prize: X=(\d+), Y=(\d+)`)
	target_eq := target_eq_regex.FindAllStringSubmatch(equation, -1)[0]
	e, f := target_eq[1], target_eq[2]

	E, _ := strconv.Atoi(e)
	F, _ := strconv.Atoi(f)
	if part_2 {
		E += 10000000000000
		F += 10000000000000
	}

	A, B := solveCramers(A_x, B_x, E, C_y, D_y, F)
	if A == -1 || B == -1 {
		return 0
	}

	return 3*A + B
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	part_1_cost := 0.0
	part_2_cost := 0.0
	for _, equation := range strings.Split(string(file), "\n\n") {
		part_1_cost += solve(equation, false)
		part_2_cost += solve(equation, true)
	}
	fmt.Printf("Part 1: %d\n", int(part_1_cost))
	fmt.Printf("Part 2: %d\n", int(part_2_cost))
}
