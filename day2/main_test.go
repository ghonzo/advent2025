// Advent of Code 2025, Day 2
package main

import (
	"testing"

	"github.com/ghonzo/advent2025/common"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line string
		want int
	}{
		{"example", common.ReadStringsFromFile("testdata/example.txt")[0], 1227775554},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.line)
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line string
		want int
	}{
		{"example", common.ReadStringsFromFile("testdata/example.txt")[0], 4174379265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.line)
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
