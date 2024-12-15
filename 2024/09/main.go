package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Empty space is denoted as "-1"
type Chunk struct {
	data int
	size int
}

func read_into_memory(memory_map string, max_chunk_size int) []Chunk {
	memory := []Chunk{}
	for d_idx, data := range strings.Split(memory_map, "") {
		data_int, _ := strconv.Atoi(string(data))
		n_chunks := int(math.Ceil(float64(data_int) / float64(max_chunk_size)))
		for i := 0; i < n_chunks; i++ {
			data_size := min(data_int-i*max_chunk_size, max_chunk_size)
			if d_idx%2 == 0 {
				memory = append(memory, Chunk{d_idx / 2, data_size})
			} else {
				memory = append(memory, Chunk{-1, data_size})
			}
		}
	}
	return memory
}

func defrag_memory(memory []Chunk) []Chunk {
	for end_ptr := len(memory) - 1; end_ptr > 0; end_ptr-- {
		for start_ptr := 0; start_ptr <= end_ptr; start_ptr++ {
			if (memory[end_ptr].data != -1) && (memory[start_ptr].data == -1) && (memory[end_ptr].size <= memory[start_ptr].size) {
				// Remaining free space
				free_space := memory[start_ptr].size - memory[end_ptr].size

				// Memory to move
				chunk_to_move := memory[end_ptr]

				// Deallocat & resize memory
				memory[start_ptr] = Chunk{-1, free_space}
				memory[end_ptr] = Chunk{-1, chunk_to_move.size}

				// Move chunk to new location
				memory = slices.Insert(memory, start_ptr, chunk_to_move)
			}
		}
	}
	return memory
}

func checksum(memory []Chunk) int {
	value := 0
	idx := 0
	for _, chunk := range memory {
		for i := 0; i < chunk.size; i++ {
			if chunk.data != -1 {
				value += chunk.data * idx
			}
			idx++
		}
	}
	return value
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	memory := read_into_memory(string(file), 1)
	fmt.Printf("Part 1: %d\n", checksum(defrag_memory(memory)))

	memory = read_into_memory(string(file), 1e9)
	fmt.Printf("Part 2: %d\n", checksum(defrag_memory(memory)))
}
