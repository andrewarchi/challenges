package gridgrouper

import (
	"reflect"
	"testing"
)

func TestGetGroupsBySquares(t *testing.T) {
	grid := BoolGrid{
		{true, false, false, true},
		{false, false, true, true},
		{true, true, false, false},
		{true, false, false, false},
	}
	expected := []PosSet{
		{Pos{0, 0}},
		{Pos{3, 0}, Pos{2, 1}, Pos{0, 2}, Pos{1, 2}, Pos{0, 3}},
	}
	if groups := getGroupsBySquares(grid); !reflect.DeepEqual(groups, expected) {
		t.Error("Expected matching groups", groups)
	}
}
