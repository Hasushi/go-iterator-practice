package benchstrings

import (
	"fmt"
	"testing"
	"iter"
)

const N = 200_000

// 疑似行を生成（"row-000001-foobar" のように）
func genLines(n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := 0; i < n; i++ {
			line := fmt.Sprintf("row-%06d-foobar", i) // ここで文字列確保が発生
			if !yield(line) { return }
		}
	}
}

// TODO: Eager と Iterator の2パターンを実装し、len>8 のものだけ length を合計

func EagerSumLength(n int) int {
	return 0
}

func IterSumLength(n int) int {
	return 0
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
