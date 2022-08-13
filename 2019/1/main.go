package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fl, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scr := bufio.NewScanner(fl)
	sum := 0
	for scr.Scan() {
		n, err := strconv.Atoi(scr.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += (n/3 - 2)
	}
	fmt.Println(sum)
}
