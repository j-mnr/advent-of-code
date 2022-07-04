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
	yes := make(map[byte]struct{})
	sum := 0
	for _, g := range grps {
		for _, b := range g {
			if !('a' <= b && b <= 'z') {
				continue
			}
			yes[b] = struct{}{}
		}
		sum += len(yes)
		yes = make(map[byte]struct{})
	}
	fmt.Println(sum)
}
