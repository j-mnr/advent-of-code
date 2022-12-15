package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const targetY = 2_000_000

var f = strings.NewReader(`
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`[1:])

type (
	coord  [2]int
	bounds [2]int
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	sensors, beacons := parseInput(f)

	var dists []int
	for i := 0; i < len(sensors); i++ {
		dists = append(dists, dist(sensors[i], beacons[i]))
	}

	var intervals []bounds
	for i, s := range sensors {
		dx := dists[i] - abs(s[1]-targetY)
		if dx <= 0 {
			continue
		}
		intervals = append(intervals, bounds{s[0] - dx, s[0] + dx})
	}

	var allowed []int
	for _, b := range beacons {
		if b[1] == targetY {
			allowed = append(allowed, b[0])
		}
	}

	minx, maxx := minMax(intervals)

	result := 0
	for n := minx; n <= maxx; n++ {
		if contains(allowed, n) {
			continue
		}

		for _, in := range intervals {
			if in[0] <= n && n <= in[1] {
				result++
				break
			}
		}
	}

	fmt.Printf("result: %v\n", result)
}

func parseInput(f io.Reader) (sensors, beacons []coord) {
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		flds := strings.Fields(scr.Text())
		sx := convInt(flds[2][2 : len(flds[2])-1])
		sy := convInt(flds[3][2 : len(flds[3])-1])
		bx := convInt(flds[8][2 : len(flds[8])-1])
		by := convInt(flds[9][2:])
		sensors = append(sensors, coord{sx, sy})
		beacons = append(beacons, coord{bx, by})
	}
	return sensors, beacons
}

func contains(a []int, b int) bool {
	for _, x := range a {
		if b == x {
			return true
		}
	}
	return false
}

func minMax(intervals []bounds) (min, max int) {
	min, max = math.MaxInt, math.MinInt
	for _, in := range intervals {
		if in[0] < min {
			min = in[0]
		}
		if in[1] > max {
			max = in[1]
		}
	}
	return min, max
}

func convInt(a string) int {
	n, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return n
}

func dist(a, b coord) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
