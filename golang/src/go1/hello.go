package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello world", os.Args[1])
	}

	os.Exit(100)
}
