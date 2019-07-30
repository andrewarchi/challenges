package gridgrouper

import (
	"math/bits"
)

// Groups computes all groups of connected dots.
func (grid BitGrid) Groups() []BitGrid {
	var groups []BitGrid
	for row := range grid {
		for grid[row] != 0 {
			group := make(BitGrid, len(grid))
			mask := ConsecutiveOnes64(grid[row])
			fillGroup(grid, group, row, mask)
			groups = append(groups, group)
		}
	}
	return groups
}

// fillGroup recursively fills a group with connected dots starting with a given
// mask for the current row.
// Precondition: 0 <= row && row < len(grid).
func fillGroup(grid, group BitGrid, row int, mask uint64) {
	mask &= grid[row]
	if mask == 0 {
		return
	}
	mask = expandMaskAdjacent(grid[row], mask)

	group[row] |= mask
	grid[row] &^= mask

	mask |= mask<<1 | mask>>1
	if prev := row - 1; prev >= 0 && grid[prev] != 0 {
		fillGroup(grid, group, prev, mask)
	}
	if next := row + 1; next < len(grid) && grid[next] != 0 {
		fillGroup(grid, group, next, mask)
	}
}

// ConsecutiveOnes64 returns a mask of the most significant series of
// consecutive ones.
func ConsecutiveOnes64(x uint64) uint64 {
	leftLen := uint64(bits.Len64(x))
	left := uint64(1<<leftLen - 1)
	rightLen := uint64(bits.Len64(left &^ x))
	right := uint64(1<<rightLen - 1)
	return left &^ right
}

// expandMaskAdjacent expands a mask to include any sequences of bits set in x
// that are adjacent to the mask.
func expandMaskAdjacent(x, mask uint64) uint64 {
	for {
		adjacent := (mask | mask<<1 | mask>>1) & x
		if mask == adjacent {
			return mask
		}
		mask = adjacent
	}
}
