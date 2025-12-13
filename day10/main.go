// Advent of Code 2025, Day 10
package main

import (
	"fmt"
	"math"
	"math/bits"
	"slices"
	"strings"

	"github.com/ghonzo/advent2025/common"
)

// Day 10: Factory
// Part 1 answer: 475
// Part 2 answer: 18273
func main() {
	fmt.Println("Advent of Code 2025, Day 10")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type button []int

type machine struct {
	lights   []bool
	buttons  []button
	joltages []int
}

func toUint(lights []bool) uint {
	// convert []bool to an int
	var v uint
	for _, light := range slices.Backward(lights) {
		v <<= 1
		if light {
			v++
		}
	}
	return v
}

func part1(entries []string) int {
	var sum int
	for _, entry := range entries {
		m := convertToMachine(entry)
		sum += minButtonPresses(m)
	}
	return sum
}

func convertToMachine(s string) machine {
	var m machine
	for f := range strings.FieldsSeq(s) {
		switch f[0] {
		case '[':
			// lights
			var lights []bool
			for p := 1; p < len(f)-1; p++ {
				lights = append(lights, f[p] == '#')
			}
			m.lights = lights
		case '(':
			m.buttons = append(m.buttons, button(common.ConvertToInts(f)))
		case '{':
			m.joltages = common.ConvertToInts(f)
		}
	}
	return m
}

func minButtonPresses(m machine) int {
	// To get to the pattern, we just need to press each button 0 or 1 times.
	// Pressing a button twice just undoes itself.
	// Let's treat it like a binary number
	var minPresses = math.MaxInt
	endState := toUint(m.lights)
	// i will represent every combination of button presses
	for i := uint(0); i < (1 << len(m.buttons)); i++ {
		var curLights uint
		for j := range bits.Len(i) {
			if (i & (1 << j)) != 0 {
				for _, b := range m.buttons[j] {
					curLights ^= (1 << b)
				}
			}
		}
		if curLights == endState {
			minPresses = min(minPresses, bits.OnesCount(i))
		}
	}
	return minPresses
}

func part2(entries []string) int {
	var sum int
	for _, entry := range entries {
		m := convertToMachine(entry)
		sum += minButtonPressesPart2(m)
	}
	return sum
}

// Records all the combinations of buttons presses (represents by a uint)
// that result in a light pattern
var buttonsPressesForPattern map[uint][]uint

var joltageDiffForButtonPress map[uint][]int

var joltageToPressesCache map[string]int

func minButtonPressesPart2(m machine) int {
	// A lot more work here, but can take inspiration from part 1.
	// We need to record the result of all the button press combinations
	// and use that to recursively determine joltages.
	buttonsPressesForPattern = make(map[uint][]uint)
	joltageDiffForButtonPress = make(map[uint][]int)
	joltageToPressesCache = make(map[string]int)
	// Now populate those first two maps
	// i will represent every combination of button presses
	for i := uint(0); i < (1 << len(m.buttons)); i++ {
		var curLights uint
		var joltageDiff = make([]int, len(m.joltages))
		for j := range bits.Len(i) {
			if (i & (1 << j)) != 0 {
				for _, b := range m.buttons[j] {
					curLights ^= (1 << b)
					joltageDiff[b]++
				}
			}
		}
		buttonsPressesForPattern[curLights] = append(buttonsPressesForPattern[curLights], i)
		joltageDiffForButtonPress[i] = joltageDiff
	}
	// Now the fun starts!
	return minPressesForJoltages(m.joltages)
}

func minPressesForJoltages(joltages []int) int {
	// First check the cache
	cacheKey := fmt.Sprintf("%v", joltages)
	if v, ok := joltageToPressesCache[cacheKey]; ok {
		return v
	}
	allZeros := true
	var curLights uint
	// Check for all zeros or any negatives as we go
	for i, joltage := range joltages {
		if joltage < 0 {
			// cannot be a valid solution
			return 9999
		} else if joltage > 0 {
			allZeros = false
			curLights ^= ((uint(joltage) & 1) << i)
		}
	}
	if allZeros {
		return 0
	}
	// Some big number
	minPresses := 9999
	// Now find all the button combinations that will generate the curLights pattern
	if allPresses, ok := buttonsPressesForPattern[curLights]; ok {
		for _, buttonPresses := range allPresses {
			diff := joltageDiffForButtonPress[buttonPresses]
			newJoltages := make([]int, len(diff))
			for i, dj := range diff {
				newJoltages[i] = (joltages[i] - dj) / 2
			}
			numPresses := bits.OnesCount(buttonPresses)
			remainingPresses := 2 * minPressesForJoltages(newJoltages)
			minPresses = min(minPresses, numPresses+remainingPresses)
		}
	}
	// Don't forget to cache it
	joltageToPressesCache[cacheKey] = minPresses
	return minPresses
}
