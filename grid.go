package gridgrouper

import (
	"math/bits"
	"strconv"
	"strings"
)

// BitGrid stores a square grid of dots in a compact manner.
type BitGrid []uint64

// BoolGrid unpacks a bit grid into a bool grid.
func (grid BitGrid) BoolGrid() BoolGrid {
	cols := grid.Cols()
	if cols == 0 {
		return nil
	}
	boolGrid := make(BoolGrid, len(grid))
	for row := range grid {
		boolRow := make([]bool, cols)
		for col := 0; col < cols; col++ {
			boolRow[col] = grid[row]&(1<<uint(col)) != 0
		}
		boolGrid[row] = boolRow
	}
	return boolGrid
}

// PosSet unpacks a bit grid into a position set.
func (grid BitGrid) PosSet() PosSet {
	var positions []Pos
	for row := range grid {
		cols := bits.Len64(grid[row])
		for col := 0; col < cols; col++ {
			if grid[row]&(1<<uint(col)) != 0 {
				positions = append(positions, Pos{row, col})
			}
		}
	}
	return positions
}

// Cols returns the width of the widest row.
func (grid BitGrid) Cols() int {
	maxWidth := 0
	for row := range grid {
		if width := bits.Len64(grid[row]); width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}

func (grid BitGrid) String() string {
	cols := grid.Cols()
	var b strings.Builder
	for row := range grid {
		for col := 0; col < cols; col++ {
			if grid[row]&(1<<uint(col)) != 0 {
				b.WriteByte('X')
			} else {
				b.WriteByte(' ')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// BoolGrid stores a square grid of dots.
type BoolGrid [][]bool

// NewBoolGrid constructs a bool grid of a given size.
func NewBoolGrid(rows, cols int) BoolGrid {
	grid := make(BoolGrid, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}
	return grid
}

// BitGrid packs a bool grid into a bit grid.
func (grid BoolGrid) BitGrid() BitGrid {
	bitGrid := make(BitGrid, len(grid))
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] {
				bitGrid[row] |= 1 << uint(col)
			}
		}
	}
	return bitGrid
}

// PosSet unpacks a bool grid into a position set.
func (grid BoolGrid) PosSet() PosSet {
	var positions []Pos
	for row := range grid {
		for col := range grid {
			if grid[row][col] {
				positions = append(positions, Pos{row, col})
			}
		}
	}
	return positions
}

func (grid BoolGrid) String() string {
	var b strings.Builder
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] {
				b.WriteByte('X')
			} else {
				b.WriteByte(' ')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// Pos contains a row and column position in a grid.
type Pos [2]int

func (pos Pos) String() string {
	var b strings.Builder
	b.WriteByte('(')
	b.WriteString(strconv.FormatInt(int64(pos[0]), 10))
	b.WriteByte(' ')
	b.WriteString(strconv.FormatInt(int64(pos[1]), 10))
	b.WriteByte(')')
	return b.String()
}

// PosSet is a position set with elements in no particular order.
type PosSet []Pos

// BitGrid packs a position set into a bit grid.
func (positions PosSet) BitGrid() BitGrid {
	grid := make(BitGrid, positions.Rows())
	for _, pos := range positions {
		grid[pos[0]] |= 1 << uint(pos[1])
	}
	return grid
}

// BoolGrid packs a position set into a bool grid.
func (positions PosSet) BoolGrid() BoolGrid {
	grid := NewBoolGrid(positions.Rows(), positions.Cols())
	for _, pos := range positions {
		grid[pos[0]][pos[1]] = true
	}
	return grid
}

// Rows returns the minimum number of rows to hold the position set.
func (positions PosSet) Rows() int {
	maxRows := 0
	for _, pos := range positions {
		if pos[0] >= maxRows {
			maxRows = pos[0] + 1
		}
	}
	return maxRows
}

// Cols returns the minimum number of columns to hold the position set.
func (positions PosSet) Cols() int {
	maxCols := 0
	for _, pos := range positions {
		if pos[1] >= maxCols {
			maxCols = pos[1] + 1
		}
	}
	return maxCols
}

// Equal compares two position sets with positions in any order.
func (positions PosSet) Equal(compare PosSet) bool {
	if len(positions) != len(compare) {
		return false
	}
	lookup := make(map[Pos]bool)
	for _, pos := range compare {
		lookup[pos] = true
	}
	if len(lookup) != len(compare) {
		return false
	}
	for _, pos := range positions {
		if !lookup[pos] {
			return false
		}
		lookup[pos] = false
	}
	return true
}
