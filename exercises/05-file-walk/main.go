package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"iter"
)

func WalkFiles(root string) iter.Seq[string] {
	return func(yield func(string) bool) {
		// TODO: filepath.WalkDir を使ってファイルだけ yield
		// yield=false のときは fs.SkipAll で中断
	}
}

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	for p := range WalkFiles(root) {
		fmt.Println("found:", p)
		if filepath.Ext(p) == ".go" {
			fmt.Println("`.go` を見つけたので break:", p)
			break
		}
	}
}
