package id64_test

import (
	"fmt"
	"testing"

	"github.com/kokizzu/id64"
	"github.com/stretchr/testify/assert"
)

func BenchmarkId64(b *testing.B) {
	for z := 0; z < b.N; z++ {
		res := id64.Gen.ID()
		_ = res
	}
}

func TestFromStr(t *testing.T) {
	i := id64.FromStr(`3KuBw----0`)
	fmt.Println(i)
	assert.Equal(t, uint64(78224544304726017), uint64(i))
}
