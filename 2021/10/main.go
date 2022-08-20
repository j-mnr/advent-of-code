package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

var (
	closing = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
		'>': '<',
	}
	openers = []byte{'{', '[', '(', '<'}
	points  = map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
)

func main() {
	r, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	var scores []int
	scr := bufio.NewScanner(r)
	for scr.Scan() {
		var stack []byte
		var open byte
		var corrupted bool
		for _, b := range scr.Bytes() {
			if bytes.ContainsRune(openers, rune(b)) {
				stack = append(stack, b)
				continue
			}
			open, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if o := closing[b]; o != open {
				corrupted = true
				break
			}
		}
		if corrupted {
			continue
		}
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score = score*5 + points[stack[i]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
