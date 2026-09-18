package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Remove("does-not-exist.txt"); err != nil {
		fmt.Println("remove failed:", err)
	}

	count := 5
	fmt.Println(count)

	fmt.Println("done")
}
