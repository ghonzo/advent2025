// Advent of Code 2025, Day 3
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2025/common"
)

// Day 3: Lobby
// Part 1 answer: 17207
// Part 2 answer: 170997883706617
func main() {
	fmt.Println("Advent of Code 2025, Day 3")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var sum int
	for _, line := range entries {
		tens, index := findBiggest(line[:len(line)-1])
		ones, _ := findBiggest(line[index+1:])
		sum += tens*10 + ones
	}
	return sum
}

// Find the largest digit in the string and also return its index
func findBiggest(line string) (int, int) {
	for n := 9; n >= 1; n-- {
		sn := strconv.Itoa(n)
		index := strings.Index(line, sn)
		if index != -1 {
			return n, index
		}
	}
	panic("no digit found")
}

func part2(entries []string) int {
	var sum int
	for _, line := range entries {
		lenLine := len(line)
		// i is the current digit we're looking for
		var joltage int
		var lastIndex int
		for i := 1; i <= 12; i++ {
			digit, index := findBiggest(line[lastIndex : lenLine+i-12])
			joltage = joltage*10 + digit
			lastIndex += index + 1
		}
		sum += joltage
	}
	return sum
}
