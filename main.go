package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, f os.FileInfo, err Error) error {
		fmt.Println(path)
	})
	fmt.Printf(context.Background().Value())
}
