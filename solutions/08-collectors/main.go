package main

import (
	"fmt"
	"iter"
)

func Collect[T any](seq iter.Seq[T]) []T {
	var out []T
	for v := range seq {
		out = append(out, v)
	}
	return out
}

func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for v := range seq {
			if !yield(v) { return }
			count++
			if count >= n { return }
		}
	}
}

func First[T any](seq iter.Seq[T]) (zero T, ok bool) {
	next, stop := iter.Pull(seq)
	defer stop()
	v, ok := next()
	return v, ok
}

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

func main() {
	evens := Filter(Count(100), func(x int) bool { return x%2 == 0 })
	first5 := Take(evens, 5)
	fmt.Println("collect first 5 evens:", Collect(first5))

	// 最初の偶数
	evens2 := Filter(Count(100), func(x int) bool { return x%2 == 0 })
	v, ok := First(evens2)
	fmt.Println("first even:", v, ok)
}
