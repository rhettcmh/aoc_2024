package main

import (
	"fmt"
	"os"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func search_horizontal(word_search [][]string, key_word string, i int, j int) int {
	if j <= len(word_search[0])-len(key_word) {
		if strings.Join(word_search[i][j:j+len(key_word)], "") == key_word {
			return 1
		}
	}
	return 0
}

func search_vertical(word_search [][]string, key_word string, i int, j int) int {
	if i <= len(word_search)-len(key_word) {
		for step := 0; step < len(key_word); step++ {
			if word_search[i+step][j] != string(key_word[step]) {
				return 0
			}
		}
		return 1
	}
	return 0
}

func search_neg_diagonal(word_search [][]string, key_word string, i int, j int) int {
	if (i <= len(word_search)-len(key_word)) && (j <= len(word_search[0])-len(key_word)) {
		for step := 0; step < len(key_word); step++ {
			if word_search[i+step][j+step] != string(key_word[step]) {
				return 0
			}
		}
		return 1
	}
	return 0
}

func search_pos_diagonal(word_search [][]string, key_word string, i int, j int) int {
	if (i >= len(key_word)-1) && (j <= len(word_search[0])-len(key_word)) {
		for step := 0; step < len(key_word); step++ {
			if word_search[i-step][j+step] != string(key_word[step]) {
				return 0
			}
		}
		return 1
	}
	return 0
}

func search_xmas(word_search [][]string, i int, j int) int {
	if (i > 0) && (j > 0) && (i < len(word_search)-1) && (j < len(word_search)-1) {
		if word_search[i][j] == "A" {
			if !((word_search[i-1][j-1] == "M" && word_search[i+1][j+1] == "S") || (word_search[i-1][j-1] == "S" && word_search[i+1][j+1] == "M")) {
				return 0
			}

			if !((word_search[i-1][j+1] == "M" && word_search[i+1][j-1] == "S") || (word_search[i-1][j+1] == "S" && word_search[i+1][j-1] == "M")) {
				return 0
			}
			return 1
		}
	}
	return 0
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(file), "\n")
	word_search := make([][]string, len(rows))
	key_word := "XMAS"
	for i, row := range rows {
		word_search[i] = strings.Split(row, "")
	}

	n_matches_p1 := 0
	n_matches_p2 := 0
	for i := range word_search {
		for j := range word_search[0] {
			n_matches_p1 += search_horizontal(word_search, key_word, i, j)
			n_matches_p1 += search_horizontal(word_search, reverse(key_word), i, j)
			n_matches_p1 += search_vertical(word_search, key_word, i, j)
			n_matches_p1 += search_vertical(word_search, reverse(key_word), i, j)
			n_matches_p1 += search_neg_diagonal(word_search, key_word, i, j)
			n_matches_p1 += search_neg_diagonal(word_search, reverse(key_word), i, j)
			n_matches_p1 += search_pos_diagonal(word_search, key_word, i, j)
			n_matches_p1 += search_pos_diagonal(word_search, reverse(key_word), i, j)
			n_matches_p2 += search_xmas(word_search, i, j)
		}
	}
	fmt.Printf("Part 1: %d\n", n_matches_p1)
	fmt.Printf("Part 2: %d\n", n_matches_p2)
}
