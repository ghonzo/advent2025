// Advent of Code 2025, Day 11
package main

import (
	"fmt"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2025/common"
)

// Day 11: Reactor
// Part 1 answer: 571
// Part 2 answer: 511378159390560
func main() {
	fmt.Println("Advent of Code 2025, Day 11")
	start := time.Now()
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d (%s)\n", part1(entries), time.Since(start))
	fmt.Printf("Part 2: %d (%s)\n", part2(entries), time.Since(start))
}

type node struct {
	name string
	next []*node
}

func part1(entries []string) int {
	nodeMap := readNodeMap(entries)
	return countPaths("you", "out", nodeMap)
}

func readNodeMap(entries []string) map[string]*node {
	nodeMap := make(map[string]*node)
	for _, entry := range entries {
		name := entry[:3]
		source := getNode(name, nodeMap)
		for destName := range strings.FieldsSeq(entry[5:]) {
			dest := getNode(destName, nodeMap)
			source.next = append(source.next, dest)
		}
	}
	return nodeMap
}

// Retrieve or create the node and store it in the given map
func getNode(name string, nodeMap map[string]*node) *node {
	n, ok := nodeMap[name]
	if !ok {
		n = &node{name: name}
		nodeMap[name] = n
	}
	return n
}

// This is brute-force. A better way would be to backtrack and accumulate pathcounts
func countPaths(src, dst string, nodeMap map[string]*node) int {
	var count int
	// These are all the nodes that are downstream of dst. Useless for part1, helpful for part2
	killNodes := findKillNodes(dst, nodeMap)
	nextNodes := []*node{nodeMap[src]}
	for len(nextNodes) > 0 {
		n := nextNodes[len(nextNodes)-1]
		nextNodes = nextNodes[:len(nextNodes)-1]
		if n.name == dst {
			count++
		} else if !killNodes.Contains(n) {
			nextNodes = append(nextNodes, n.next...)
		}
	}
	return count
}

// Return the set of all nodes that are "downstream" of the given node
func findKillNodes(name string, nodeMap map[string]*node) mapset.Set[*node] {
	toVisit := mapset.NewThreadUnsafeSet(nodeMap[name].next...)
	visited := mapset.NewThreadUnsafeSet[*node]()
	for !toVisit.IsEmpty() {
		n, _ := toVisit.Pop()
		if !visited.Contains(n) {
			visited.Add(n)
			toVisit.Append(n.next...)
		}
	}
	return visited
}

func part2(entries []string) int {
	nodeMap := readNodeMap(entries)
	// This is one of those rare Advent Of Code cases where the solution doesn't take
	// either 0.1 seconds or 10,000 years. This brute force method takes about four minutes
	// with my data set and computer. I know we can be more efficient by backtracing and
	// accumulating paths combinatorially, but I'm just going to take the win and move on.
	part1 := countPaths("svr", "fft", nodeMap)
	part2 := countPaths("fft", "dac", nodeMap)
	part3 := countPaths("dac", "out", nodeMap)
	return part1 * part2 * part3
}
