// Advent of Code 2025, Day 8
package main

import (
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2025/common"
)

// Day 8: Playground
// Part 1 answer: 67488
// Part 2 answer: 3767453340
func main() {
	fmt.Println("Advent of Code 2025, Day 8")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, 1000))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type box struct {
	x, y, z   int
	connected mapset.Set[*box]
}

type pair struct {
	a, b *box
	dist int
}

// This is actually the square of the distance
func distance(a, b *box) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func part1(entries []string, connect int) int {
	var allBoxes []*box
	for _, entry := range entries {
		coords := common.ConvertToInts(entry)
		b := &box{
			x:         coords[0],
			y:         coords[1],
			z:         coords[2],
			connected: mapset.NewThreadUnsafeSet[*box](),
		}
		allBoxes = append(allBoxes, b)
	}
	var allPairs []pair
	for i := 0; i < len(allBoxes)-1; i++ {
		for j := i + 1; j < len(allBoxes); j++ {
			allPairs = append(allPairs, pair{
				a:    allBoxes[i],
				b:    allBoxes[j],
				dist: distance(allBoxes[i], allBoxes[j]),
			})
		}
	}
	// Sort pairs by distance
	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i].dist < allPairs[j].dist
	})
	// Now connect boxes
	for _, p := range allPairs[:connect] {
		p.a.connected.Add(p.b)
		p.b.connected.Add(p.a)
	}
	visited := mapset.NewThreadUnsafeSet[*box]()
	var circuitSizes []int
	// Now find all the circuit sizes
	for _, b := range allBoxes {
		if !visited.Contains(b) {
			circuit := findCircuit(b)
			circuitSizes = append(circuitSizes, circuit.Cardinality())
			visited = visited.Union(circuit)
		}
	}
	// Now find the three largest circuit sizes and multiply them
	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})
	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func findCircuit(start *box) mapset.Set[*box] {
	toVisit := mapset.NewThreadUnsafeSet(start)
	circuit := mapset.NewThreadUnsafeSet[*box]()
	for {
		b, ok := toVisit.Pop()
		if !ok {
			return circuit
		}
		if !circuit.Contains(b) {
			circuit.Add(b)
			toVisit = toVisit.Union(b.connected)
		}
	}
}

func part2(entries []string) int {
	var allBoxes []*box
	for _, entry := range entries {
		coords := common.ConvertToInts(entry)
		b := &box{
			x:         coords[0],
			y:         coords[1],
			z:         coords[2],
			connected: mapset.NewThreadUnsafeSet[*box](),
		}
		allBoxes = append(allBoxes, b)
	}
	var allPairs []pair
	for i := 0; i < len(allBoxes)-1; i++ {
		for j := i + 1; j < len(allBoxes); j++ {
			allPairs = append(allPairs, pair{
				a:    allBoxes[i],
				b:    allBoxes[j],
				dist: distance(allBoxes[i], allBoxes[j]),
			})
		}
	}
	// Sort pairs by distance
	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i].dist < allPairs[j].dist
	})
	// Now connect boxes
	allConnected := false
outer:
	for _, p := range allPairs {
		p.a.connected.Add(p.b)
		p.b.connected.Add(p.a)
		// Any unconnected boxes?
		if !allConnected {
			for _, b := range allBoxes {
				if b.connected.IsEmpty() {
					continue outer
				}
			}
			allConnected = true
		}
		if findCircuit(allBoxes[0]).Cardinality() == len(allBoxes) {
			return p.a.x * p.b.x
		}
	}
	panic("No solution found")
}
