package main

import (
	"errors"
	"fmt"
	"iter"
)

// TODO: Lines を実装。10行に1回エラーを出し、(zero, err) を1回だけ流して終了。
func Lines(limit int) iter.Seq2[string, error] {
	return nil
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
