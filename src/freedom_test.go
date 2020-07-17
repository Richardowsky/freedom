package src

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var n = 6
var k = 2
var array = []int{2, 3, 6, 5, 4, 10}
var n2 = 10
var k2 = 2
var array2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Test1(t *testing.T) {
	excpect := 3
	assert.Equal(t, Solution(n, k, array), excpect)
}

func Test2(t *testing.T) {
	excpect := 6
	assert.Equal(t, Solution(n2, k2, array2), excpect)
}

func Benchmark1(b *testing.B) {
	benchmarkZero(n, k, array, b)
}

func Benchmark2(b *testing.B) {
	benchmarkZero(n2, k2, array2, b)
}

func benchmarkZero(n, k int, array []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solution(n, k, array)
	}
}
