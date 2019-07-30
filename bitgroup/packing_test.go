package bitgroup

import (
	"reflect"
	"testing"
)

func TestPacking(t *testing.T) {
	tests := []struct {
		Width int
		Grid  Grid
		Bool  [][]bool
		Pos   []Pos
	}{
		{
			Width: 4,
			Grid:  Grid{0x9, 0xc, 0x3, 0x1},
			Bool: [][]bool{
				{true, false, false, true},
				{false, false, true, true},
				{true, true, false, false},
				{true, false, false, false},
			},
			Pos: []Pos{
				Pos{0, 0}, Pos{0, 3},
				Pos{1, 2}, Pos{1, 3},
				Pos{2, 0}, Pos{2, 1},
				Pos{3, 0},
			},
		},
		{
			Width: 10,
			Grid:  Grid{0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34},
			Bool: [][]bool{
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
			Pos: []Pos{
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
		deepCompare(t, i, "UnpackBool", test.Grid.UnpackBool(test.Width), test.Bool)
		deepCompare(t, i, "UnpackPos", test.Grid.UnpackPos(), test.Pos)
		deepCompare(t, i, "PackBool", PackBool(test.Bool), test.Grid)
		deepCompare(t, i, "PackPos", PackPos(test.Pos), test.Grid)
	}
}

func TestUnpackBoolInvalid(t *testing.T) {
	grid := Grid{0x9, 0xc, 0x3, 0x1}
	deepCompare(t, 0, "UnpackBool(-1)", grid.UnpackBool(-1), [][]bool(nil))
	deepCompare(t, 1, "UnpackBool(64)", grid.UnpackBool(65), [][]bool(nil))
	deepCompare(t, 2, "UnpackBool(100)", grid.UnpackBool(100), [][]bool(nil))
}

func TestGridString(t *testing.T) {
	want := `X  X                                                            
  XX                                                            
XX                                                              
X                                                               
`
	grid := Grid{0x9, 0xc, 0x3, 0x1}
	deepCompare(t, 0, "String", grid.String(), want)
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
