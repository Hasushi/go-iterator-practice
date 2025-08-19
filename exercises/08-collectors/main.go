package main

import (
	"fmt"
	"iter"
)

// TODO: Collect / Take / First を実装
func Collect[T any](seq iter.Seq[T]) []T               { return nil }
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T]   { return nil }
func First[T any](seq iter.Seq[T]) (zero T, ok bool)   { return zero, false }

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
	// 1..100 の偶数の先頭5個を収集
	fmt.Println("TODO")
}
