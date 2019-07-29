package bitgroup

import (
	"strconv"
	"strings"
)

// Pos contains the row and column position in a grid.
type Pos [2]int

// PackBool packs a bool grid into a compact grid.
func PackBool(boolGrid [][]bool) Grid {
	grid := make(Grid, len(boolGrid))
	for row := range boolGrid {
		for col := range boolGrid[row] {
			if boolGrid[row][col] {
				grid[row] |= 1 << uint(col)
			}
		}
	}
	return grid
}

// PackPos packs positions into a compact grid.
func PackPos(positions []Pos) Grid {
	grid := make(Grid, len(positions))
	for _, pos := range positions {
		grid[pos[0]] |= 1 << uint(pos[1])
	}
	return grid
}

// UnpackBool unpacks a compact grid into a bool grid.
func (grid Grid) UnpackBool(width int) [][]bool {
	if width < 0 || width > 64 {
		return nil
	}
	var boolGrid [][]bool
	for row := range grid {
		boolRow := make([]bool, width)
		for col := 0; col < width; col++ {
			boolRow[col] = grid[row]&(1<<uint(col)) != 0
		}
		boolGrid[row] = boolRow
	}
	return boolGrid
}

// UnpackPos unpacks a compact grid into positions.
func (grid Grid) UnpackPos() []Pos {
	var positions []Pos
	for row := range grid {
		for col := 0; col < 64; col++ {
			if grid[row]&(1<<uint(col)) != 0 {
				positions = append(positions, Pos{row, col})
			}
		}
	}
	return positions
}

func (grid Grid) String() string {
	var b strings.Builder
	for row := range grid {
		for col := 0; col < 64; col++ {
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

func (pos Pos) String() string {
	var b strings.Builder
	b.WriteByte('(')
	b.WriteString(strconv.FormatInt(int64(pos[0]), 10))
	b.WriteByte(' ')
	b.WriteString(strconv.FormatInt(int64(pos[1]), 10))
	b.WriteByte(')')
	return b.String()
}
