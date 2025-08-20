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
		_ = filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fmt.Printf("Name: %s\n", d.Name())
			fmt.Printf("IsDir: %v\n", d.IsDir())
			if d.IsDir() {
				return nil
			}
			if !yield(p) {
				return fs.SkipAll
			}
			return nil
		})
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
