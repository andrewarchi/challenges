package bitgroup

import (
	"reflect"
	"testing"
)

func TestGroups(t *testing.T) {
	tests := []struct {
		Grid   [][]bool
		Groups [][][]bool
	}{
		{
			Grid: [][]bool{
				{true, false, false, true},
				{false, false, true, true},
				{true, true, false, false},
				{true, false, false, false},
			},
			Groups: [][][]bool{
				[][]bool{
					{false, false, false, true},
					{false, false, true, true},
					{true, true, false, false},
					{true, false, false, false},
				},
				[][]bool{
					{true, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
			},
		},
		{
			Grid: [][]bool{
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
			Groups: [][][]bool{
				[][]bool{
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{false, false, true, true, false, true, true, false, false, false},
					{false, false, true, false, true, true, false, false, false, false},
					{true, true, false, false, true, false, false, false, false, false},
					{true, false, false, false, false, false, false, false, false, false},
					{true, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{true, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, true, true, false, false},
					{false, false, true, true, false, true, true, false, false, false},
					{false, false, true, false, true, true, false, false, false, false},
					{true, true, false, false, true, false, false, false, false, false},
					{true, false, false, false, false, false, false, false, false, false},
					{true, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, true, true, false, false},
					{false, false, true, true, false, true, true, false, false, false},
					{false, false, true, false, true, true, false, false, false, false},
				},
				[][]bool{
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, false},
					{false, false, false, false, false, false, false, false, false, true},
					{false, false, false, false, false, false, false, false, false, false},
				},
			},
		},
	}

	for i, test := range tests {
		grid := PackBool(test.Grid)
		groups := grid.Groups()
		if len(groups) != len(test.Groups) {
			t.Errorf("test %d: got len %d, want %d", i, len(groups), len(test.Groups))
			continue
		}
		for g := 0; g < len(groups); g++ {
			want := PackBool(test.Groups[g])
			if !reflect.DeepEqual(groups[g], want) {
				t.Errorf("test %d: group %d not equal", i, g)
				t.Errorf("  got:\n%v", groups[g])
				t.Errorf("  want:\n%v", want)
			}
		}
	}
}
