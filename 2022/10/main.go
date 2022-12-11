package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scr := bufio.NewScanner(f)
	spritePos := 0
	var img strings.Builder
	for cycle := 0; scr.Scan(); cycle++ {
		flds := strings.Fields(scr.Text())
		switch flds[0] {
		case "noop":
			write(&img, spritePos, cycle)
		default: // addx instruction
			write(&img, spritePos, cycle)

			cycle++
			write(&img, spritePos, cycle)
			n, err := strconv.Atoi(flds[1])
			if err != nil {
				log.Fatal(err)
			}
			spritePos += n
		}
	}
	fmt.Printf("%v\n", img.String())
}

func write(img *strings.Builder, spritePos, cycle int) {
	switch cycle {
	case 40, 80, 120, 160, 200, 240:
		img.WriteByte('\n')
	}
	if spritePos <= cycle%40 && cycle%40 <= spritePos+2 {
		img.WriteRune('█')
		return
	}
	img.WriteByte(' ')
}
