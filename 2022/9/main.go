package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// R 4
// U 4
// L 3
// D 1
// R 4
// D 1
// L 5
// R 2
var f = strings.NewReader(`
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2 `[1:])

type direction uint8

const (
	unknown = iota
	up
	down
	left
	right
)

type coordinate [2]int

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scr := bufio.NewScanner(f)
	allCoords := map[coordinate]struct{}{}
	head, tail := coordinate{10000, 10000}, coordinate{10000, 10000}
	for scr.Scan() {
		f := strings.Fields(scr.Text())
		dir := Direction(f[0])
		steps, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < steps; i++ {
			switch dir {
			case up:
				head[0]++
			case down:
				head[0]--
			case left:
				head[1]--
			case right:
				head[1]++
			}
			catch(&head, &tail)
			allCoords[tail] = struct{}{}
		}
	}
	fmt.Println("All unique positions visited:", len(allCoords))
}

func Direction(s string) direction {
	switch s {
	case "R":
		return right
	case "D":
		return down
	case "U":
		return up
	case "L":
		return left
	default:
		panic(s + " is not a direction")
	}
}

func catch(head, tail *coordinate) {
	dif0, dif1 := abs(head[0])-abs(tail[0]), abs(head[1])-abs(tail[1])
	if (dif0 == 0 && dif1 == 0) || (dif0 == 1 && dif1 == 1) ||
		(dif0 == 1 && dif1 == 0) || (dif0 == 0 && dif1 == 1) {
		return
	}
	if (abs(dif0) == 2 && abs(dif1) == 1) || (abs(dif0) == 1 && abs(dif1) == 2) { // diagnoal
		switch {
		case dif0 < 0 && dif1 < 0:
			tail[0]--
			tail[1]--
		case dif0 >= 0 && dif1 < 0:
			tail[0]++
			tail[1]--
		case dif0 < 0 && dif1 >= 0:
			tail[0]--
			tail[1]++
		case dif0 >= 0 && dif1 >= 0:
			tail[0]++
			tail[1]++
		}
		return
	}

	switch {
	case abs(dif0) == 2:
		if dif0 < 0 {
			tail[0]--
			return
		}
		tail[0]++
	case abs(dif1) == 2:
		if dif1 < 0 {
			tail[1]--
			return
		}
		tail[1]++
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
