package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input")
	defer f.Close()
	part1(f)

	f, _ = os.Open("input")
	defer f.Close()
	part2(f)
}

func part2(f io.Reader) {
	var wndw []int
	scn := bufio.NewScanner(f)
	for i := 0; i < 3; i++ {
		scn.Scan()
		wndw = append(wndw, convert(scn.Text()))
	}
	count := 0
	prev := sum(wndw)
	for scn.Scan() {
		wndw = wndw[1:]
		wndw = append(wndw, convert(scn.Text()))
		next := sum(wndw)
		if prev < next {
			count++
		}
		prev = next
	}
	fmt.Println(count)
}

func sum(sl []int) int {
	sum := 0
	for _, n := range sl {
		sum += n
	}
	return sum
}

func part1(f io.Reader) {
	scn := bufio.NewScanner(f)
	scn.Scan()
	prev := convert(scn.Text())
	count := 0
	for scn.Scan() {
		next := convert(scn.Text())
		if prev < next {
			count++
		}
		prev = next
	}
	fmt.Println(count)
}

func convert(s string) int {
	d, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return d
}
