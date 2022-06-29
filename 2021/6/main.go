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
	var ctr [9]uint64
	split := bytes.Split(b, []byte{','})
	for _, bsl := range split {
		ctr[int8(bsl[0]-'0')]++
	}
	fmt.Println(ctr)
	for i := 0; i < 256; i++ {
		tmp := ctr[0]
		for j := 0; j < 9; j++ {
			if j == len(ctr)-1 {
				ctr[j] = tmp
				ctr[len(ctr)-3] += tmp
				break
			}
			ctr[j] = ctr[j+1]
		}
	}
	fmt.Println(ctr, add(ctr))
}

func add(s [9]uint64) uint64 {
	var sum uint64
	for _, n := range s {
		sum += uint64(n)
	}
	return sum
}
