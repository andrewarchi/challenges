package main

import (
	"testing"

	"github.com/andrewarchi/gridgrouper/bitgroup"
)

func BenchmarkRecursive_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroups(getGrid4())
	}
}

func BenchmarkIterativeBySquares_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroupsBySquares(getGrid4())
	}
}

func BenchmarkBitwise_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := bitgroup.Grid{0x9, 0xc, 0x3, 0x1}
		grid.Groups()
	}
}

func BenchmarkRecursive_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroups(getGrid10())
	}
}

func BenchmarkIterativeBySquares_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroupsBySquares(getGrid10())
	}
}

func BenchmarkBitwise_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := bitgroup.Grid{0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34}
		grid.Groups()
	}
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
