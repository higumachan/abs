package main

import (
	"os"
	"path/filepath"
)

func main() {
	path := os.Args[1]
	path, err := filepath.Abs(path)
	if err != nil {
		println(err.Error())
		return
	}
	println(path)
}
