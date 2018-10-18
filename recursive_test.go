package main

import (
	"fmt"
	"testing"
)

func TestGetGroups(t *testing.T) {
	grid := [][]bool{
		{true, false, true, true},
		{false, false, true, false},
		{true, true, false, false},
		{true, false, false, false},
	}
	fmt.Println(getGroups(grid))
}
