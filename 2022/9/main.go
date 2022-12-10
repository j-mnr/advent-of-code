package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord [2]int

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	knots := make([]coord, 10)
	tailVisit := make(map[coord]struct{})
	tailVisit[knots[0]] = struct{}{}
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		op, amount := parseLine(scr.Text())
		dx, dy := direction(op)
		for i := 0; i < amount; i++ {
			move(knots, dx, dy)
			tailVisit[knots[len(knots)-1]] = struct{}{}
		}
	}
	fmt.Printf("knots: %v\n", knots)
	fmt.Println(len(tailVisit))
}

func direction(op string) (x, y int) {
	switch op {
	case "R":
		return 1, 0
	case "U":
		return 0, 1
	case "L":
		return -1, 0
	case "D":
		return 0, -1
	}
	panic("No other direction exists!")
}

func parseLine(s string) (string, int) {
	flds := strings.Fields(s)
	amount, err := strconv.Atoi(flds[1])
	if err != nil {
		log.Fatal(err)
	}
	return flds[0], amount
}

func move(knots []coord, dx, dy int) {
	knots[0][0] += dx
	knots[0][1] += dy

	for i := 1; i < 10; i++ {
		tx, ty := knots[i][0], knots[i][1]

		if !touching(knots[i-1], knots[i]) {
			hx, hy := knots[i-1][0], knots[i-1][1]
			signx, signy := 0, 0
			if hx != tx {
				signx = (hx - tx) / abs(hx-tx)
			}
			if hy != ty {
				signy = (hy - ty) / abs(hy-ty)
			}

			tx += signx
			ty += signy
		}
		knots[i] = coord{tx, ty}
	}
}

func touching(c1, c2 coord) bool {
	return abs(c1[0]-c2[0]) <= 1 && abs(c1[1]-c2[1]) <= 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
