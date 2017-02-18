package popcount

import "testing"

var x uint64 = 0x73fa00003a30abed

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(x)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(x)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount4(x)
	}
}
