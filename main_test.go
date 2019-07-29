package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/andrewarchi/gridgrouper/bitgroup"
)

func TestDisplayPerformance(t *testing.T) {
	fmt.Println("4x4 grid with 100,000 iterations")
	start := time.Now()
	for i := 0; i < 100000; i++ {
		getGroups(getGrid4())
	}
	fmt.Println("Recursive took", time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		getGroupsBySquares(getGrid4())
	}
	fmt.Println("Iterative by squares took", time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		bitGrid := bitgroup.PackBool(getGrid4())
		bitGrid.Groups()
	}
	fmt.Println("Bitwise took", time.Since(start))

	fmt.Println("\n10x10 grid with 100,000 iterations")
	start = time.Now()
	for i := 0; i < 100000; i++ {
		getGroups(getGrid10())
	}
	fmt.Println("Recursive took", time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		getGroupsBySquares(getGrid10())
	}
	fmt.Println("Iterative by squares took", time.Since(start))
	start = time.Now()
	for i := 0; i < 100000; i++ {
		bitGrid := bitgroup.PackBool(getGrid10())
		bitGrid.Groups()
	}
	fmt.Println("Bitwise took", time.Since(start))
}

func getGrid4() [][]bool {
	return [][]bool{
		{true, false, true, true},
		{false, false, true, false},
		{true, true, false, false},
		{true, false, false, false},
	}
}

func getGrid10() [][]bool {
	return [][]bool{
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
}
