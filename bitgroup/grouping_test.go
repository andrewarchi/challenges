package bitgroup

import (
	"fmt"
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
		{
			Grid: [][]bool{
				{false, true, false, false},
				{true, true, true, true},
			},
			Groups: [][][]bool{
				[][]bool{
					{false, true, false, false},
					{true, true, true, true},
				},
			},
		},
		{
			Grid: [][]bool{
				{false, true, false, false},
				{true, false, false, true},
			},
			Groups: [][][]bool{
				[][]bool{
					{false, true, false, false},
					{true, false, false, false},
				},
				[][]bool{
					{false, false, false, false},
					{false, false, false, true},
				},
			},
		},
		{
			Grid: [][]bool{
				{false, true, false, false},
				{true, false, false, true},
				{false, true, true, false},
			},
			Groups: [][][]bool{
				[][]bool{
					{false, true, false, false},
					{true, false, false, true},
					{false, true, true, false},
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
			deepCompare(t, i, fmt.Sprintf("group %d", g), groups[g], want)
		}
	}
}

func BenchmarkGroups_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := Grid{0x9, 0xc, 0x3, 0x1} // same as 4x4 grid above
		grid.Groups()
	}
}

func BenchmarkGroups_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := Grid{0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34} // same as 10x10 grid above
		grid.Groups()
	}
}
