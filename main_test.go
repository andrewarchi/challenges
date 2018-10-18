package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDisplayPerformance(t *testing.T) {
	grid := [][]bool{
		{true, false, true, true},
		{false, false, true, false},
		{true, true, false, false},
		{true, false, false, false},
	}
	fmt.Println("4x4 grid with 100,000 iterations")
	start := time.Now()
	for i := 0; i < 100000; i++ {
		getGroups(grid)
	}
	fmt.Println("Recursive took", time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		getGroupsBySquares(grid)
	}
	fmt.Println("Iterative by squares took", time.Since(start))

	grid = [][]bool{
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
	}
	fmt.Println("\n10x10 grid with 100,000 iterations")
	start = time.Now()
	for i := 0; i < 100000; i++ {
		getGroups(grid)
	}
	fmt.Println("Recursive took", time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		getGroupsBySquares(grid)
	}
	fmt.Println("Iterative by squares took", time.Since(start))
}
