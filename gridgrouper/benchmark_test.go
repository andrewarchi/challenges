package gridgrouper

import (
	"testing"
)

func getBoolGrid4() BoolGrid {
	return BoolGrid{
		{true, false, true, true},
		{false, false, true, false},
		{true, true, false, false},
		{true, false, false, false},
	}
}
func getBoolGrid10() BoolGrid {
	return BoolGrid{
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
func getBitGrid4() BitGrid {
	return BitGrid{0xd, 0x4, 0x3, 0x1}
}
func getBitGrid10() BitGrid {
	return BitGrid{0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34, 0x213, 0xc1, 0x26d, 0x34}
}

func TestBenchmarkGridsEqual(t *testing.T) {
	boolPos4 := getBoolGrid4().PosSet()
	boolPos10 := getBoolGrid10().PosSet()
	bitPos4 := getBitGrid4().PosSet()
	bitPos10 := getBitGrid10().PosSet()
	if !boolPos4.Equal(bitPos4) {
		t.Errorf("boolGrid4: %v, bitGrid4: %v", boolPos4, bitPos4)
	}
	if !boolPos10.Equal(bitPos10) {
		t.Errorf("boolGrid10: %v, bitGrid10: %v", boolPos10, bitPos10)
	}
}

func BenchmarkRecursive_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroups(getBoolGrid4())
	}
}

func BenchmarkIterativeBySquares_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroupsBySquares(getBoolGrid4())
	}
}

func BenchmarkBitwise_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getBitGrid4().Groups()
	}
}

func BenchmarkRecursive_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroups(getBoolGrid10())
	}
}

func BenchmarkIterativeBySquares_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGroupsBySquares(getBoolGrid10())
	}
}

func BenchmarkBitwise_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getBitGrid10().Groups()
	}
}
