// Advent of Code 2025, Day 1
package main

import (
	"fmt"

	"github.com/ghonzo/advent2025/common"
)

// Day 1: Secret Entrance
// Part 1 answer: 1066
// Part 2 answer: 6223
func main() {
	fmt.Println("Advent of Code 2025, Day 1")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	pos := 50
	var count int
	for _, line := range entries {
		delta := common.Atoi(line[1:])
		if line[0] == 'L' {
			delta *= -1
		}
		pos += delta
		pos = common.Mod(pos, 100)
		if pos == 0 {
			count++
		}
	}
	return count
}

func part2(entries []string) int {
	pos := 50
	var count int
	for _, line := range entries {
		delta := common.Atoi(line[1:])
		if line[0] == 'L' {
			delta *= -1
		}
		startAtZero := pos == 0
		pos += delta
		for pos < 0 {
			pos += 100
			if startAtZero {
				// I know this is strange, maybe we need a better variable name
				startAtZero = false
				continue
			}
			if pos != 0 || delta < -100 {
				count++
			}
		}
		for pos >= 100 {
			pos -= 100
			if pos != 0 {
				count++
			}
		}
		if pos == 0 {
			count++
		}
	}
	return count
}
