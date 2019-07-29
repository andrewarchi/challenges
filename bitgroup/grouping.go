package bitgroup

import (
	"math/bits"
)

// Grid stores a square grid of dots in a compact manner.
type Grid []uint64

// Groups computes all groups of connected dots.
func (grid Grid) Groups() []Grid {
	var groups []Grid
	for row := range grid {
		for grid[row] != 0 {
			group := make(Grid, len(grid))
			mask := getConsecutiveOnes(grid[row])
			fillGroup(grid, group, row, mask)
			groups = append(groups, group)
		}
	}
	return groups
}

// fillGroup recursively fills a group with connected dots starting a given
// mask for the current row.
// Preconditions: row >= 0, row < len(grid), grid[row] != 0, mask != 0.
func fillGroup(grid, group Grid, row int, mask uint64) {
	if mask == 0 {
		return
	}
	group[row] |= mask
	grid[row] &^= mask

	mask |= mask<<1 | mask>>1
	if prev := row - 1; prev >= 0 && grid[prev] != 0 {
		fillGroup(grid, group, prev, grid[prev]&mask)
	}
	if next := row + 1; next < len(grid) && grid[next] != 0 {
		fillGroup(grid, group, next, grid[next]&mask)
	}
}

// getConsecutiveOnes returns a mask of the most significant series of
// consecutive ones.
func getConsecutiveOnes(x uint64) uint64 {
	leftLen := uint64(bits.Len64(x))
	left := uint64(1<<leftLen - 1)
	rightLen := uint64(bits.Len64(left &^ x))
	right := uint64(1<<rightLen - 1)
	return left &^ right
}
