package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var f = strings.NewReader(`
addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`[1:])

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scr := bufio.NewScanner(f)
	cycle := 1
	regX := 1
	sum := 0
	for scr.Scan() {
		inst, num := "", 0
		for i, f := range strings.Fields(scr.Text()) {
			switch i {
			case 0:
				inst = f
			case 1:
				n, err := strconv.Atoi(f)
				if err != nil {
					log.Fatal(err)
				}
				num = n
			}
		}

		cycle++
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			fmt.Println("Signal strength for:", cycle, "is", cycle * regX)
			sum += (cycle * regX)
		}
		if inst == "noop" {
			continue
		}
		regX += num
		cycle++
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			fmt.Println("Signal strength for:", cycle, "is", cycle * regX)
			sum += (cycle * regX)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

