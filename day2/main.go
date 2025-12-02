// Advent of Code 2025, Day 2
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/advent2025/common"
)

// Day 2: Gift Shop
// Part 1 answer: 38437576669
// Part 2 answer: 49046150754
func main() {
	fmt.Println("Advent of Code 2025, Day 2")
	line := common.ReadStringsFromFile("input.txt")[0]
	fmt.Printf("Part 1: %d\n", part1(line))
	fmt.Printf("Part 2: %d\n", part2(line))
}

// This all would have been trivial if go regexp supported backreferences...
// part1 would be (\d+)\1 and part2 would be (\d+)\1+

func part1(line string) int {
	var sum int
	ids := common.ConvertToInts(line)
	for i := 0; i < len(ids); i += 2 {
		lowId := ids[i]
		highId := ids[i+1]
		for id := lowId; id <= highId; id++ {
			idStr := strconv.Itoa(id)
			if len(idStr)%2 != 0 {
				continue
			}
			if idStr[:len(idStr)/2] == idStr[len(idStr)/2:] {
				sum += id
			}
		}
	}
	return sum
}

func part2(line string) int {
	var sum int
	ids := common.ConvertToInts(line)
	for i := 0; i < len(ids); i += 2 {
		lowId := ids[i]
		highId := ids[i+1]
		for id := lowId; id <= highId; id++ {
			idStr := strconv.Itoa(id)
			if isRepeating(idStr) {
				sum += id
			}
		}
	}
	return sum
}

func isRepeating(s string) bool {
	n := len(s)
outer:
	for lenMatch := 1; lenMatch <= n/2; lenMatch++ {
		if n%lenMatch != 0 {
			continue
		}
		firstPart := s[:lenMatch]
		for parts := 1; parts < n/lenMatch; parts++ {
			if s[parts*lenMatch:(parts+1)*lenMatch] != firstPart {
				continue outer
			}
		}
		return true
	}
	return false
}
