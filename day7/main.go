// Advent of Code 2025, Day 7
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2025/common"
)

// Day 7: Laboratories
// Part 1 answer: 1660
// Part 2 answer: 305999729392659
func main() {
	fmt.Println("Advent of Code 2025, Day 7")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var count int
	beam := make([]bool, len(entries[0]))
	start := strings.Index(entries[0], "S")
	beam[start] = true
	for i := 2; i < len(entries); i += 2 {
		line := entries[i]
		nextBeam := make([]bool, len(entries[0]))
		for x, b := range beam {
			if b {
				if line[x] == '^' {
					nextBeam[x-1] = true
					nextBeam[x+1] = true
					count++
				} else {
					nextBeam[x] = true
				}
			}
		}
		beam = nextBeam
	}
	return count
}

var memo [][]int

func part2(entries []string) int {
	memo = make([][]int, len(entries))
	for i := range memo {
		memo[i] = make([]int, len(entries[0]))
	}
	start := strings.Index(entries[0], "S")
	return timelines(entries, 2, start)
}

// Returns the number of timelines at this given row and position
func timelines(entries []string, row, pos int) int {
	if row >= len(entries) {
		return 1
	}
	if entries[row][pos] == '^' {
		if memo[row][pos] == 0 {
			memo[row][pos] = timelines(entries, row+2, pos-1) + timelines(entries, row+2, pos+1)
		}
		return memo[row][pos]
	}
	return timelines(entries, row+2, pos)
}
