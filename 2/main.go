package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions, err := os.Open("input")
	check(err)
	defer instructions.Close()

	horz, depth := followCourse(instructions)
	fmt.Println(horz * depth)
}

// followCourse calculates the final horizontal position and depth from a set
// of submarine instructions.
func followCourse(instructions io.Reader) (horz int, depth int) {
	scn := bufio.NewScanner(instructions)
	aim := 0
	for scn.Scan() {
		cmd := strings.Fields(scn.Text())
		amt, err := strconv.Atoi(cmd[1]) 
		check(err)

		switch cmd[0] {
		case "forward":
			horz += amt
			depth = depth + aim * amt
		case "down":
			aim += amt
		case "up":
			aim -= amt
		}
	}
	return horz, depth
}

// check panics if err != nil
func check(err error) {
	if err != nil {
		panic(err)
	}
}
