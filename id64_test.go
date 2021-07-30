package id64

import "testing"

func BenchmarkId64(b *testing.B) {
	for z := 0; z < b.N; z++ {
		res := Gen.ID()
		_ = res
	}
}
