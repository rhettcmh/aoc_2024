package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input_str := string(file)

	part_1 := 0
	part_2 := 0
	ignore_mul := false
	r, _ := regexp.Compile(`(mul\((\d+),(\d+)\))|(do\(\))|(don\'t\(\))`)
	for _, match := range r.FindAllStringSubmatch(input_str, -1) {
		if match[0] == "do()" {
			ignore_mul = false
		} else if match[0] == "don't()" {
			ignore_mul = true
		} else {
			a, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			part_1 += a * b
			if ignore_mul == false {
				part_2 += a * b
			}
		}

	}
	fmt.Printf("Part 1: %d\n", part_1)
	fmt.Printf("Part 2: %d\n", part_2)
}
