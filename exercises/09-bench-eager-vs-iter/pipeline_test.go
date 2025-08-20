package bencheageriter

import (
	"context"
	"testing"
	"iter"
)

const (
	N = 1_000_000
	K = 50_000
)

// ========== TODO: 実装対象（Eager / Iterator / Channel） ==========

// Eager: 中間スライスを作る実装
func EagerSumSquaresEvenFirstK(n, k int) int {
	// TODO: 1..n から偶数だけをスライスに集め、2乗して、先頭k個を合計
	return 0
}

// Iterator: iter.Seq を使う実装
func Count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) { return }
		}
	}
}

func Filter[A any](seq iter.Seq[A], pred func(A) bool) iter.Seq[A] {
	return func(yield func(A) bool) {
		for a := range seq {
			if pred(a) {
				if !yield(a) { return }
			}
		}
	}
}

func Map[A any, B any](seq iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range seq {
			if !yield(f(a)) { return }
		}
	}
}

func IterSumSquaresEvenFirstK(n, k int) int {
	// TODO: Count -> Filter(偶数) -> Map(2乗) を合成し、先頭k個を合計
	return 0
}

// Channel: goroutine + chanでストリーム。中断可能にするため context を使う
func genEvensSquares(ctx context.Context, n int) <-chan int {
	// TODO: バッファ少なめのチャネルを作成し、偶数の2乗を送る。
	// ctx.Done() を選択して中断に対応。
	return nil
}

func ChannelSumSquaresEvenFirstK(n, k int) int {
	// TODO: ctx を作り、k 個合計したら cancel して終了
	return 0
}

// ========== correctness test ==========

func TestSameResult(t *testing.T) {
	E := EagerSumSquaresEvenFirstK(N, K)
	I := IterSumSquaresEvenFirstK(N, K)
	C := ChannelSumSquaresEvenFirstK(N, K)
	if !(E == I && I == C) {
		t.Fatalf("results differ: eager=%d iter=%d chan=%d", E, I, C)
	}
}

// ========== benchmarks ==========

func BenchmarkEager(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = EagerSumSquaresEvenFirstK(N, K)
	}
}

func BenchmarkIter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = IterSumSquaresEvenFirstK(N, K)
	}
}

func BenchmarkChannel(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ChannelSumSquaresEvenFirstK(N, K)
	}
}
