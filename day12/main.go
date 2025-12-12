// Advent of Code 2025, Day 12
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ghonzo/advent2025/common"
)

// Day 12: Christmas Tree Farm
// Part 1 answer: 510
func main() {
	fmt.Println("Advent of Code 2025, Day 12")
	start := time.Now()
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d (%s)\n", part1(entries), time.Since(start))
}

func part1(entries []string) int {
	var count int
	areas := []int{6, 7, 7, 7, 5, 7}
	for _, entry := range entries {
		if len(entry) > 2 && entry[2] == 'x' {
			spaceArea := common.Atoi(entry[:2]) * common.Atoi(entry[3:5])
			var shapeArea int
			for i, numShapesStr := range strings.Fields(entry[7:]) {
				shapeArea += areas[i] * common.Atoi(numShapesStr)
			}
			if shapeArea <= spaceArea {
				count++
			}
		}
	}
	return count
}
