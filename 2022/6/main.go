package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	f   = strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb") // 7
	fiv = strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz")
	six = strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg")
	ten = strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	elv = strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
)

func main() {
	// marker(f)
	// marker(fiv)
	// marker(six)
	// marker(ten)
	// marker(elv)
	stream, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	marker(stream)
}

func marker(stream []byte) {
	seen := [4]byte{stream[0], stream[1], stream[2], stream[3]}
	for i, b := range stream[4:] {
		if dupl(seen) {
			seen = rotate(seen, b)
			continue
		}
		// fmt.Println(seen)
		fmt.Println("stream start:", i+4)
		break
	}
}

func dupl(seen [4]byte) bool {
	for i, a := range seen {
		for _, b := range seen[i+1:] {
			if a == b {
				return true
			}
		}
	}
	return false
}

func rotate(a [4]byte, b byte) [4]byte {
	a[0], a[1], a[2], a[3] = a[1], a[2], a[3], b
	return a
}
