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

// --------- Eager ---------
func EagerSumSquaresEvenFirstK(n, k int) int {
	// 偶数だけ収集
	arr := make([]int, 0, n/2+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			arr = append(arr, i)
		}
	}
	// 2乗
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	if k > len(arr) { k = len(arr) }
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	return sum
}

// --------- Iterator ---------
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
	seq := Count(n)
	seq = Filter(seq, func(x int) bool { return x%2 == 0 })
	seq = Map(seq, func(x int) int { return x * x })
	sum, seen := 0, 0
	for v := range seq {
		sum += v
		seen++
		if seen >= k { break }
	}
	return sum
}

// --------- Channel ---------
func genEvensSquares(ctx context.Context, n int) <-chan int {
	ch := make(chan int, 64)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				v := i * i
				select {
				case ch <- v:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return ch
}

func ChannelSumSquaresEvenFirstK(n, k int) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sum, seen := 0, 0
	for v := range genEvensSquares(ctx, n) {
		sum += v
		seen++
		if seen >= k {
			cancel()
			break
		}
	}
	return sum
}

// --------- Tests & Bench ---------
func TestSameResult(t *testing.T) {
	E := EagerSumSquaresEvenFirstK(N, K)
	I := IterSumSquaresEvenFirstK(N, K)
	C := ChannelSumSquaresEvenFirstK(N, K)
	if !(E == I && I == C) {
		t.Fatalf("results differ: eager=%d iter=%d chan=%d", E, I, C)
	}
}

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
