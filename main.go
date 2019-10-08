package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "."
	}

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
