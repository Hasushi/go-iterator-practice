package main

import (
	"errors"
	"fmt"
	"iter"
)

// TODO: Lines を実装。10行に1回エラーを出し、(zero, err) を1回だけ流して終了。
func Lines(limit int) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for i := 1; i <= limit; i++ {
			var line string
			var err error
			if i%10 == 0 {
				err = errors.New("mock read error at line 10 multiple")
			} else {
				line = fmt.Sprintf("line-%02d", i)
			}
			if !yield(line, err) {
				return
			}
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
