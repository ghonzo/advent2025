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
	// First read the shapes. We just need to record the areas
	var areas [6]int
	for i := range 6 {
		var area int
		for lineNum := 5*i + 1; lineNum <= 5*i+3; lineNum++ {
			area += strings.Count(entries[lineNum], "#")
		}
		areas[i] = area
	}
	for _, entry := range entries[30:] {
		colonIndex := strings.IndexByte(entry, ':')
		spaceDims := common.ConvertToInts(entry[:colonIndex])
		spaceArea := spaceDims[0] * spaceDims[1]
		var shapeArea int
		for i, numShapesStr := range strings.Fields(entry[colonIndex+2:]) {
			shapeArea += areas[i] * common.Atoi(numShapesStr)
		}
		if shapeArea <= spaceArea {
			count++
		}
	}
	return count
}
