// Advent of Code 2025, Day 5
package main

import (
	"fmt"
	"sort"

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

type Range struct {
	start, end int
}

func (r Range) Includes(v int) bool {
	return v >= r.start && v <= r.end
}

// containing the union of those ranges, merging any overlapping or
// adjacent ranges.
func mergeOverlappingRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return nil
	}

	// Sort ranges by their start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		lastMerged := &merged[len(merged)-1]
		current := ranges[i]

		if current.start <= lastMerged.end { // Overlap or adjacent
			if current.end > lastMerged.end {
				lastMerged.end = current.end
			}
		} else { // No overlap
			merged = append(merged, current)
		}
	}
	return merged
}

func part1(entries []string) int {
	var ranges []Range
	var fresh int
	var ids bool
	for _, entry := range entries {
		if entry == "" {
			ids = true
			continue
		}
		if ids {
			for _, r := range ranges {
				if r.Includes(common.Atoi(entry)) {
					fresh++
					break
				}
			}
		} else {
			var r Range
			pair := common.ConvertToInts(entry)
			r.start = pair[0]
			r.end = pair[1]
			ranges = append(ranges, r)
		}
	}
	return fresh
}

func part2(entries []string) int {
	var ranges []Range
	for _, entry := range entries {
		if entry == "" {
			break
		}
		var r Range
		pair := common.ConvertToInts(entry)
		r.start = pair[0]
		r.end = pair[1]
		ranges = append(ranges, r)
	}
	// Now find all the overlapping ranges and merge them
	merged := mergeOverlappingRanges(ranges)
	var sum int
	for _, r := range merged {
		sum += r.end - r.start + 1
	}
	return sum
}
