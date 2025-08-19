package main

import (
	"errors"
	"fmt"
	"iter"
)

func Lines(limit int) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for i := 1; i <= limit; i++ {
			if i%10 == 0 {
				// エラーを1回だけ流して終了
				_ = yield("", errors.New("mock read error at line 10 multiple"))
				return
			}
			line := fmt.Sprintf("line-%02d", i)
			if !yield(line, nil) { return }
		}
	}
}

func main() {
	for line, err := range Lines(35) {
		if err != nil {
			fmt.Println("エラーを受信:", err)
			break
		}
		fmt.Println("line:", line)
	}
}
