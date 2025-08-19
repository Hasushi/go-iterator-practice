package main

import (
	"fmt"
	"iter"
)

func Naturals() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; ; i++ {
			if !yield(i) { return }
		}
	}
}

func main() {
	next, stop := iter.Pull(Naturals())
	defer stop()

	for i := 0; i < 10; i++ {
		v, ok := next()
		if !ok { break }
		fmt.Println(v)
	}
}
