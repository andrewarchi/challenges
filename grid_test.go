package gridgrouper

import (
	"reflect"
	"testing"
)

func TestConversions(t *testing.T) {
	tests := []struct {
		BitGrid  BitGrid
		BoolGrid BoolGrid
		PosSet   PosSet
	}{
		{
			BitGrid: BitGrid{0xd, 0x4, 0x3, 0x1},
			BoolGrid: BoolGrid{
				{true, false, true, true},
				{false, false, true, false},
				{true, true, false, false},
				{true, false, false, false},
			},
			PosSet: PosSet{
				Pos{0, 0}, Pos{0, 2}, Pos{0, 3},
				Pos{1, 2},
				Pos{2, 0}, Pos{2, 1},
				Pos{3, 0},
			},
		},
		{
			BitGrid: BitGrid{0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34},
			BoolGrid: BoolGrid{
				{true, false, true, true, false, true, true, false, false, true},
				{false, false, true, false, true, true, false, false, false, false},
				{true, true, false, false, true, false, false, false, false, true},
				{true, false, false, false, false, false, true, true, false, false},
				{true, false, true, true, false, true, true, false, false, true},
				{false, false, true, false, true, true, false, false, false, false},
				{true, true, false, false, true, false, false, false, false, true},
				{true, false, false, false, false, false, true, true, false, false},
				{true, false, true, true, false, true, true, false, false, true},
				{false, false, true, false, true, true, false, false, false, false},
			},
			PosSet: PosSet{
				Pos{0, 0}, Pos{0, 2}, Pos{0, 3}, Pos{0, 5}, Pos{0, 6}, Pos{0, 9},
				Pos{1, 2}, Pos{1, 4}, Pos{1, 5},
				Pos{2, 0}, Pos{2, 1}, Pos{2, 4}, Pos{2, 9},
				Pos{3, 0}, Pos{3, 6}, Pos{3, 7},
				Pos{4, 0}, Pos{4, 2}, Pos{4, 3}, Pos{4, 5}, Pos{4, 6}, Pos{4, 9},
				Pos{5, 2}, Pos{5, 4}, Pos{5, 5},
				Pos{6, 0}, Pos{6, 1}, Pos{6, 4}, Pos{6, 9},
				Pos{7, 0}, Pos{7, 6}, Pos{7, 7},
				Pos{8, 0}, Pos{8, 2}, Pos{8, 3}, Pos{8, 5}, Pos{8, 6}, Pos{8, 9},
				Pos{9, 2}, Pos{9, 4}, Pos{9, 5},
			},
		},
	}

	for i, test := range tests {
		deepCompare(t, i, "BitGrid.BoolGrid()", test.BitGrid.BoolGrid(), test.BoolGrid)
		deepCompare(t, i, "BitGrid.PosSet()", test.BitGrid.PosSet(), test.PosSet)
		deepCompare(t, i, "BoolGrid.BitGrid()", test.BoolGrid.BitGrid(), test.BitGrid)
		deepCompare(t, i, "BoolGrid.PosSet()", test.BoolGrid.PosSet(), test.PosSet)
		deepCompare(t, i, "PosSet.BitGrid()", test.PosSet.BitGrid(), test.BitGrid)
		deepCompare(t, i, "PosSet.BoolGrid()", test.PosSet.BoolGrid(), test.BoolGrid)
	}
}

func TestBitGridString(t *testing.T) {
	want := `X  X
  XX
XX  
X   
`
	grid := BitGrid{0x9, 0xc, 0x3, 0x1}
	deepCompare(t, 0, "BitGrid.String()", grid.String(), want)
}

func TestBoolGridString(t *testing.T) {
	want := `X  X
  XX
XX  
X   
`
	grid := BoolGrid{
		{true, false, false, true},
		{false, false, true, true},
		{true, true, false, false},
		{true, false, false, false},
	}
	deepCompare(t, 0, "BoolGrid.String()", grid.String(), want)
}

func TestPosString(t *testing.T) {
	want := "(1 2)"
	pos := Pos{1, 2}
	deepCompare(t, 0, "String", pos.String(), want)
}

func deepCompare(t *testing.T, index int, label string, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("test %d: %s\ngot:\n%v\nwant:\n%v", index, label, got, want)
	}
}
