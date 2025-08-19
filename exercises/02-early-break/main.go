package main

import (
	"fmt"
	"iter"
)

// TODO: Count(n) を実装
func Count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		// TODO
	}
}

func isLCM35(x int) bool { // 5と7の公倍数
	return x%5 == 0 && x%7 == 0
}

func main() {
	for v := range Count(10000) {
		fmt.Println("受信:", v)
		if isLCM35(v) {
			fmt.Println("ループ: 5と7の公倍数が見つかったのでbreak ->", v)
			break
		}
	}
}
