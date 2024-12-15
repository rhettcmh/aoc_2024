package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func iterate_stone(stone int, iters_remaining int) int {
	cache_key := fmt.Sprintf("%d,%d", stone, iters_remaining)
	if value, hit := cache[cache_key]; hit {
		return value
	}

	if iters_remaining == 0 {
		return 1
	}

	n_stones := 0
	stone_str := strconv.Itoa(stone)
	if stone == 0 {
		n_stones += iterate_stone(1, iters_remaining-1)
	} else if len(stone_str)%2 == 0 {
		stone_a, _ := strconv.Atoi(stone_str[:len(stone_str)/2])
		stone_b, _ := strconv.Atoi(stone_str[len(stone_str)/2:])
		n_stones += iterate_stone(stone_a, iters_remaining-1)
		n_stones += iterate_stone(stone_b, iters_remaining-1)
	} else {
		n_stones += iterate_stone(stone*2024, iters_remaining-1)
	}

	cache[cache_key] = n_stones
	return n_stones
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for part_i, n_blinks := range []int{25, 75} {
		n_deriative_stones := 0
		for _, stone := range strings.Split(string(file), " ") {
			stone_int, _ := strconv.Atoi(stone)
			n_deriative_stones += iterate_stone(stone_int, n_blinks)
		}
		fmt.Printf("Part %d: %d\n", part_i+1, n_deriative_stones)
	}
}
