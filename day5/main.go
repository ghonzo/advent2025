// Advent of Code 2025, Day 5
package main

import (
	"fmt"

	"github.com/ghonzo/advent2025/common"
)

// Day 5: Cafeteria
// Part 1 answer: 638
// Part 2 answer: 352946349407338
func main() {
	fmt.Println("Advent of Code 2025, Day 5")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var intervals []common.Interval
	// Read the ranges
	var lineNum int
	for lineNum = 0; entries[lineNum] != ""; lineNum++ {
		pair := common.ConvertToInts(entries[lineNum])
		intervals = append(intervals, common.Interval{Start: pair[0], End: pair[1]})
	}
	// Now read the IDs to check
	var fresh int
	for _, entry := range entries[lineNum+1:] {
		id := common.Atoi(entry)
		for _, t := range intervals {
			if t.Includes(id) {
				fresh++
				break
			}
		}
	}
	return fresh
}

func part2(entries []string) int {
	var intervals []common.Interval
	// Read the ranges
	for _, entry := range entries {
		if entry == "" {
			break
		}
		pair := common.ConvertToInts(entry)
		intervals = append(intervals, common.Interval{Start: pair[0], End: pair[1]})
	}
	// Now find all the overlapping ranges and merge them
	merged := common.MergeOverlappingIntervals(intervals)
	var sum int
	for _, t := range merged {
		sum += t.End - t.Start + 1
	}
	return sum
}
