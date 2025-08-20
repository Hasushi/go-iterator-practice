package benchstrings

import (
	"fmt"
	"testing"
	"iter"
)

const N = 200_000

func genLines(n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := 0; i < n; i++ {
			line := fmt.Sprintf("row-%06d-foobar", i)
			if !yield(line) { return }
		}
	}
}

// Eager: すべての行を []string に保持してから処理
func EagerSumLength(n int) int {
	arr := make([]string, 0, n)
	for s := range genLines(n) {
		arr = append(arr, s)
	}
	sum := 0
	for _, s := range arr {
		if len(s) > 8 {
			sum += len(s)
		}
	}
	return sum
}

// Iterator: 逐次処理（保持しない）
func IterSumLength(n int) int {
	sum := 0
	for s := range genLines(n) {
		if len(s) > 8 {
			sum += len(s)
		}
	}
	return sum
}

func TestSame(t *testing.T) {
	if EagerSumLength(10_000) != IterSumLength(10_000) {
		t.Fatal("results differ")
	}
}

func BenchmarkEager(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EagerSumLength(N)
	}
}

func BenchmarkIter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = IterSumLength(N)
	}
}
