package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	split := bytes.Split(b, []byte{','})
	fish := make([]int, 0, len(split)*2)
	for _, bsl := range split {
		fish = append(fish, int(bsl[0] - '0'))
	}
	for i := 0; i < 80; i++ {
		for j := range fish {
			fish[j]--
			if fish[j] < 0 {
				fish[j] = 6
				fish = append(fish, 8)
			}
		}
	}
	fmt.Println(len(fish))
}
