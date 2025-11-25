<!-- Copilot / AI Agent instructions for the `advent2025` repo -->

# Quick Project Summary
- This repository contains Advent of Code 2025 solutions in Go. Each day lives in a top-level `dayN/` folder (for example `day1/`, `day2/`, ...).
- Shared utilities are in the `common` package. Example helpers: `common.ReadStringsFromFile(...)` used by many tests.
- The module path is `github.com/ghonzo/advent2025` (see `go.mod`).

# What to edit
- Day folders: each `dayN/` contains `main.go`, `main_test.go` and a `testdata/` directory with example inputs. `main.go` typically exposes `part1(entries []string) int` and `part2(entries []string) int` (or similar signatures) which tests call.
- `common/`: utilities shared across days (file parsing, grid/point helpers, min/max, etc.).
- `leaderboard.go` and root-level files may aggregate results across days.

# Test & build workflow (how to run things)
- Run all tests: `go test ./...` (from repo root). Tests are organized per-package; they rely on relative paths like `testdata/example.txt`.
- Run a single day's tests: `go test ./day1 -run Test_part1` or `go test ./day1 -run Test_part2`.
- Build: `go build ./...` or `go build ./day1` for a single package.

# Repository conventions and important patterns
- Input files: the real puzzle input is stored as `dayN/input.txt`. Example inputs for tests are under `dayN/testdata/*.txt`.
- Tests call `common.ReadStringsFromFile("testdata/example.txt")` (relative to package dir). When running `go test ./dayN`, the package’s working dir is the package directory and this path resolves correctly.
- Tests are named `Test_part1`/`Test_part2` and use table-driven tests. Use the same naming to keep test discovery and selective runs easy.
- Keep functions small and deterministic. Each day's `part1`/`part2` should be pure functions taking inputs (usually `[]string`) and returning simple types (int, string, etc.) so they are easy to unit test.

# Patterns & examples (copyable)
- Read example file in tests: `entries := common.ReadStringsFromFile("testdata/example.txt")`
- Test assertion pattern:
  - `if got := part1(entries); got != tt.want { t.Errorf("part1() = %v, want %v", got, tt.want) }`

# What AI agents should avoid changing
- Do not change the module path in `go.mod` (unless explicitly instructed) — many imports use `github.com/ghonzo/advent2025`.
- Do not rename `testdata` directories or move example files — tests expect those relative paths.

# Typical small tasks an agent can perform
- Add or fix a day's solution: add/modify `main.go` and update `main_test.go` to include example assertions.
- Add utility functions to `common/` and update call sites across day packages.
- Improve test coverage by adding new `testdata/example*.txt` files and table rows in `main_test.go`.

# When in doubt
- Run `go test ./...` locally to verify changes.
- Inspect `common/` and a day's `main.go`/`main_test.go` to learn the specific shape expected for that day.

# Ask the maintainer
- If a change requires altering the module path, testdata layout, or package boundaries, confirm with the repo owner before changing — those are global-breaking modifications.

---
If you'd like, I can: run the full test suite, add a missing `part1`/`part2` implementation for a given day, or merge in any existing agent guidance you have. What should I do next?
