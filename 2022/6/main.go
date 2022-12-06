package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stream, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	seen := [14]byte{}
	for i, b := range stream[:14] {
		seen[i] = b
	}
	for i, b := range stream[14:] {
		if dupl(seen) {
			seen = rotate(seen, b)
			continue
		}
		fmt.Println(seen)
		fmt.Println("stream start:", i+14)
		break
	}
}

func dupl(seen [14]byte) bool {
	for i, a := range seen {
		for _, b := range seen[i+1:] {
			if a == b {
				return true
			}
		}
	}
	return false
}

func rotate(a [14]byte, b byte) [14]byte {
	for i := 0; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a[13] = b
	return a
}
