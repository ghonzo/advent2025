// Advent of Code 2025, Day 9
package main

import (
	"fmt"
	"time"

	"github.com/ghonzo/advent2025/common"
)

// Day 9: Movie Theater
// Part 1 answer: 4759930955
// Part 2 answer: 1525241870
func main() {
	fmt.Println("Advent of Code 2025, Day 9")
	start := time.Now()
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d (%s)\n", part1(entries), time.Since(start))
	fmt.Printf("Part 2: %d (%s)\n", part2(entries), time.Since(start))
}

func part1(entries []string) int {
	points := readPoints(entries)
	var maxArea int
	for i := 0; i < len(points)-1; i++ {
		p0 := points[i]
		for j := i + 1; j < len(points); j++ {
			p1 := points[j]
			maxArea = max(maxArea, rect{p0, p1}.area())
		}
	}
	return maxArea
}

func part2(entries []string) int {
	points := readPoints(entries)
	var maxArea int
	for i := 0; i < len(points)-1; i++ {
		p0 := points[i]
		for j := i + 1; j < len(points); j++ {
			p1 := points[j]
			r := rect{p0, p1}
			area := r.area()
			if area <= maxArea {
				// don't even bother
				continue
			}
			// Now make sure the bounding rectangle is completely within the polygon
			if r.allInside(points) {
				maxArea = area
			}
		}
	}
	return maxArea
}

func readPoints(entries []string) []common.Point {
	var points []common.Point
	for _, entry := range entries {
		c := common.ConvertToInts(entry)
		points = append(points, common.NewPoint(c[0], c[1]))
	}
	return points
}

type line struct {
	pA, pB common.Point
}

func (l line) isHorizontal() bool {
	return l.pA.Y() == l.pB.Y()
}

func (l line) containsPoint(p common.Point) bool {
	return between(p.X(), l.pA.X(), l.pB.X()) && between(p.Y(), l.pA.Y(), l.pB.Y())
}

func between(x, a, b int) bool {
	return x >= min(a, b) && x <= max(a, b)
}

// Rectangle defined by two red tiles
type rect struct {
	a, b common.Point
}

func (r rect) area() int {
	return (common.Abs(r.a.X()-r.b.X()) + 1) * (common.Abs(r.a.Y()-r.b.Y()) + 1)
}

// Return all the lines of the bounding box
func (r rect) boundaryLines() []line {
	if r.a.X() == r.b.X() || r.a.Y() == r.b.Y() {
		// Degenerate case of just one line to check
		return []line{{r.a, r.b}}
	}
	// sorry for the strange numbering
	p0 := common.NewPoint(r.b.X(), r.a.Y())
	p1 := common.NewPoint(r.a.X(), r.b.Y())
	return []line{{r.a, p0}, {p0, r.b}, {r.b, p1}, {p1, r.a}}

}

func (r rect) allInside(polygon []common.Point) bool {
	// We know that r.a and r.b are "in". Let's make sure the other corners are (unless it's degenerate)
	if r.a.X() != r.b.X() && r.a.Y() != r.b.Y() {
		if !pointInPolygon(common.NewPoint(r.a.X(), r.b.Y()), polygon) ||
			!pointInPolygon(common.NewPoint(r.b.X(), r.a.Y()), polygon) {
			return false
		}
	}
	// Now make sure no segments of the polygon cross the bounding rectangle
	boundaryLines := r.boundaryLines()
	// Loop through all edges of the polygon
	for i := range polygon {
		edge := line{polygon[i], polygon[(i+1)%len(polygon)]} // Connect last vertex to first
		for _, l := range boundaryLines {
			if linesCross(l, edge) {
				return false
			}
		}
	}
	return true
}

func linesCross(a, b line) bool {
	if a.isHorizontal() {
		if b.isHorizontal() {
			return false
		}
		// a is horiztonal, b is vertical
		return crosses(b.pA.Y(), b.pB.Y(), a.pA.Y()) && crosses(a.pA.X(), a.pB.X(), b.pA.X())
	} else {
		if !b.isHorizontal() {
			return false
		}
		// a is vertical, b is horizontal
		return crosses(b.pA.X(), b.pB.X(), a.pA.X()) && crosses(a.pA.Y(), a.pB.Y(), b.pA.Y())
	}
}

func crosses(a, b, c int) bool {
	// Crosses if opposite and not zero
	return common.Sgn(a-c)*common.Sgn(b-c) == -1
}

func pointInPolygon(p common.Point, polygon []common.Point) bool {
	var inside bool
	for i := range polygon {
		edge := line{polygon[i], polygon[(i+1)%len(polygon)]} // Connect last vertex to first
		// if it's actually on one of the edges, it's considered in
		if edge.containsPoint(p) {
			return true
		}
		// ray cast
		if ((edge.pB.Y() < p.Y() && edge.pA.Y() >= p.Y()) || (edge.pA.Y() < p.Y() && edge.pB.Y() >= p.Y())) &&
			(edge.pB.X() <= p.X() || edge.pA.X() <= p.X()) {
			if edge.pB.X()+(p.Y()-edge.pB.Y())/(edge.pA.Y()-edge.pB.Y())*(edge.pA.X()-edge.pB.X()) < p.X() {
				inside = !inside
			}
		}
	}
	return inside
}
