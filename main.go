package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		fmt.Println(path)
		return err
	})
}
