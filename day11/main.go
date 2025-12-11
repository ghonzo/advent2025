// Advent of Code 2025, Day 11
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2025/common"
)

// Day 11: Reactor
// Part 1 answer: 571
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2025, Day 11")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type node struct {
	name string
	next []*node
}

func part1(entries []string) int {
	nodeMap := make(map[string]*node)
	for _, entry := range entries {
		name := entry[:3]
		source := getNode(name, nodeMap)
		for destName := range strings.FieldsSeq(entry[5:]) {
			dest := getNode(destName, nodeMap)
			source.next = append(source.next, dest)
		}
	}
	var count int
	nextNodes := []*node{nodeMap["you"]}
	for len(nextNodes) > 0 {
		n := nextNodes[len(nextNodes)-1]
		nextNodes = nextNodes[:len(nextNodes)-1]
		if n.name == "out" {
			count++
		} else {
			nextNodes = append(nextNodes, n.next...)
		}
	}
	return count
}

func getNode(name string, nodeMap map[string]*node) *node {
	n, ok := nodeMap[name]
	if !ok {
		n = &node{name: name}
		nodeMap[name] = n
	}
	return n
}

type path struct {
	cur      *node
	dac, fft bool
}

func (p path) nextPaths() []path {
	paths := make([]path, len(p.cur.next))
	for i, n := range p.cur.next {
		paths[i] = path{cur: n, dac: p.dac || p.cur.name == "dac", fft: p.fft || p.cur.name == "fft"}
	}
	return paths
}

func part2(entries []string) int {
	nodeMap := make(map[string]*node)
	for _, entry := range entries {
		name := entry[:3]
		source := getNode(name, nodeMap)
		for destName := range strings.FieldsSeq(entry[5:]) {
			dest := getNode(destName, nodeMap)
			source.next = append(source.next, dest)
		}
	}
	var count int
	paths := []path{{cur: nodeMap["svr"]}}
	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		if p.cur.name == "out" {
			if p.dac && p.fft {
				count++
			}
		} else {
			paths = append(paths, p.nextPaths()...)
		}
	}
	return count
}
