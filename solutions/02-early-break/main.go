package main

import (
	"fmt"
	"iter"
)

func Count(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		defer fmt.Println("イテレータ: 片付け完了")
		for i := 1; i <= n; i++ {
			fmt.Println("イテレータ: 送出直前", i)
			if !yield(i) {
				fmt.Println("イテレータ: yield=false を受けたので停止")
				return
			}
		}
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
