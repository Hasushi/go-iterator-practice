package main

import (
	"fmt"
	"iter"
)

// TODO: Map, Filter を実装
func Map[A any, B any](seq iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for v := range seq {
			fmt.Printf("Map: %v\n", v)
			fv := f(v)
			if !yield(fv) {
				return 
			}
		}
	}
}

func Filter[A any](seq iter.Seq[A], pred func(A) bool) iter.Seq[A] {
	return func(yield func(A) bool){
		for v := range seq {
			fmt.Printf("Filter: %v\n", v);
			if !pred(v) { continue };
			if !yield(v) {
				fmt.Println("Filter 終了")
				return 
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
	// 1..100 -> 偶数 -> 2乗 -> 10個見たらbreak
	ci := Count(100)
	fi := Filter(ci, func(x int) bool { return x%2 == 0 })
	mi := Map(fi, func(x int) int { return x*x })

	// TODO: 実装して動かす
	i := 0
	for v := range mi {
		fmt.Printf("%d\n",v)
		i++
		if i ==10 { break }
	}
}
