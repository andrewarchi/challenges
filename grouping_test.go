package gridgrouper

import (
	"fmt"
	"testing"
)

func TestGroups(t *testing.T) {
	tests := []struct {
		Grid   BoolGrid
		Groups []BoolGrid
	}{
		{
			Grid: BoolGrid{
				{true, false, true, true},
				{false, false, true, false},
				{true, true, false, false},
				{true, false, false, false},
			},
			Groups: []BoolGrid{
				BoolGrid{
					{false, false, true, true},
					{false, false, true, false},
					{true, true, false, false},
					{true, false, false, false},
				},
				BoolGrid{
					{true, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
					{false, false, false, false},
				},
			},
		},
		{
			Grid: BoolGrid{
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
			Groups: []BoolGrid{
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
				BoolGrid{
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
			Grid: BoolGrid{
				{false, true, false, false},
				{true, true, true, true},
			},
			Groups: []BoolGrid{
				BoolGrid{
					{false, true, false, false},
					{true, true, true, true},
				},
			},
		},
		{
			Grid: BoolGrid{
				{false, true, false, false},
				{true, false, false, true},
			},
			Groups: []BoolGrid{
				BoolGrid{
					{false, true, false, false},
					{true, false, false, false},
				},
				BoolGrid{
					{false, false, false, false},
					{false, false, false, true},
				},
			},
		},
		{
			Grid: BoolGrid{
				{false, true, false, false},
				{true, false, false, true},
				{false, true, true, false},
			},
			Groups: []BoolGrid{
				BoolGrid{
					{false, true, false, false},
					{true, false, false, true},
					{false, true, true, false},
				},
			},
		},
	}

	for i, test := range tests {
		grid := test.Grid.BitGrid()
		groups := grid.Groups()
		if len(groups) != len(test.Groups) {
			t.Errorf("test %d: got len %d, want %d", i, len(groups), len(test.Groups))
			continue
		}
		for g := 0; g < len(groups); g++ {
			want := test.Groups[g].BitGrid()
			deepCompare(t, i, fmt.Sprintf("group %d", g), groups[g], want)
		}
	}
}
