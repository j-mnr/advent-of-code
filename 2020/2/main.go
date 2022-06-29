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
		posns, ltr, pass := strings.Split(flds[0], "-"), flds[1][0], flds[2]
		pos1, err := strconv.Atoi(posns[0])
		if err != nil {
			log.Fatal(err)
		}
		pos2, err := strconv.Atoi(posns[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Println(pos1-1, pos2-1, string(ltr), pass)

		switch {
		case pass[pos1-1] == ltr && pass[pos2-1] == ltr:
			// no-op
		case pass[pos1-1] == ltr || pass[pos2-1] == ltr:
			valid++
		}
	}
	log.Println(valid)
}
