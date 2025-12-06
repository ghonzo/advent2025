// Advent of Code 2025, Day 6
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2025/common"
)

// Day 6: Trash Compactor
// Part 1 answer: 5335495999141
// Part 2 answer: 10142723156431
func main() {
	fmt.Println("Advent of Code 2025, Day 6")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var numbers [][]int
	var lineNum int
	for lineNum = 0; entries[lineNum][0] != '*'; lineNum++ {
		numbers = append(numbers, common.ConvertToInts(entries[lineNum]))
	}
	var sep = regexp.MustCompile(`\s+`)
	var sum int
	for i, op := range sep.Split(entries[lineNum], -1) {
		if op == "*" {
			subTotal := 1
			for _, n := range numbers {
				subTotal *= n[i]
			}
			sum += subTotal
		} else if op == "+" {
			for _, n := range numbers {
				sum += n[i]
			}
		}
	}
	return sum
}

func part2(entries []string) int {
	numLines := len(entries)
	var op byte
	var sum int
	var subTotal int
	for col := 0; ; col++ {
		if col < len(entries[numLines-1]) {
			if entries[numLines-1][col] == '*' {
				subTotal = 1
				op = '*'
			} else if entries[numLines-1][col] == '+' {
				subTotal = 0
				op = '+'
			}
		}
		num := readCol(entries, col)
		if num == 0 {
			sum += subTotal
			if col >= len(entries[numLines-1]) {
				break
			}
		} else {
			if op == '*' {
				subTotal *= num
			} else if op == '+' {
				subTotal += num
			}
		}
	}
	return sum
}

func readCol(entries []string, col int) int {
	var num int
	for row := 0; row < len(entries)-1; row++ {
		if col < len(entries[row]) {
			c := entries[row][col]
			if c >= '0' && c <= '9' {
				num = num*10 + int(c-'0')
			}
		}
	}
	return num
}
