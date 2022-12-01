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

// var f = strings.NewReader("R8,U5,L5,D3\nU7,R6,D4,L4")

// var f = strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7")

// var f = strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83")

type point struct{ x, y int }

type point2 struct {
	x, y  int
	steps int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	first, second := make(map[point]struct{}), make(map[point]struct{})
	steps1, steps2 := make(map[point]int), make(map[point]int)
	scr := bufio.NewScanner(f)

	scr.Scan() // first wire's points
	appnd := func(p point, steps int) {
		first[p] = struct{}{}
		steps1[p] = steps
	}
	origin, steps := &point{}, 0
	for _, vector := range strings.Split(scr.Text(), ",") {
		parsePoints(appnd, origin, &steps, vector)
	}

	scr.Scan() // second wire's points
	appnd = func(p point, steps int) {
		second[p] = struct{}{}
		steps2[p] = steps
	}
	origin, steps = &point{}, 0
	for _, vector := range strings.Split(scr.Text(), ",") {
		parsePoints(appnd, origin, &steps, vector)
	}

	var intersections []point2
	for p := range first {
		if _, ok := second[p]; ok {
			intersections = append(intersections, point2{
				x:     p.x,
				y:     p.y,
				steps: steps1[p] + steps2[p],
			})
		}
	}

	min := math.MaxInt64
	for _, p := range intersections {
		if min > p.steps {
			min = p.steps
		}
	}
	fmt.Printf("intersections: %v\n", intersections)
	fmt.Printf("min: %v\n", min)
}

func parsePoints(appnd func(p point, steps int), origin *point, steps *int, vector string) {
	d, err := strconv.Atoi(vector[1:])
	if err != nil {
		log.Fatal(err)
	}
	switch vector[0] {
	case 'R':
		for i := 1; i <= d; i++ {
			origin.x++
			*steps++
			appnd(*origin, *steps)
		}
	case 'U':
		for i := 1; i <= d; i++ {
			origin.y++
			*steps++
			appnd(*origin, *steps)
		}
	case 'L':
		for i := 1; i <= d; i++ {
			origin.x--
			*steps++
			appnd(*origin, *steps)
		}
	case 'D':
		for i := 1; i <= d; i++ {
			origin.y--
			*steps++
			appnd(*origin, *steps)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
