package main

import (
	"fmt"
	"iter"
)

// TODO: Fibonacci(limit int) iter.Seq[int] を実装
func Fibonacci(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// TODO: a,b を使って 0,1,1,2,... を生成し、yield(a)。
		var a, b int = 0, 1
		for a <= limit {
			if !yield(a) {
				fmt.Println("yieldがfalseを返したので終了")
				return
			}
			a, b = b, a + b
		}
	}
}

func main() {
	for num := range Fibonacci(100) {
		fmt.Printf("ループ: %d を受け取りました\n", num)
		if num > 20 {
			fmt.Println("ループ: 20を超えたのでbreak")
			break
		}
	}
}
