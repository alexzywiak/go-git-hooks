package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Printf(context.Background().Value())
}
