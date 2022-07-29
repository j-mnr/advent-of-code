package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var b = strings.NewReader(`
F10
N3
F7
R90
F11`[1:])

//go:generate go run github.com/dmarkham/enumer -type=action
type action int8

const (
	unknown action = iota
	n              // north
	e              // east
	s              // south
	w              // west
	l              // left
	r              // right
	f              // forward
)

type ship struct {
	currDirection action
	ewPosition    int
	nsPosition    int
}

// rotate will take in left or right and rotate the direction by a certain
// amount of degrees
func rotate(amount int, way, dir action) action {
	amount = amount/90 - 1
	switch way {
	case l: // left
		if dir = dir - action(amount) - 1; dir < 1 {
			dir += 4
		}
		return dir
	case r: // right
		if dir = dir + action(amount) + 1; dir > 4 {
			dir -= 4
		}
		return dir
	default:
		panic("HOW DARE YOU ENTER MY DOMAIN!")
	}
}

func (ship *ship) updatePos(a action, amount int) {
	switch a {
	case n: // north
		ship.nsPosition += amount
	case s: // south
		ship.nsPosition -= amount
	case e: // east
		ship.ewPosition += amount
	case w: // west
		ship.ewPosition -= amount
	case l: // left
		ship.currDirection = rotate(amount, a, ship.currDirection)
	case r: // right
		ship.currDirection = rotate(amount, a, ship.currDirection)
	case f: // forward
		ship.updatePos(ship.currDirection, amount)
	}
}

func main() {
	b, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scr := bufio.NewScanner(b)
	ship := ship{currDirection: e}
	for scr.Scan() {
		instructions := scr.Text()
		a, err := actionString(string(instructions[0]))
		if err != nil {
			log.Fatal(err)
		}
		amount, err := strconv.Atoi(instructions[1:])
		ship.updatePos(a, amount)
	}
	fmt.Println(math.Abs(float64(ship.nsPosition)) + math.Abs(float64(ship.ewPosition)))
}
