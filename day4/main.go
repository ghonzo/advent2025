// Advent of Code 2025, Day 4
package main

import (
	"fmt"

	"github.com/ghonzo/advent2025/common"
)

// Day 4: Printing Department
// Part 1 answer: 1549
// Part 2 answer: 8887
func main() {
	fmt.Println("Advent of Code 2025, Day 4")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var count int
	grid := common.ArraysGridFromLines(entries)
	for p := range grid.AllPoints() {
		if grid.Get(p) == '@' {
			var adjacent int
			for s := range p.SurroundingPoints() {
				if v, _ := grid.CheckedGet(s); v == '@' {
					adjacent++
				}
			}
			if adjacent < 4 {
				count++
			}
		}
	}
	return count
}

func part2(entries []string) int {
	var count int
	grid := common.ArraysGridFromLines(entries)
	for {
		var pointsToRemove []common.Point
		for p := range grid.AllPoints() {
			if grid.Get(p) == '@' {
				var adjacent int
				for s := range p.SurroundingPoints() {
					if v, _ := grid.CheckedGet(s); v == '@' {
						adjacent++
					}
				}
				if adjacent < 4 {
					pointsToRemove = append(pointsToRemove, p)
				}
			}
		}
		if len(pointsToRemove) == 0 {
			return count
		}
		for _, p := range pointsToRemove {
			grid.Set(p, '.')
			count++
		}
	}
}
