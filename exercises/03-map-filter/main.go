package main

import (
	"fmt"
	"iter"
)

// TODO: Map, Filter を実装
func Map[A any, B any](seq iter.Seq[A], f func(A) B) iter.Seq[B] {
	// TODO
	return nil
}

func Filter[A any](seq iter.Seq[A], pred func(A) bool) iter.Seq[A] {
	// TODO
	return nil
}

func Count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) { return }
		}
	}
}

func main() {
	// 1..100 -> 偶数 -> 2乗 -> 10個見たらbreak
	_ = Count
	_ = Map
	_ = Filter

	// TODO: 実装して動かす
	for v := range Count(10) {
		fmt.Println("placeholder:", v)
		break
	}
}
