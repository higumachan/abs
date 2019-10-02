package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := os.Args[1]

	_, err2 := os.Stat(path)
	if os.IsNotExist(err2) {
		fmt.Printf("abs: \"%s\": No such file or directory\n", path)
		return
	}

	path, err := filepath.Abs(path)
	if err != nil {
		println(err.Error())
		return
	}

	println(path)
}
