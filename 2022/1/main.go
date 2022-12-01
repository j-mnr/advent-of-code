package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var kcalCount []int
	for _, elf := range bytes.Split(f, []byte("\n\n")) {
		sum := 0
		for _, kcal := range bytes.Split(elf, []byte("\n")) {
			if bytes.Equal(kcal, []byte("")) {
				continue
			}
			d, err := strconv.Atoi(string(kcal))
			if err != nil {
				log.Println(err)
				continue
			}
			sum += d
		}
		kcalCount = append(kcalCount, sum)
	}
	sort.Ints(kcalCount)
	fmt.Println("The elf with the most calories has:", kcalCount[len(kcalCount)-1], "calories")
}
