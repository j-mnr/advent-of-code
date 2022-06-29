package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var f = strings.NewReader(`
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`[1:])

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	valid := 0
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		flds := strings.Fields(scr.Text())
		minmax, ltr, pass := strings.Split(flds[0], "-"), flds[1][0], flds[2]
		min, err := strconv.Atoi(minmax[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Println(min, max, string(ltr), pass)

		if c := strings.Count(pass, string(ltr)); c >= min && c <= max {
			valid++
		}
	}
	log.Println(valid)
}
