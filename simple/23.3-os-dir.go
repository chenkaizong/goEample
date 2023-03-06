package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		fmt.Printf("Directory: %s\n", path)
	} else {
		fmt.Printf("File: %s\n", path)
	}

	return nil
}

func main() {
	root := "./"

	os.Mkdir("abc", 0777)
	os.MkdirAll("abc/bbb/ddd", 0777)

	// 遍历目录
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Printf("Error walking the path %s: %v", root, err)
	}

	os.RemoveAll("abc")
}
