package main

import (
	"fmt"
	"iter"
)

// モックのページネーションAPI: 1..100 を5個ずつ返す
func FetchPage(cursor string) (items []int, next string) {
	start := 1
	if cursor != "" {
		fmt.Println("API: fetch next with cursor =", cursor)
		fmt.Sscanf(cursor, "c%d", &start)
	}
	items = []int{start, start+1, start+2, start+3, start+4}
	if start+4 >= 100 {
		return items, "" // end
	}
	next = fmt.Sprintf("c%d", start+5)
	fmt.Println("API: returns", items, "next=", next)
	return
}

func StreamAll() iter.Seq[int] {
	return func(yield func(int) bool) {
		cursor := ""
		for {
			items, next := FetchPage(cursor)
			for _, v := range items {
				if !yield(v) { return } // 早期終了はここで止まる
			}
			if next == "" { return }
			cursor = next
		}
	}
}

func main() {
	count := 0
	for v := range StreamAll() {
		fmt.Println("recv:", v)
		count++
		if count >= 20 {
			fmt.Println("20件受け取ったので break")
			break
		}
	}
}
