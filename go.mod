module github.com/ghonzo/advent2025

go 1.25

// For generics constraints
require golang.org/x/exp v0.0.0-20251125195548-87e1e737ad39

// Easier JSON parsing for leaderboard.go
require (
	github.com/tidwall/gjson v1.18.0
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
)

// Set
require github.com/deckarep/golang-set/v2 v2.8.0

// Priorirty queue (Used in day 16)
//require github.com/oleiade/lane/v2 v2.0.0
