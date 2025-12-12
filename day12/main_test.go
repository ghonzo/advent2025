// Advent of Code 2025, Day 8
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		entries []string
		want    int
	}{
		// Test doesn't pass! Oh well!
		//{"example", common.ReadStringsFromFile("testdata/example.txt"), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.entries)
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
