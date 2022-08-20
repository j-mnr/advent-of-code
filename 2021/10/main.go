package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

var r = strings.NewReader(`
{([(<{}[<>[]}>{[]{[(<()>
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{`[1:])

var (
	closing = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
		'>': '<',
	}
	openers = []byte{'{', '[', '(', '<'}
	points  = map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func main() {
	r, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scr := bufio.NewScanner(r)
	var illChars []byte
	for scr.Scan() {
		var stack []byte
		var open byte
		for _, b := range scr.Bytes() {
			if bytes.ContainsRune(openers, rune(b)) {
				stack = append(stack, b)
				continue
			}
			open, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if o, ok := closing[b]; !ok {
				panic("What the hell is this? " + string(b))
			} else if o != open {
				illChars = append(illChars, b)
				break
			}
		}
	}
	score := 0
	for _, b := range illChars {
		score += points[b]
	}
	fmt.Println(score)
}
