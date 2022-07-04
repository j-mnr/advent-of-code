package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	grps := bytes.Split(f, []byte("\n\n"))
	sum := 0
	for _, g := range grps {
		everyone := make(map[byte]int)
		for _, b := range g {
			if !('a' <= b && b <= 'z') {
				continue
			}
			everyone[b]++
		}

		nMembers := len(bytes.Split(g, []byte("\n")))
		for _, cnt := range everyone {
			if nMembers == cnt {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
