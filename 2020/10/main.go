package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var test = []int{
	16,
	10,
	15,
	5,
	1,
	11,
	7,
	19,
	6,
	12,
	4,
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	test = make([]int, 100)
	for i, b := range bytes.Split(f, []byte("\n")) {
		n, err := strconv.Atoi(string(b))
		if err != nil {
			break
		}
		test[i] = n
	}
	sort.Ints(test)

	diffOne, diffThree := 0, 0
	for i := 0; i < len(test)-1; i++ {
		if i == len(test)-2 {
			fmt.Println(test[i+1])
		}
		switch test[i+1] - test[i] {
		case 1:
			diffOne++
		case 3:
			diffThree++
		default:
			fmt.Println(test[i+1], test[i])
		}
	}
	fmt.Println(diffOne, diffThree, (diffOne+1) * (diffThree+1))
}
