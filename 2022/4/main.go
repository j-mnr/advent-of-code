package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var f = strings.NewReader(`
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`[1:])

type bounds [2]int

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	contained := 0
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		first, second, ok := bytes.Cut(scr.Bytes(), []byte(","))
		if !ok {
			continue
		}
		b1, ok1 := convert(first)
		b2, ok2 := convert(second)
		if !ok1 || !ok2 {
			fmt.Println("conversion failed")
		}
		if b1[0] > b2[0] {
			b1, b2 = b2, b1
		}
		if b1[1] < b2[0] {
			continue
		}
		fmt.Println("contained in together", b1, b2)
		contained++
	}

	fmt.Printf("contained: %d\n", contained)
}

func convert(a []byte) (bounds, bool) {
	n1, n2, ok := bytes.Cut(a, []byte("-"))
	if !ok {
		fmt.Println("bad things happened")
		return bounds{}, false
	}
	start, err := strconv.Atoi(string(n1))
	if err != nil {
		fmt.Println("bad things happened", err)
		return bounds{}, false
	}
	end, err := strconv.Atoi(string(n2))
	if err != nil {
		fmt.Println("bad things happened", err)
		return bounds{}, false
	}
	return bounds{start, end}, true
}
