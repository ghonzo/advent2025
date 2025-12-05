package common

import "sort"

// Interval represents a closed interval [Start, End]
// TODO: Can make this generic with [T cmp.Ordered]
type Interval struct {
	Start, End int
}

// Includes returns true if the value is within (closed) interval
func (t Interval) Includes(v int) bool {
	return v >= t.Start && v <= t.End
}

// Merge any overlapping or adjacent intervals and return a slice
// containing the union of those intervals.
func MergeOverlappingIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}
	// Sort intervals by their Start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}

	for _, t := range intervals[1:] {
		lastMerged := &merged[len(merged)-1]
		if t.Start <= lastMerged.End { // Overlap or adjacent
			if t.End > lastMerged.End {
				lastMerged.End = t.End
			}
		} else { // No overlap
			merged = append(merged, t)
		}
	}
	return merged
}
