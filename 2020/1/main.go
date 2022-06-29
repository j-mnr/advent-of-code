package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

var input = strings.NewReader(`
1721
979
366
299
675
1456`[1:])

const target = 2020

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[int]int, 200)
	for _, line := range bytes.Split(f, []byte("\n"))[:200] {
		n, err := strconv.Atoi(string(line))
		if err != nil {
			log.Println(err)
			continue
		}
		m[n]++
	}

	for y := range m {
		s := target - y
		for x := range m {
			z := s - x
			if _, found := m[z]; found {
				log.Printf("%d+%d+%d=2020 | %d*%d*%d=%d",
					z, y, x, z, y, x, z*y*x)
				return
			}
		}
	}
}
