package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	err := os.WriteFile("notes.txt", []byte("hello file\nline two\nline three\n"), 0644)

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("notes.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("line:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// data, err := os.ReadFile("notes.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(string(data))

	r := strings.NewReader("copied from a string\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	g, err := os.Open("notes.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = g.Close() }()

	if _, err := io.Copy(os.Stdout, g); err != nil {
		log.Fatal(err)
	}

}
