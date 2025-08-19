package main

import (
	"fmt"
	"iter"
)

func Map[A any, B any](seq iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range seq {
			b := f(a)
			if !yield(b) { return }
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

func Count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) { return }
		}
	}
}

func main() {
	seq := Count(100)
	seq = Filter(seq, func(x int) bool { return x%2 == 0 })   // 偶数
	m := Map(seq, func(x int) int { return x * x })           // 2乗

	seen := 0
	for v := range m {
		fmt.Println(v)
		seen++
		if seen >= 10 {
			break
		}
	}
}
