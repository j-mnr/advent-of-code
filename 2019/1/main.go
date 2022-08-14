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
	sum := 0
	for scr := bufio.NewScanner(fl); scr.Scan(); {
		n, err := strconv.Atoi(scr.Text())
		if err != nil {
			log.Fatal(err)
		}
		for n = n/3 - 2; n > 0; n = n/3 - 2 {
			sum += n
		}
	}
	fmt.Println(sum)
}
