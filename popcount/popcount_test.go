package popcount

import (
	"math/rand"
	"testing"
)

func allEqual(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func TestAllCountsEqual(t *testing.T) {
	r := uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
	for r < 2 {
		r = uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
	}
	results := []int{PopCount(r), PopCount2(r), PopCount3(r), PopCount4(r)}
	if !allEqual(results) {
		t.Error("Not all results are equal")
	}
}
