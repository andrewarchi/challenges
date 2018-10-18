package main

import (
	"reflect"
	"testing"
)

func TestGetGroupsBySquares(t *testing.T) {
	grid := [][]bool{
		{true, false, false, true},
		{false, false, true, true},
		{true, true, false, false},
		{true, false, false, false},
	}
	expected := [][]position{
		{position{0, 0}},
		{position{3, 0}, position{2, 1}, position{0, 2}, position{1, 2}, position{0, 3}},
	}
	if groups := getGroupsBySquares(grid); !reflect.DeepEqual(groups, expected) {
		t.Error("Expected matching groups", groups)
	}
}
