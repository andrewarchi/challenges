package gridgrouper

import (
	"fmt"
	"testing"
)

func TestGetGroups(t *testing.T) {
	grid := BoolGrid{
		{true, false, true, true},
		{false, false, true, false},
		{true, true, false, false},
		{true, false, false, false},
	}
	fmt.Println(getGroups(grid))
}
